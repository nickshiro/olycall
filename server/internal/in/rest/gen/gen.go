package gen

import (
	"context"
	"errors"

	"olycall-server/internal/core/domain"
	"olycall-server/pkg/typesocket"

	"github.com/google/uuid"
)

const (
	MethodSendMessage = "sendMessage"
	EventNewMessage   = "newMessage"
)

type SendMessageParams struct {
	UserID    uuid.UUID  `json:"user_id"`
	ReplyToID *uuid.UUID `json:"reply_to_id"`
	Content   *string    `json:"content"`
}

type SendMessageError interface {
	typesocket.RPCErrorable
	sendMessageError()
}

type UserNotFoundError struct {
	Message string    `json:"-"`
	UserID  uuid.UUID `json:"user_id"`
}

func (e UserNotFoundError) sendMessageError() {}
func (e UserNotFoundError) Error() string {
	return e.Message
}

func (e UserNotFoundError) ToRPCError() *typesocket.RPCError {
	return &typesocket.RPCError{
		Code:    -32009,
		Message: e.Message,
		Data:    e,
	}
}

func OnSendMessage(
	s *typesocket.Server,
	handler func(ctx context.Context, params *SendMessageParams) (domain.Message, error),
) {
	typesocket.Register(s, MethodSendMessage, func(
		ctx context.Context,
		params *SendMessageParams,
	) (domain.Message, *typesocket.RPCError) {
		res, err := handler(ctx, params)
		if err != nil {
			var specErr SendMessageError
			if errors.As(err, &specErr) {
				return domain.Message{}, specErr.ToRPCError()
			}

			return domain.Message{}, typesocket.NewInternalError(err.Error())
		}

		return res, nil
	})
}

func EmitNewMessageEvent(data *domain.Message) []byte {
	return typesocket.MakeEvent(EventNewMessage, data)
}

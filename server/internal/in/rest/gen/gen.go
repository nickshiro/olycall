package gen

import (
	"context"
	"errors"
	"time"

	"olycall-server/pkg/typesocket"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	AvatarURL *string   `json:"avatar_url"`
}

type Message struct {
	ID                 uuid.UUID  `json:"id"`
	Sender             User       `json:"sender"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          *time.Time `json:"updated_at"`
	ReplyToID          *uuid.UUID `json:"reply_to_id"`
	ForwardedMessageID *uuid.UUID `json:"forwarded_message_id"`
	Content            *string    `json:"content"`
}

type ChatType string

const (
	ChatTypeGroup  ChatType = "group"
	ChatTypeDirect ChatType = "direct"
)

type Chat struct {
	ID          uuid.UUID `json:"id"`
	Type        ChatType  `json:"type"`
	Muted       bool      `json:"muted"`
	Pinned      bool      `json:"pinned"`
	LastMessage *Message  `json:"last_message"`
	Name        string    `json:"name"`
	AvatarURL   *string   `json:"avatar_url"`
}

const (
	MethodSendMessage = "sendMessage"
	MethodGetChats    = "getChats"
	MethodSearchUsers = "searchUsers"
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
	handler func(ctx context.Context, params *SendMessageParams) (Message, error),
) {
	typesocket.Register(s, MethodSendMessage, func(
		ctx context.Context,
		params *SendMessageParams,
	) (Message, *typesocket.RPCError) {
		res, err := handler(ctx, params)
		if err != nil {
			var specErr SendMessageError
			if errors.As(err, &specErr) {
				return Message{}, specErr.ToRPCError()
			}

			return Message{}, typesocket.NewInternalError(err.Error())
		}

		return res, nil
	})
}

func OnSearchUsers(
	s *typesocket.Server,
	handler func(ctx context.Context, params string) ([]User, error),
) {
	typesocket.Register(s, MethodSearchUsers, func(
		ctx context.Context,
		params *string,
	) ([]User, *typesocket.RPCError) {
		res, err := handler(ctx, *params)
		if err != nil {
			return []User{}, typesocket.NewInternalError(err.Error())
		}

		return res, nil
	})
}

func OnGetChats(
	s *typesocket.Server,
	handler func(ctx context.Context) ([]Chat, error),
) {
	typesocket.Register(s, MethodGetChats, func(
		ctx context.Context,
		params *struct{},
	) ([]Chat, *typesocket.RPCError) {
		res, err := handler(ctx)
		if err != nil {
			return []Chat{}, typesocket.NewInternalError(err.Error())
		}

		return res, nil
	})
}

func EmitNewMessageEvent(data *Message) []byte {
	return typesocket.MakeEvent(EventNewMessage, data)
}

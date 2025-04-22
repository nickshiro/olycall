package rest

import (
	"context"
	"errors"
	"net/http"

	"olycall-server/internal/core"
	"olycall-server/internal/core/domain"
	"olycall-server/internal/in/rest/gen"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type wsContextKey string

const (
	userIDContextKey wsContextKey = "userID"
)

func GetUserIDFromContext(ctx context.Context) (uuid.UUID, bool) {
	userID, ok := ctx.Value(userIDContextKey).(uuid.UUID)
	return userID, ok
}

func (c Controller) makeWsHandler(h func(
	w http.ResponseWriter,
	r *http.Request,
) error,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			c.logger.InfoContext(r.Context(), "WS handler", "error", err)
		}
	}
}

func (c Controller) primaryWs(w http.ResponseWriter, r *http.Request) error {
	accessToken := c.getAccessTokenFromCtx(r.Context())

	conn, err := c.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	// FIXME: check if user exists
	userID, err := c.service.GetUserIDFromAccessToken(accessToken)
	if err != nil {
		return err
	}

	c.connectionStore.CreateConn(r.Context(), userID, conn)

	ctxWithUserID := context.WithValue(r.Context(), userIDContextKey, userID)

	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			return err
		}

		resp, err := c.tserver.HandleRequest(ctxWithUserID, data)
		if err != nil {
			return err
		}

		if err := conn.WriteMessage(websocket.TextMessage, resp); err != nil {
		}
	}
}

func (c Controller) handleSendMessage(
	ctx context.Context,
	params *gen.SendMessageParams,
) (domain.Message, error) {
	senderUserID, ok := ctx.Value(userIDContextKey).(uuid.UUID)
	if !ok {
		panic("failed to get user id from context")
	}

	newMessage, err := c.service.SendMessageToUser(ctx, &core.SendMessageToUserParams{
		UserID:       senderUserID,
		TargetUserID: params.UserID,
		ReplyToID:    params.ReplyToID,
		Content:      params.Content,
	})
	if err != nil {
		unwrappedErr := UnwrapAll(err)
		if errors.Is(err, domain.ErrUserNotFound) {
			return domain.Message{}, gen.UserNotFoundError{
				Message: unwrappedErr.Error(),
				UserID:  params.UserID,
			}
		}

		return domain.Message{}, unwrappedErr
	}

	return newMessage, nil
}

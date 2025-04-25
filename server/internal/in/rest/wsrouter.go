package rest

import (
	"context"
	"errors"
	"fmt"
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
		return fmt.Errorf("upgrade websocket: %w", err)
	}
	defer conn.Close()

	userID, err := c.service.GetUserIDFromAccessToken(r.Context(), accessToken)
	if err != nil {
		return fmt.Errorf("get user_id from access token: %w", err)
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
) (gen.Message, error) {
	senderUserID, ok := ctx.Value(userIDContextKey).(uuid.UUID)
	if !ok {
		panic("send message: failed to get user id from context")
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
			return gen.Message{}, gen.UserNotFoundError{
				Message: unwrappedErr.Error(),
				UserID:  params.UserID,
			}
		}

		return gen.Message{}, unwrappedErr
	}

	return gen.Message{
		ID: newMessage.ID,
		Sender: gen.User{
			ID:        newMessage.Sender.ID,
			Username:  newMessage.Sender.Username,
			Name:      newMessage.Sender.Name,
			AvatarURL: newMessage.Sender.AvatarURL,
		},
		CreatedAt:          newMessage.CreatedAt,
		UpdatedAt:          newMessage.UpdatedAt,
		ReplyToID:          newMessage.ReplyToID,
		ForwardedMessageID: newMessage.ForwardedMessageID,
		Content:            newMessage.Content,
	}, nil
}

func (c Controller) handleGetChats(ctx context.Context) ([]gen.Chat, error) {
	senderUserID, ok := ctx.Value(userIDContextKey).(uuid.UUID)
	if !ok {
		panic("get chats: failed to get user id from context")
	}

	chatList, err := c.service.GetChats(ctx, senderUserID)
	if err != nil {
		unwrappedErr := UnwrapAll(err)
		return []gen.Chat{}, unwrappedErr
	}

	chats := make([]gen.Chat, len(chatList))
	for i, chat := range chatList {
		chats[i] = gen.Chat{
			ID:    chat.ID,
			Type:  gen.ChatType(chat.Type),
			Muted: chat.Muted,
			LastMessage: &gen.Message{
				ID:                 chat.LastMessage.ID,
				Sender:             gen.User(chat.LastMessage.Sender),
				CreatedAt:          chat.LastMessage.CreatedAt,
				UpdatedAt:          chat.LastMessage.UpdatedAt,
				ReplyToID:          chat.LastMessage.ReplyToID,
				ForwardedMessageID: chat.LastMessage.ForwardedMessageID,
				Content:            chat.LastMessage.Content,
			},
			Name:      chat.Name,
			AvatarURL: chat.AvatarURL,
		}
	}

	return chats, nil
}

func (c Controller) handleSearchUsers(
	ctx context.Context,
	query string,
) ([]gen.User, error) {
	usersResp, err := c.service.SearchUsers(ctx, query)
	if err != nil {
		unwrappedErr := UnwrapAll(err)
		return []gen.User{}, unwrappedErr
	}

	users := make([]gen.User, len(usersResp))
	for i, user := range usersResp {
		users[i] = gen.User(user)
	}

	return users, nil
}

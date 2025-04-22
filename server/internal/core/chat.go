package core

import (
	"context"
	"fmt"
	"time"

	"olycall-server/internal/core/domain"
	"olycall-server/internal/core/ports/chatstore"
	"olycall-server/pkg/uuidrule"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

type SendMessageToUserParams struct {
	UserID       uuid.UUID  `json:"user_id"`
	TargetUserID uuid.UUID  `json:"target_user_id"`
	ReplyToID    *uuid.UUID `json:"reply_to_id"`
	Content      *string    `json:"content"`
}

func (s Service) SendMessageToUser(
	ctx context.Context,
	params *SendMessageToUserParams,
) (domain.Message, error) {
	if err := validation.ValidateStructWithContext(ctx, params,
		validation.Field(&params.UserID, validation.By(uuidrule.Required), validation.NotIn(params.TargetUserID)),
		validation.Field(&params.TargetUserID, validation.By(uuidrule.Required)),
		validation.Field(&params.Content, validation.Required),
	); err != nil {
		return domain.Message{}, wrapInvalidParamsErr(err)
	}

	chatID, err := s.chatStore.GetDirectChatIDBy2UsersID(ctx, params.UserID, params.TargetUserID)
	if err != nil {
		return domain.Message{}, fmt.Errorf("get direct_chat_id by 2 user_ids: %w", err)
	}

	now := time.Now()

	if chatID == nil {
		c := uuid.New()
		chatID = &c

		if err := s.chatStore.CreateChat(ctx, &chatstore.Chat{
			ID:   c,
			Type: chatstore.ChatTypeDirect,
		}); err != nil {
			return domain.Message{}, fmt.Errorf("create chat: %w", err)
		}

		if err := s.chatStore.CreateDirectChat(ctx, &chatstore.DirectChat{
			ChatID:  c,
			User1ID: params.UserID,
			User2ID: params.TargetUserID,
		}); err != nil {
			return domain.Message{}, fmt.Errorf("create direct chat: %w", err)
		}
	}

	messageID := uuid.New()
	if err := s.chatStore.CreateMessage(ctx, &chatstore.Message{
		ID:        messageID,
		SenderID:  params.UserID,
		ChatID:    *chatID,
		ReplyToID: params.ReplyToID,
		Content:   params.Content,
		CreatedAt: now,
		UpdatedAt: nil,
	}); err != nil {
		return domain.Message{}, fmt.Errorf("create message: %w", err)
	}

	targetUserConns := s.connectionStore.GetConnsByUserID(params.TargetUserID)

	newMessage := domain.Message{
		ID:        messageID,
		SenderID:  params.UserID,
		Content:   params.Content,
		CreatedAt: now,
	}

	s.notificationsProvider.NewMessage(ctx, targetUserConns, &newMessage)

	return newMessage, nil
}

// func (s Service) GetChats(
// 	ctx context.Context,
// 	userID uuid.UUID,
// ) (domain.ChatList, error) {

// }

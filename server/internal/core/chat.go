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
		// TODO: wrap into transaction, use deferable constraints

		if err := s.chatStore.CreateChat(ctx, &chatstore.Chat{
			ID:   c,
			Type: domain.ChatTypeDirect,
		}); err != nil {
			return domain.Message{}, fmt.Errorf("create chat: %w", err)
		}

		if err := s.chatStore.CreateUserChat(ctx, &chatstore.UserChat{
			UserID: params.UserID,
			ChatID: c,
			Pinned: false,
			Muted:  false,
		}); err != nil {
			return domain.Message{}, fmt.Errorf("create sender user_chat: %w", err)
		}

		if err := s.chatStore.CreateUserChat(ctx, &chatstore.UserChat{
			UserID: params.TargetUserID,
			ChatID: c,
			Pinned: false,
			Muted:  false,
		}); err != nil {
			return domain.Message{}, fmt.Errorf("create target user_chat: %w", err)
		}

		if err := s.chatStore.CreateDirectChat(ctx, &chatstore.DirectChat{
			ChatID:  c,
			User1ID: params.UserID,
			User2ID: params.TargetUserID,
		}); err != nil {
			return domain.Message{}, fmt.Errorf("create direct_chat: %w", err)
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

	user, err := s.userStore.GetUserByID(ctx, params.UserID)
	if err != nil {
		return domain.Message{}, fmt.Errorf("get user by id: %w", err)
	}

	if user == nil {
		return domain.Message{}, domain.ErrUserNotFound
	}

	newMessage := domain.Message{
		ID: messageID,
		Sender: domain.User{
			ID:        user.ID,
			Username:  user.Username,
			Name:      user.Name,
			AvatarURL: user.AvatarURL,
		},
		Content:   params.Content,
		CreatedAt: now,
		UpdatedAt: nil,
	}

	s.notificationsProvider.NewMessage(ctx, targetUserConns, &newMessage)

	return newMessage, nil
}

func (s Service) GetChats(
	ctx context.Context,
	userID uuid.UUID,
) ([]domain.Chat, error) {
	userChatsResp, err := s.chatStore.GetUserChats(ctx, userID)
	if err != nil {
		return []domain.Chat{}, fmt.Errorf("get user chats: %w", err)
	}

	chats := make([]domain.Chat, len(userChatsResp))

	for i, userChatResp := range userChatsResp {
		chat := domain.Chat{
			ID:     userChatResp.ChatID,
			Type:   userChatResp.Type,
			Muted:  userChatResp.Muted,
			Pinned: userChatResp.Pinned,
		}

		if userChatResp.LastMessageID != nil {
			chat.LastMessage = &domain.Message{
				ID:        *userChatResp.LastMessageID,
				CreatedAt: *userChatResp.LastMessageCreatedAt,
				UpdatedAt: userChatResp.LastMessageUpdatedAt,
				Content:   userChatResp.LastMessageContent,
			}
		}

		switch userChatResp.Type {
		case domain.ChatTypeDirect:
			chat.Name = *userChatResp.Username
			chat.AvatarURL = userChatResp.UserAvatarURL

		default:
			panic("unknown chat type")
		}

		chats[i] = chat
	}

	return chats, nil
}

package domain

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID        uuid.UUID  `json:"id"`
	SenderID  uuid.UUID  `json:"sender_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	// ReplyToID          *uuid.UUID `json:"reply_to_id"`
	// ForwardedMessageID *uuid.UUID `json:"forwarded_message_id"`
	Content *string `json:"content"`
}

type Chat struct {
	ID          uuid.UUID `json:"id"`
	Type        string    `json:"type"`
	Name        string    `json:"name"`
	AvatarURL   *string   `json:"avatar_url"`
	LastMessage *Message  `json:"last_message"`
}

type ChatList struct {
	Pinned []Chat `json:"pinned"`
	Others []Chat `json:"others"`
}

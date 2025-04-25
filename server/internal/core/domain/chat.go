package domain

import (
	"time"

	"github.com/google/uuid"
)

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
	Pinned      bool      `json:"pinned"`
	Muted       bool      `json:"muted"`
	LastMessage *Message  `json:"last_message"`
	Name        string    `json:"name"`
	AvatarURL   *string   `json:"avatar_url"`
}

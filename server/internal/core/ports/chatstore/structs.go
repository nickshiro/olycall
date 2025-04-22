package chatstore

import (
	"time"

	"github.com/google/uuid"
)

type ChatType string

const (
	ChatTypeGroup  ChatType = "group"
	ChatTypeDirect ChatType = "direct"
)

type Chat struct {
	ID   uuid.UUID `json:"id"`
	Type ChatType  `json:"type"`
}

type GroupChat struct {
	ChatID      uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	AvatarURL   *string   `json:"avatar_url"`
	Description *string   `json:"description"`
}

type GroupChatMember struct {
	GroupChatID uuid.UUID `json:"group_chat_id"`
	UserID      uuid.UUID `json:"user_id"`
	JoinedAt    time.Time `json:"joined_at"`
}

type DirectChat struct {
	ChatID  uuid.UUID `json:"chat_id"`
	User1ID uuid.UUID `json:"user1_id"`
	User2ID uuid.UUID `json:"user2_id"`
}

type Message struct {
	ID                 uuid.UUID  `json:"id"`
	SenderID           uuid.UUID  `json:"sender_id"`
	ChatID             uuid.UUID  `json:"chat_id"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          *time.Time `json:"updated_at"`
	ReplyToID          *uuid.UUID `json:"reply_to_id"`
	ForwardedMessageID *uuid.UUID `json:"forwarded_message_id"`
	Content            *string    `json:"content"`
}

type File struct {
	ID              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	MimeType        string    `json:"mime_type"`
	SizeBytes       int64     `json:"size_bytes"`
	UploadTimestamp time.Time `json:"upload_timestamp"`
}

type MessageFile struct {
	MessageID uuid.UUID `json:"message_id"`
	FileID    uuid.UUID `json:"file_id"`
}

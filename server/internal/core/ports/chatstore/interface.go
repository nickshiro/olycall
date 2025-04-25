package chatstore

import (
	"context"

	"github.com/google/uuid"
)

type ChatStore interface {
	CreateChat(ctx context.Context, chat *Chat) error
	GetChatByID(ctx context.Context, id uuid.UUID) (*Chat, error)

	CreateGroupChatMember(ctx context.Context, groupChatMember *GroupChatMember) error
	// GetChatIDBy2UserIDs(ctx context.Context, user1ID uuid.UUID, user2ID uuid.UUID) (*uuid.UUID, error)

	CreateUserChat(ctx context.Context, userChat *UserChat) error

	CreateDirectChat(ctx context.Context, directChat *DirectChat) error
	GetDirectChatIDBy2UsersID(ctx context.Context, user1ID uuid.UUID, user2ID uuid.UUID) (*uuid.UUID, error)

	GetUserChats(ctx context.Context, userID uuid.UUID) ([]UserChatsResp, error)

	CreateMessage(ctx context.Context, message *Message) error

	WithTx(ctx context.Context, fn func(ctx context.Context, store ChatStore) error) error
}

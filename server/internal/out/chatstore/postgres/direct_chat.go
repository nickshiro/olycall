package postgres

import (
	"context"

	"olycall-server/internal/core/ports/chatstore"

	"github.com/google/uuid"
)

func (s ChatStore) CreateDirectChat(ctx context.Context, directChat *chatstore.DirectChat) error {
	_, err := s.db.Exec(ctx,
		`
		INSERT INTO
		    direct_chat (
		        chat_id,
		        user1_id,
		        user2_id
		    )
		VALUES ($1, $2, $3)
		`,
		directChat.ChatID,
		directChat.User1ID,
		directChat.User2ID,
	)

	return s.mapError(err)
}

func (s ChatStore) GetDirectChatIDBy2UsersID(
	ctx context.Context,
	user1ID uuid.UUID,
	user2ID uuid.UUID,
) (*uuid.UUID, error) {
	var chatID uuid.UUID
	if err := s.db.QueryRow(ctx,
		`
		SELECT chat_id
		FROM direct_chat
		WHERE (user1_id = $1 AND user2_id = $2)
		   OR (user1_id = $2 AND user2_id = $1);
		`,
		user1ID,
		user2ID,
	).Scan(&chatID); err != nil {
		return nil, s.mapError(err)
	}

	return &chatID, nil
}

package postgres

import (
	"context"

	"olycall-server/internal/core/ports/chatstore"

	"github.com/google/uuid"
)

func (s ChatStore) CreateChat(ctx context.Context, chat *chatstore.Chat) error {
	q := s.getQuerier()
	_, err := q.Exec(ctx,
		`
		INSERT INTO chat (
		    id,
		    type
		) VALUES ($1, $2)
		`,
		chat.ID,
		chat.Type,
	)

	return s.mapError(err)
}

func (s ChatStore) GetChatByID(ctx context.Context, id uuid.UUID) (*chatstore.Chat, error) {
	q := s.getQuerier()

	var resp chatstore.Chat
	if err := q.QueryRow(ctx,
		`
		SELECT
		    id,
		    type
		FROM chat
		WHERE id = $1
		LIMIT 1
		`,
		id,
	).Scan(
		&resp.ID,
		&resp.Type,
	); err != nil {
		return nil, s.mapError(err)
	}

	resp.ID = id

	return &resp, nil
}

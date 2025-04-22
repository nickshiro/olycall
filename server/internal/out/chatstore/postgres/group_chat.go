package postgres

import (
	"context"

	"olycall-server/internal/core/ports/chatstore"
)

func (s ChatStore) CreateGroupChat(ctx context.Context, groupChat *chatstore.GroupChat) error {
	q := s.getQuerier()

	_, err := q.Exec(ctx,
		`
		INSERT INTO
		    group_chat (
		        chat_id,
				name,
				avatar_url,
				description
		    )
		VALUES ($1, $2, $3, $4)
		`,
		groupChat.ChatID,
		groupChat.Name,
		groupChat.AvatarURL,
		groupChat.Description,
	)

	return s.mapError(err)
}

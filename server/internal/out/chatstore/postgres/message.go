package postgres

import (
	"context"

	"olycall-server/internal/core/ports/chatstore"
)

func (s ChatStore) CreateMessage(ctx context.Context, message *chatstore.Message) error {
	q := s.getQuerier()

	_, err := q.Exec(ctx,
		`
		INSERT INTO message (
			id,
			sender_id,
			chat_id,
			created_at,
			updated_at,
			reply_to_id,
			forwarded_message_id,
			content
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		`,
		message.ID,
		message.SenderID,
		message.ChatID,
		message.CreatedAt,
		message.UpdatedAt,
		message.ReplyToID,
		message.ForwardedMessageID,
		message.Content,
	)

	return s.mapError(err)
}

// func (s ChatStore) GetMessageByID(ctx context.Context, id uuid.UUID) (*chatstore.Chat, error) {
// 	q := s.getQuerier()

// 	var resp chatstore.Chat
// 	if err := q.QueryRow(ctx,
// 		`
// 		SELECT
// 	        name,
// 	        avatar_url
// 		FROM chat
// 		WHERE id = $1
// 		LIMIT 1
// 		`,
// 		id,
// 	).Scan(
// 		&resp.Name,
// 		&resp.AvatarURL,
// 	); err != nil {
// 		return nil, s.mapError(err)
// 	}

// 	resp.ID = id

// 	return &resp, nil
// }

// func (s ChatStore) UpdateChat(ctx context.Context, params *chatstore.UpdateChatParams) (bool, error) {
// 	q := s.getQuerier()

// 	_, err := q.Exec(ctx,
// 		`
// 		UPDATE chat
// 		SET
// 	        name = $2
// 	        avatar_url = $3
//     	WHERE id = $1
// 		 `,
// 		params.ID,
// 		params.Name,
// 		params.AvatarURL,
// 	)
// 	if err != nil {
// 		return false, s.mapError(err)
// 	}

// 	return true, nil
// }

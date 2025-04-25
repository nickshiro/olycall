package postgres

import (
	"context"

	"olycall-server/internal/core/ports/chatstore"

	"github.com/google/uuid"
)

func (s ChatStore) CreateUserChat(ctx context.Context, userChat *chatstore.UserChat) error {
	_, err := s.db.Exec(ctx,
		`
		INSERT INTO
		    user_chat (
		        user_id,
		        chat_id,
				pinned,
				muted
		    )
		VALUES ($1, $2, $3, $4)
		`,
		userChat.UserID,
		userChat.ChatID,
		userChat.Pinned,
		userChat.Muted,
	)

	return s.mapError(err)
}

func (s ChatStore) GetUserChats(
	ctx context.Context,
	userID uuid.UUID,
) ([]chatstore.UserChatsResp, error) {
	rows, err := s.db.Query(ctx,
		`
		WITH
		user_chats AS (
			SELECT uc.chat_id, uc.pinned, uc.muted, c.type
			FROM user_chat uc
			JOIN chat c ON c.id = uc.chat_id
			WHERE uc.user_id = $1
		),

		last_messages AS (
			SELECT DISTINCT ON (chat_id)
				chat_id,
				id AS message_id,
				content,
				created_at,
				updated_at,
				reply_to_id,
				forwarded_message_id
			FROM message
			ORDER BY chat_id, created_at DESC
		),

		direct_chat_profiles AS (
			SELECT
				dc.chat_id,
				u.id AS other_user_id,
				u.username,
				u.avatar_url
			FROM direct_chat dc
			JOIN app_user u ON (u.id = dc.user1_id OR u.id = dc.user2_id)
			WHERE u.id <> $1
		)

		SELECT
			uc.chat_id,
			uc.type,
			uc.pinned,
			uc.muted,

			lm.message_id AS last_message_id,
			lm.content AS last_message_content,
			lm.created_at AS last_message_created_at,
			lm.updated_at AS last_message_updated_at,
			lm.reply_to_id AS last_message_reply_to_id,
			lm.forwarded_message_id AS last_message_forwarded_message_id,

			dcp.other_user_id AS user_id,
			dcp.username,
			dcp.avatar_url
		FROM user_chats uc
		LEFT JOIN last_messages lm ON lm.chat_id = uc.chat_id
		LEFT JOIN direct_chat_profiles dcp ON dcp.chat_id = uc.chat_id
		ORDER BY uc.pinned DESC, lm.created_at DESC NULLS LAST;
		`,
		userID,
	)
	if err != nil {
		return nil, s.mapError(err)
	}
	defer rows.Close()

	var chats []chatstore.UserChatsResp

	for rows.Next() {
		var chat chatstore.UserChatsResp

		if err := rows.Scan(
			&chat.ChatID,
			&chat.Type,
			&chat.Pinned,
			&chat.Muted,
			&chat.LastMessageID,
			&chat.LastMessageContent,
			&chat.LastMessageCreatedAt,
			&chat.LastMessageUpdatedAt,
			&chat.LastMessageReplyToID,
			&chat.LastMessageForwardedID,
			&chat.UserID,
			&chat.Username,
			&chat.UserAvatarURL,
		); err != nil {
			return nil, s.mapError(err)
		}

		chats = append(chats, chat)
	}

	if err := rows.Err(); err != nil {
		return nil, s.mapError(err)
	}

	return chats, nil
}

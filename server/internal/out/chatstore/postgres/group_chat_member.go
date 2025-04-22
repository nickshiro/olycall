package postgres

import (
	"context"

	"olycall-server/internal/core/ports/chatstore"
)

func (s ChatStore) CreateGroupChatMember(ctx context.Context, groupChatMember *chatstore.GroupChatMember) error {
	_, err := s.db.Exec(ctx,
		`
		INSERT INTO
		    group_chat_member (
		        group_chat_id,
		        user_id,
		        joined_at
		    )
		VALUES ($1, $2, $3)
		`,
		groupChatMember.GroupChatID,
		groupChatMember.UserID,
		groupChatMember.JoinedAt,
	)

	return s.mapError(err)
}

// func (s ChatStore) GetChatIDBy2UserIDs(
// 	ctx context.Context,
// 	user1ID uuid.UUID,
// 	user2ID uuid.UUID,
// ) (*uuid.UUID, error) {
// 	var chatID uuid.UUID
// 	if err := s.db.QueryRow(ctx,
// 		`
// 		SELECT cm.chat_id
// 		FROM chat_member cm
// 		WHERE cm.user_id IN ($1, $2)
// 		GROUP BY cm.chat_id
// 		HAVING COUNT(*) = 2
// 			AND COUNT(*) = (
// 				SELECT COUNT(*)
// 				FROM chat_member cm2
// 				WHERE cm2.chat_id = cm.chat_id
// 			);
// 		`,
// 		user1ID,
// 		user2ID,
// 	).Scan(&chatID); err != nil {
// 		return nil, s.mapError(err)
// 	}
//
// 	return &chatID, nil
// }

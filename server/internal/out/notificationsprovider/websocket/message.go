package websocket

import (
	"context"

	"olycall-server/internal/core/domain"
	"olycall-server/internal/in/rest/gen"

	"github.com/gorilla/websocket"
)

func (p *NotificationsProvider) NewMessage(
	ctx context.Context,
	conns []*websocket.Conn,
	data *domain.Message,
) {
	msg := gen.EmitNewMessageEvent(&gen.Message{
		ID:                 data.ID,
		Sender:             gen.User(data.Sender),
		CreatedAt:          data.CreatedAt,
		UpdatedAt:          data.UpdatedAt,
		ReplyToID:          data.ReplyToID,
		ForwardedMessageID: data.ForwardedMessageID,
		Content:            data.Content,
	})

	for _, conn := range conns {
		if conn == nil {
			continue
		}

		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			continue
		}
	}
}

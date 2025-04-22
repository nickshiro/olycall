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
	msg := gen.EmitNewMessageEvent(data)

	for _, conn := range conns {
		if conn == nil {
			continue
		}

		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			continue
		}
	}
}

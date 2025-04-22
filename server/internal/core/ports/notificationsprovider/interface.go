package notificationsprovider

import (
	"context"

	"olycall-server/internal/core/domain"

	"github.com/gorilla/websocket"
)

type NotificationsProvider interface {
	NewMessage(ctx context.Context, conns []*websocket.Conn, data *domain.Message)
}

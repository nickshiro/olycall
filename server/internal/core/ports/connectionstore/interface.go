package connectionstore

import (
	"context"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type ConnectionStore interface {
	CreateConn(ctx context.Context, userID uuid.UUID, conn *websocket.Conn)
	DeleteConn(ctx context.Context, userID uuid.UUID, conn *websocket.Conn)
	GetConnsByUserID(userID uuid.UUID) []*websocket.Conn
	GetConnsByUserIDs(userIDs []uuid.UUID) []*websocket.Conn
}

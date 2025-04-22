package inmemory

import (
	"context"
	"slices"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func (p *ConnectionStore) CreateConn(ctx context.Context, userID uuid.UUID, conn *websocket.Conn) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.connections[userID] = append(p.connections[userID], conn)
}

func (p *ConnectionStore) DeleteConn(ctx context.Context, userID uuid.UUID, conn *websocket.Conn) {
	p.mu.Lock()
	defer p.mu.Unlock()

	conns := p.connections[userID]
	for i, c := range conns {
		if c == conn {
			conns = slices.Delete(conns, i, i+1)

			break
		}
	}

	if len(conns) == 0 {
		delete(p.connections, userID)
	} else {
		p.connections[userID] = conns
	}
}

func (p *ConnectionStore) GetConnsByUserID(userID uuid.UUID) []*websocket.Conn {
	p.mu.RLock()
	defer p.mu.RUnlock()

	// Return a copy to avoid exposing internal slice
	conns := p.connections[userID]
	result := make([]*websocket.Conn, len(conns))
	copy(result, conns)

	return result
}

func (p *ConnectionStore) GetConnsByUserIDs(userIDs []uuid.UUID) []*websocket.Conn {
	p.mu.RLock()
	defer p.mu.RUnlock()

	var result []*websocket.Conn

	for _, userID := range userIDs {
		conns := p.connections[userID]
		result = append(result, conns...)
	}

	return result
}

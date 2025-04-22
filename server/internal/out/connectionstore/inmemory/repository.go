package inmemory

import (
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type ConnectionStore struct {
	connections map[uuid.UUID][]*websocket.Conn
	mu          sync.RWMutex
}

func New() *ConnectionStore {
	return &ConnectionStore{
		connections: make(map[uuid.UUID][]*websocket.Conn),
		mu:          sync.RWMutex{},
	}
}

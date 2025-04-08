package service

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Email     string
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

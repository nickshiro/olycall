package userstore

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Email     string
	Username  string
	CreatedAt time.Time
}

type UpdateUserParams struct {
	ID       uuid.UUID
	Username string
}

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
	UpdatedAt time.Time
}

type CreateUserParams struct {
	ID        uuid.UUID
	Email     string
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UpdateUserParams struct {
	ID        uuid.UUID
	Username  string
	UpdatedAt time.Time
}

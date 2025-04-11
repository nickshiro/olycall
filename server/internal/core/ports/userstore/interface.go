package userstore

import (
	"context"

	"github.com/google/uuid"
)

type UserStore interface {
	CreateUser(ctx context.Context, params *CreateUserParams) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*User, error)
	UpdateUser(ctx context.Context, params *UpdateUserParams) (bool, error)
}

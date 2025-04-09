package userstore

import (
	"context"

	"github.com/google/uuid"
)

type UserStore interface {
	CreateUser(context.Context, *CreateUserParams) error
	GetUserByEmail(context.Context, string) (*User, error)
	GetUserByID(context.Context, uuid.UUID) (*User, error)
	UpdateUser(context.Context, *UpdateUserParams) (bool, error)
}

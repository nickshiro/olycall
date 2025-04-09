package oauthstatestore

import (
	"context"

	"github.com/google/uuid"
)

type OAuthStateStore interface {
	CreateOAuthState(ctx context.Context, params *CreateOAuthStateParams) error
	GetOAuthState(ctx context.Context, id uuid.UUID) (*OAuthState, error)
	DeleteOAuthState(ctx context.Context, id uuid.UUID) (bool, error)
}

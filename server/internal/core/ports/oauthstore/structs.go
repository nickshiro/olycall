package oauthstore

import (
	"time"

	"github.com/google/uuid"
)

type CreateOauthIdentityParams struct {
	ID         int
	UserID     uuid.UUID
	ProviderID int
	Subject    string
	CreatedAt  time.Time
}

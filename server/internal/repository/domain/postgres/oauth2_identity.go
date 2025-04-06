package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type CreateOauth2IdentityParams struct {
	ID         int
	UserID     uuid.UUID
	ProviderID int
	Subject    string
	CreatedAt  time.Time
}

func (r Repository) CreateOauth2Identity(ctx context.Context, params *CreateOauth2IdentityParams) error {
	_, err := r.db.Exec(ctx,
		`
		INSERT INTO
		    oauth2_identity (
		        id,
		        user_id,
		        provider_id,
		        subject,
			    created_at
		    )
		VALUES ($1, $2, $3, $4, $5)
		`,
		params.ID,
		params.UserID,
		params.ProviderID,
		params.Subject,
		params.CreatedAt,
	)

	return r.handleError(err)
}

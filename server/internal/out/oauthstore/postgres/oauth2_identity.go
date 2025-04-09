package postgres

import (
	"context"

	"olycall-server/internal/core/ports/oauthstore"
	"olycall-server/pkg/pg"
)

func (r Repository) CreateOauthIdentity(ctx context.Context, params *oauthstore.CreateOauthIdentityParams) error {
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

	return pg.MapError(err) // nolint: wrapcheck
}

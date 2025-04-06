package postgres

import (
	"context"

	"olycall-server/internal/repository/domain"

	"github.com/google/uuid"
)

func (r Repository) CreateUser(ctx context.Context, params *domain.CreateUserParams) error {
	_, err := r.db.Exec(ctx,
		`
		INSERT INTO
		    app_user (
		        id,
		        email,
		        username,
			    created_at,
			    updated_at
		    )
		VALUES ($1, $2, $3, $4, $5)
		`,
		params.ID,
		params.Email,
		params.Username,
		params.CreatedAt,
		params.UpdatedAt,
	)

	return r.handleError(err)
}

func (r Repository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var resp domain.User
	err := r.db.QueryRow(ctx,
		`
		SELECT
	        id,
	        username,
		    created_at,
		    updated_at
		FROM app_user
		WHERE email = $1
		LIMIT 1
		`,
		email,
	).Scan(
		&resp.ID,
		&resp.Username,
		&resp.CreatedAt,
		&resp.UpdatedAt,
	)
	if err != nil {
		return nil, r.handleError(err)
	}

	resp.Email = email
	return &resp, nil
}

func (r Repository) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	var resp domain.User
	err := r.db.QueryRow(ctx,
		`
		SELECT
	        email,
	        username,
		    created_at,
		    updated_at
		FROM app_user
		WHERE id = $1
		LIMIT 1
		`,
		id,
	).Scan(
		&resp.Email,
		&resp.Username,
		&resp.CreatedAt,
		&resp.UpdatedAt,
	)
	if err != nil {
		return nil, r.handleError(err)
	}

	resp.ID = id
	return &resp, nil
}

func (r Repository) UpdateUser(ctx context.Context, arg *domain.UpdateUserParams) (bool, error) {
	_, err := r.db.Exec(ctx,
		`
		UPDATE app_user
		SET
	        username = $2,
		    updated_at = $3
    	WHERE id = $1
		 `,
		arg.ID,
		arg.Username,
		arg.UpdatedAt,
	)
	if err != nil {
		return false, r.handleError(err)
	}

	return true, nil
}

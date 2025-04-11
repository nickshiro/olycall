package postgres

import (
	"context"
	"olycall-server/internal/core/ports/userstore"

	"github.com/google/uuid"
)

func (s UserStore) CreateUser(ctx context.Context, params *userstore.CreateUserParams) error {
	_, err := s.db.Exec(ctx,
		`
		INSERT INTO
		    app_user (
		        id,
		        email,
		        username,
			    created_at
		    )
		VALUES ($1, $2, $3, $4)
		`,
		params.ID,
		params.Email,
		params.Username,
		params.CreatedAt,
	)

	return s.mapError(err)
}

func (s UserStore) GetUserByEmail(ctx context.Context, email string) (*userstore.User, error) {
	var resp userstore.User
	if err := s.db.QueryRow(ctx,
		`
		SELECT
	        id,
	        username,
		    created_at
		FROM app_user
		WHERE email = $1
		LIMIT 1
		`,
		email,
	).Scan(
		&resp.ID,
		&resp.Username,
		&resp.CreatedAt,
	); err != nil {
		return nil, s.mapError(err)
	}

	resp.Email = email

	return &resp, nil
}

func (s UserStore) GetUserByID(ctx context.Context, id uuid.UUID) (*userstore.User, error) {
	var resp userstore.User
	if err := s.db.QueryRow(ctx,
		`
		SELECT
	        email,
	        username,
		    created_at
		FROM app_user
		WHERE id = $1
		LIMIT 1
		`,
		id,
	).Scan(
		&resp.Email,
		&resp.Username,
		&resp.CreatedAt,
	); err != nil {
		return nil, s.mapError(err)
	}

	resp.ID = id

	return &resp, nil
}

func (s UserStore) UpdateUser(ctx context.Context, arg *userstore.UpdateUserParams) (bool, error) {
	_, err := s.db.Exec(ctx,
		`
		UPDATE app_user
		SET
	        username = $2
    	WHERE id = $1
		 `,
		arg.ID,
		arg.Username,
	)
	if err != nil {
		return false, s.mapError(err)
	}

	return true, nil
}

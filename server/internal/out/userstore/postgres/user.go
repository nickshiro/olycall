package postgres

import (
	"context"

	"olycall-server/internal/core/ports/userstore"

	"github.com/google/uuid"
)

func (s UserStore) CreateUser(ctx context.Context, user *userstore.User) error {
	_, err := s.db.Exec(ctx,
		`
		INSERT INTO
		    app_user (
		        id,
		        email,
		        username,
				name,
				avatar_url,
			    created_at
		    )
		VALUES ($1, $2, $3, $4, $5, $6)
		`,
		user.ID,
		user.Email,
		user.Username,
		user.Name,
		user.AvatarURL,
		user.CreatedAt,
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
			name,
			avatar_url,
		    created_at
		FROM app_user
		WHERE email = $1
		LIMIT 1
		`,
		email,
	).Scan(
		&resp.ID,
		&resp.Username,
		&resp.Name,
		&resp.AvatarURL,
		&resp.CreatedAt,
	); err != nil {
		return nil, s.mapError(err)
	}

	resp.Email = email

	return &resp, nil
}

func (s UserStore) GetUserIDByEmail(ctx context.Context, email string) (*uuid.UUID, error) {
	var id uuid.UUID
	if err := s.db.QueryRow(ctx,
		`
		SELECT id
		FROM app_user
		WHERE email = $1
		LIMIT 1
		`,
		email,
	).Scan(
		&id,
	); err != nil {
		return nil, s.mapError(err)
	}

	return &id, nil
}

func (s UserStore) GetUserByID(ctx context.Context, id uuid.UUID) (*userstore.User, error) {
	var resp userstore.User
	if err := s.db.QueryRow(ctx,
		`
		SELECT
	        email,
	        username,
			name,
			avatar_url,
		    created_at
		FROM app_user
		WHERE id = $1
		LIMIT 1
		`,
		id,
	).Scan(
		&resp.Email,
		&resp.Username,
		&resp.Name,
		&resp.AvatarURL,
		&resp.CreatedAt,
	); err != nil {
		return nil, s.mapError(err)
	}

	resp.ID = id

	return &resp, nil
}

func (s UserStore) CheckUserByID(ctx context.Context, id uuid.UUID) (bool, error) {
	var exists bool
	if err := s.db.QueryRow(ctx,
		`
		SELECT EXISTS (
			SELECT 1
			FROM app_user
			WHERE id = $1
		)
		`,
		id,
	).Scan(&exists); err != nil {
		return false, s.mapError(err)
	}

	return exists, nil
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

func (s UserStore) SearchUsersByUsername(ctx context.Context, query string) ([]userstore.User, error) {
	rows, err := s.db.Query(ctx,
		`
		SELECT 
			id,
			email,
			username,
			name,
			avatar_url,
			created_at
		FROM app_user
		WHERE username ILIKE '%' || $1 || '%' OR name ILIKE '%' || $1 || '%';
		 `,
		query,
	)
	if err != nil {
		return nil, s.mapError(err)
	}

	var chats []userstore.User

	for rows.Next() {
		var chat userstore.User

		if err := rows.Scan(
			&chat.ID,
			&chat.Email,
			&chat.Username,
			&chat.Name,
			&chat.AvatarURL,
			&chat.CreatedAt,
		); err != nil {
			return nil, s.mapError(err)
		}

		chats = append(chats, chat)
	}

	return chats, nil
}

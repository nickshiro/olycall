package postgres

import (
	"errors"

	"olycall-server/internal/core/ports/userstore"
	"olycall-server/pkg/pg"
)

func (s UserStore) mapError(err error) error {
	if err := pg.MapError(err); err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil
		}

		var conflictError pg.ConflictError
		if errors.As(err, &conflictError) {
			switch conflictError.Column {
			case "email":
				return userstore.ErrEmailConflict
			case "username":
				return userstore.ErrUsernameConflict
			}
		}
	}

	return err
}

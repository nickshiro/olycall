package postgres

import (
	"errors"

	"olycall-server/internal/core/ports/chatstore"
	"olycall-server/pkg/pg"
)

func (s ChatStore) mapError(err error) error {
	if err := pg.MapError(err); err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil
		}

		// var conflictError pg.ConflictError
		// if errors.As(err, &conflictError) {
		// 	switch conflictError.Column {
		// 	case "email":
		// 		return userstore.ErrEmailConflict

		// 	case "username":
		// 		return userstore.ErrUsernameConflict
		// 	}
		// }

		var foreignKeyError pg.ForeignKeyError
		if errors.As(err, &foreignKeyError) {
			switch foreignKeyError.TableName {
			case "app_user":
				return chatstore.ErrUserNotFound
			}
		}
	}

	return err
}

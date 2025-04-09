package pg

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var ErrNoRows = errors.New("no rows")

type ConflictError struct {
	Column string
}

func (e ConflictError) Error() string {
	return fmt.Sprintf("unique constraint violation on column %s", e.Column)
}

func MapError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return ErrNoRows
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return ConflictError{
				Column: pgErr.TableName,
			}
		default:
			return err
		}
	}

	return err
}

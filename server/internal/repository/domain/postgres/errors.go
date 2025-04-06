package postgres

import (
	"errors"

	"olycall-server/internal/repository/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r Repository) handleError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return nil
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return domain.ConflictError{
				Field: pgErr.TableName,
			}
		default:
			return err
		}
	}

	return err
}

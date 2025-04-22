package pg

import (
	"errors"
	"regexp"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var tableNameRegex = regexp.MustCompile(`is not present in table "([^"]+)"`)

var ErrNoRows = errors.New("no rows")

type ConflictError struct {
	Column string
}

func (e ConflictError) Error() string {
	return "unique constraint violation on column " + e.Column
}

type ForeignKeyError struct {
	TableName string
}

func (e ForeignKeyError) Error() string {
	return "foreign key constraint violation on table" + e.TableName
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

		case "23503":
			matches := tableNameRegex.FindStringSubmatch(pgErr.Detail)
			if len(matches) == 2 {
				return ForeignKeyError{
					TableName: matches[1],
				}
			}

		default:
			return err
		}
	}

	return err
}

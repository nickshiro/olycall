package service

import (
	"errors"
	"fmt"
)

type ConflictError struct {
	Field string
}

func (e ConflictError) Error() string {
	return fmt.Sprintf("%s field unique constraint violation", e.Field)
}

var (
	ErrNotFound     = errors.New("not found")
	ErrInvalidToken = errors.New("invalid token")
	ErrUserNotFound = wrapNotFoundErr("user")
	ErrNoteNotFound = wrapNotFoundErr("note")
	ErrTagNotFound  = wrapNotFoundErr("tag")
)

func wrapNotFoundErr(field string) error {
	return fmt.Errorf("%s %w", field, ErrNotFound)
}

func wrapInvalidParamsErr(err error) error {
	return fmt.Errorf("invalid params: %w", err)
}

package uuidrule

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidUUID = errors.New("invalid uuid type")
	ErrEmptyUUID   = errors.New("uuid must not be empty")
)

func Required(value any) error {
	id, ok := value.(uuid.UUID)
	if !ok {
		return ErrInvalidUUID
	}

	if id == uuid.Nil {
		return ErrEmptyUUID
	}

	return nil
}

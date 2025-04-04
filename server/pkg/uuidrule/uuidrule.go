package uuidrule

import (
	"errors"

	"github.com/google/uuid"
)

func Required(value interface{}) error {
	id, ok := value.(uuid.UUID)
	if !ok {
		return errors.New("invalid uuid type")
	}

	if id == uuid.Nil {
		return errors.New("uuid must not be empty")
	}

	return nil
}

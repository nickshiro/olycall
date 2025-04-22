package core

import (
	"errors"

	"olycall-server/internal/core/domain"
	"olycall-server/internal/core/ports/chatstore"
)

func (s Service) mapError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, chatstore.ErrUserNotFound) {
		return domain.ErrUserNotFound
	}

	return err
}

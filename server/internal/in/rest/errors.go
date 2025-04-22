package rest

import (
	"errors"
	"net/http"

	"olycall-server/internal/core/domain"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var (
	ErrMethodNotAllowed = errors.New("method not allowed")
	ErrNotFound         = errors.New("not found")
)

func UnwrapAll(err error) error {
	for {
		unwrapped := errors.Unwrap(err)
		if unwrapped == nil {
			return err
		}

		err = unwrapped
	}
}

func (c Controller) handleError(err error) handlerResponse {
	var validationErrors validation.Errors
	if errors.As(err, &validationErrors) {
		return handlerResponse{
			Body:    validationErrors,
			Status:  http.StatusBadRequest,
			Headers: nil,
		}
	}

	status := http.StatusInternalServerError

	switch {
	case errors.Is(err, domain.ErrInvalidToken):
		status = http.StatusUnauthorized

	case errors.Is(err, domain.ErrUserNotFound):
		status = http.StatusNotFound

	case errors.Is(err, domain.ErrOAuthStateNotFound):
		status = http.StatusBadRequest
	}

	return handlerResponse{
		Body:    err,
		Status:  status,
		Headers: nil,
	}
}

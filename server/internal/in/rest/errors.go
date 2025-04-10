package rest

import (
	"errors"
	"net/http"

	"olycall-server/internal/core/domain"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
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

func (c Controller) handleError(ctx echo.Context, err error) error {
	var validationErrors validation.Errors
	if errors.As(err, &validationErrors) {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: validationErrors})
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

	return ctx.JSON(status, ErrorResponse{Error: UnwrapAll(err).Error()})
}

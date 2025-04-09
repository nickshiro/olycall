package rest

import (
	"errors"
	"net/http"

	"olycall-server/internal/core/domain"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (c Controller) handleError(err error) handlerResponse {
	var validationErrors validation.Errors
	if errors.As(err, &validationErrors) {
		return handlerResponse{
			Body:    validationErrors,
			Status:  http.StatusBadRequest,
			IsError: true,
		}
	}

	status := http.StatusInternalServerError
	switch {
	case errors.Is(err, domain.ErrInvalidToken):
		status = http.StatusUnauthorized
		// case errors.Is(err, domain.ErrNotFound):
		// 	status = http.StatusNotFound
		// default:
		// var errConflict service.ConflictError
		// if errors.As(err, &errConflict) {
		// 	status = http.StatusConflict
		// }
	}

	return handlerResponse{
		Body:    err.Error(),
		Status:  status,
		IsError: true,
	}
}

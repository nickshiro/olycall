package rest

import (
	"log/slog"

	"olycall-server/internal/core/service/auth"
)

type Controller struct {
	authService *auth.AuthService
	logger      *slog.Logger
}

func NewController(
	authService *auth.AuthService,
	logger *slog.Logger,
) *Controller {
	return &Controller{
		authService: authService,
		logger:      logger,
	}
}

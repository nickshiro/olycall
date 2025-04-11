package rest

import (
	"log/slog"
	"olycall-server/internal/core"
)

type Controller struct {
	service *core.Service
	logger  *slog.Logger
}

func NewController(
	service *core.Service,
	logger *slog.Logger,
) *Controller {
	return &Controller{
		service: service,
		logger:  logger,
	}
}

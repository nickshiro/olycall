package rest

import (
	"context"
	"log/slog"

	"olycall-server/internal/service"

	"github.com/google/uuid"
)

type iService interface {
	GetGoogleLoginURL(context.Context, *service.GetGoogleLoginURLParams) (*service.GetGoogleLoginURLResponse, error)
	HandleGoogleCallback(context.Context, *service.HandleGoogleCallbackParams) (*service.HandleGoogleCallbackResponse, error) //nolint: lll
	GetUser(context.Context, uuid.UUID) (*service.User, error)
	GetMe(context.Context, string) (*service.User, error)
	RefreshTokens(context.Context, *service.RefreshTokensParams) (*service.RefreshTokensResponse, error)
}

type Controller struct {
	service iService
	logger  *slog.Logger
}

func NewController(service iService, logger *slog.Logger) *Controller {
	return &Controller{
		service: service,
		logger:  logger,
	}
}

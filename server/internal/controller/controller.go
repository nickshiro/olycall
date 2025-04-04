package controller

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/websocket"
)

type iService interface{}

type Controller struct {
	service  iService
	upgrader websocket.Upgrader
	logger   *slog.Logger
}

func NewController(service iService, logger *slog.Logger) *Controller {
	return &Controller{
		service: service,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		},
		logger: logger,
	}
}

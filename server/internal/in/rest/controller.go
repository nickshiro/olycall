package rest

import (
	"log/slog"
	"net/http"

	"olycall-server/internal/core"
	"olycall-server/internal/in/rest/gen"
	"olycall-server/internal/out/connectionstore/inmemory"
	"olycall-server/pkg/typesocket"

	"github.com/gorilla/websocket"
)

type Controller struct {
	service         *core.Service
	tserver         *typesocket.Server
	upgrader        *websocket.Upgrader
	connectionStore *inmemory.ConnectionStore
	logger          *slog.Logger
}

func NewController(
	service *core.Service,
	connectionStore *inmemory.ConnectionStore,
	logger *slog.Logger,
) *Controller {
	c := Controller{
		service:         service,
		connectionStore: connectionStore,
		upgrader: &websocket.Upgrader{
			// TODO: make opt-in
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		logger: logger,
	}

	c.tserver = c.getwsmux()

	return &c
}

func (c Controller) getwsmux() *typesocket.Server {
	s := typesocket.NewServer()

	gen.OnSendMessage(s, c.handleSendMessage)

	return s
}

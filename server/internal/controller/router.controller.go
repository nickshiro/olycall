package controller

import (
	"fmt"
	"net/http"

	// _ "github.com/xhhx-space/olycall-server/docs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

func methodNotAllowedHandler(_ *http.Request) handlerResponse {
	return handlerResponse{
		Body:    "the requested method is not allowed on this endpoint",
		Status:  http.StatusMethodNotAllowed,
		IsError: true,
	}
}

func (c Controller) GetMux() http.Handler {
	r := chi.NewRouter()
	r.MethodNotAllowed(c.makeHandler(methodNotAllowedHandler))

	r.Use(middleware.Recoverer)
	// r.Use(c.requestIDMw)
	// r.Use(c.requestLoggingMw)
	r.Use(cors.AllowAll().Handler)

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, "OK")
	})
	r.Route("/api", func(r chi.Router) {
		// r.Get("/join", c.handleWithError(c.join))
	})
	return r
}

func (c Controller) handleWithError(h func(w http.ResponseWriter, r *http.Request) error) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			c.logger.InfoContext(r.Context(), fmt.Sprintf("error in handler %s: %s", r.URL, err.Error()))
			w.Write([]byte(err.Error()))
		}
	}
}

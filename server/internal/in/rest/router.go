package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "olycall-server/docs"
)

func methodNotAllowedHandler(_ *http.Request) handlerResponse {
	return handlerResponse{
		Body:    ErrMethodNotAllowed,
		Status:  http.StatusMethodNotAllowed,
		Headers: nil,
	}
}

func notFoundHandler(_ *http.Request) handlerResponse {
	return handlerResponse{
		Body:    ErrNotFound,
		Status:  http.StatusNotFound,
		Headers: nil,
	}
}

func (c Controller) GetMux() http.Handler {
	r := chi.NewRouter()
	r.MethodNotAllowed(c.makeHandler(methodNotAllowedHandler))
	r.NotFound(c.makeHandler(notFoundHandler))

	r.Use(middleware.Recoverer)
	r.Use(middleware.RedirectSlashes)
	// r.Use(c.requestIDMw)
	// r.Use(c.requestLoggingMw)
	r.Use(cors.AllowAll().Handler)

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok")) //nolint: errcheck
		})
		r.Route("/auth", func(r chi.Router) {
			r.Get("/refresh", c.makeHandler(c.refresh))
			r.Get("/google", c.google)
			r.Get("/google-callback", c.googleCallback)
		})
		r.Route("/user", func(r chi.Router) {
			r.Route("/me", func(r chi.Router) {
				r.Use(c.accessTokenMw)
				r.Get("/", c.makeHandler(c.getMe))
				r.Put("/", c.makeHandler(c.putMe))
			})
			r.Route("/{user-id}", func(r chi.Router) {
				r.Use(c.userMw)
				r.Get("/", c.makeHandler(c.getUser))
			})
		})
		r.Route("/ws", func(r chi.Router) {
			r.Get("/primary", c.primaryWs)
		})
	})

	return r
}

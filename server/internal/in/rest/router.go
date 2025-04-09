package rest

import (
	"net/http"

	_ "olycall-server/docs"

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
	r.Use(c.requestIDMw)
	r.Use(c.requestLoggingMw)
	r.Use(cors.AllowAll().Handler)

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Route("/api", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/refresh", c.makeHandler(c.refresh))
			r.Get("/google", c.google)
			r.Get("/google-callback", c.makeHandler(c.googleCallback))
		})
		// r.Route("/user", func(r chi.Router) {
		// 	r.With(c.accessTokenMw).Get("/me", c.makeHandler(c.me))
		// 	r.Route("/{user-id}", func(r chi.Router) {
		// 		r.Use(c.userMw)
		// 		r.Get("/", c.makeHandler(c.getUser))
		// 	})
		// })
	})

	return r
}

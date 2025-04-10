package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "olycall-server/docs"
)

func (c Controller) GetMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	// e.Use(c.requestIDMw)
	// e.Use(c.requestLoggingMw)
	e.Use(middleware.CORS())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	api := e.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.GET("/refresh", c.refresh)
			auth.GET("/google", c.google)
			auth.GET("/google-callback", c.googleCallback)
		}
		user := api.Group("/user")
		{
			user.GET("/me", c.getMe, c.accessTokenMw)
			user.PUT("/me", c.putMe, c.accessTokenMw)
			user.GET("/:user-id", c.getUser, c.userMw)
		}
	}

	return e
}

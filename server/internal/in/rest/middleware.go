package rest

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type AccessTokenCtx struct {
	echo.Context
}

func (c *AccessTokenCtx) SetAccessToken(value string) {
	c.Set("access-token", value)
}

func (c *AccessTokenCtx) GetAccessToken() string {
	value, _ := c.Get("access-token").(string)
	return value
}

type UserIDCtx struct {
	echo.Context
}

func (c *UserIDCtx) SetUserID(value uuid.UUID) {
	c.Set("user-id", value)
}

func (c *UserIDCtx) GetUserID() uuid.UUID {
	value, _ := c.Get("user-id").(uuid.UUID)
	return value
}

// func (c Controller) requestLoggingMw(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(ctx echo.Context) error {
// 		c.logger.Info("request",
// 			"method", ctx.Request().Method,
// 			"url", ctx.Request().URL.String(),
// 			"remote_addr", ctx.Request().RemoteAddr,
// 			"headers", ctx.Request().Header,
// 			"body", ctx.Request().Body,
// 		)
// 		return next(ctx)
// 	}
// }

func (c Controller) userMw(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		uc := UserIDCtx{ctx}
		userID := ctx.Param("user-id")
		userUUID, err := uuid.Parse(userID)
		if err != nil {
			return ctx.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
		}

		uc.SetUserID(userUUID)
		return next(uc)
	}
}

func (c Controller) accessTokenMw(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ac := AccessTokenCtx{ctx}

		accessToken, err := c.getAccessToken(ctx)
		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, ErrorResponse{Error: err.Error()})
		}

		ac.SetAccessToken(accessToken)
		return next(ac)
	}
}

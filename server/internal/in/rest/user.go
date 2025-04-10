package rest

import (
	"net/http"

	"olycall-server/internal/core"
	"olycall-server/internal/core/domain"

	"github.com/labstack/echo/v4"
)

// @Summary	Get user by ID
// @Tags	Users
// @Param	user-id	path		string	true	"User ID (UUID)"	format(uuid)
// @Success	200		{object}	SuccessResponse[domain.User]
// @Failure	401		{object}	ErrorResponse
// @Failure	404		{object}	ErrorResponse
// @Failure	500		{object}	ErrorResponse
// @Router	/user/{user-id} [get]
func (c Controller) getUser(echoCtx echo.Context) error {
	ctx := echoCtx.(UserIDCtx) // nolint: errcheck

	getUserResp, err := c.service.GetUser(ctx.Request().Context(), ctx.GetUserID())
	if err != nil {
		return c.handleError(ctx, err)
	}

	return ctx.JSON(
		http.StatusOK,
		SuccessResponse[*domain.User]{Data: getUserResp},
	)
}

// @Summary	Get current user
// @Tags	Users
// @Success	200	{object}	SuccessResponse[domain.User]
// @Failure	401	{object}	ErrorResponse
// @Failure	500	{object}	ErrorResponse
// @Router	/user/getMe [get]
func (c Controller) getMe(echoCtx echo.Context) error {
	ctx := echoCtx.(AccessTokenCtx) // nolint: errcheck

	getUserResp, err := c.service.GetMe(ctx.Request().Context(), ctx.GetAccessToken())
	if err != nil {
		return c.handleError(ctx, err)
	}

	return ctx.JSON(
		http.StatusOK,
		SuccessResponse[*domain.User]{Data: getUserResp},
	)
}

type putMeBody struct {
	Username string `json:"username" validate:"required"`
} // @name PutMeBody

// @Summary	Update current user
// @Tags	Users
// @Param	body	body	putMeBody	true	"Request body"
// @Success	200	{object}	SuccessResponse[domain.User]
// @Failure	400	{object}	ErrorResponse
// @Failure	401	{object}	ErrorResponse
// @Failure	500	{object}	ErrorResponse
// @Router	/user/me [put]
func (c Controller) putMe(echoCtx echo.Context) error {
	ctx := echoCtx.(AccessTokenCtx) // nolint: errcheck
	var body putMeBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}

	getUserResp, err := c.service.UpdateMe(ctx.Request().Context(), &core.UpdateMeParams{
		AccessToken: ctx.GetAccessToken(),
		Username:    "",
	})
	if err != nil {
		return c.handleError(ctx, err)
	}

	return ctx.JSON(
		http.StatusOK,
		SuccessResponse[*domain.User]{Data: getUserResp},
	)
}

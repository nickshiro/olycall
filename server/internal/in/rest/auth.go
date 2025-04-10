package rest

import (
	"net/http"

	"olycall-server/internal/core"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// @Summary	Refreshes auth token pair
// @Tags	Auth
// @Param	RefresnTokenCookie	header		string		true	"Refresh token cookie"
// @Header	200					{string}	Set-Cookie	"access_token=<token>; HttpOnly; Secure; SameSite=Strict"
// @Header	200					{string}	Set-Cookie	"refresh_token=<token>; HttpOnly; Secure; SameSite=Strict"
// @Success	200					{object}	SuccessResponse[string]
// @Failure	400					{object}	ErrorResponse
// @Failure	500					{object}	ErrorResponse
// @Router	/auth/refresh [get]
func (c Controller) refresh(ctx echo.Context) error {
	refreshToken, err := c.getRefreshToken(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
	}

	refreshTokensResp, err := c.service.RefreshTokens(ctx.Request().Context(), &core.RefreshTokensParams{
		RefreshToken: refreshToken,
	})
	if err != nil {
		return c.handleError(ctx, err)
	}

	c.addAuthTokensPairToHeader(ctx, refreshTokensResp.AccessToken, refreshTokensResp.RefreshToken)
	return ctx.JSON(http.StatusOK, SuccessResponse[string]{Data: "ok"})
}

// @Summary	Endpoint that redirects to google oauth page
// @Tags	Auth
// @Param	redirect_uri	query		string	true	"URI to redirect to after successful login"
// @Success	302				{string}	string	"Redirect to Google's OAuth2 consent screen"
// @Failure	400				{object}	ErrorResponse
// @Failure	500				{object}	ErrorResponse
// @Router	/auth/google [get]
func (c Controller) google(ctx echo.Context) error {
	params := struct {
		RedirectURI string `query:"redirect_uri"`
	}{}
	if err := ctx.Bind(&params); err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
	}

	resp, err := c.service.GetGoogleLoginURL(ctx.Request().Context(), &core.GetGoogleLoginURLParams{
		RedirectURI: params.RedirectURI,
	})
	if err != nil {
		return c.handleError(ctx, err)
	}

	return ctx.Redirect(http.StatusFound, resp.URL)
}

func (c Controller) googleCallback(ctx echo.Context) error {
	params := struct {
		Code  string    `query:"code" json:"code"`
		State uuid.UUID `query:"state" json:"state"`
	}{}
	if err := ctx.Bind(&params); err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
	}

	resp, err := c.service.HandleGoogleCallback(ctx.Request().Context(), &core.HandleGoogleCallbackParams{
		Code:         params.Code,
		OAuthStateID: params.State,
	})
	if err != nil {
		return c.handleError(ctx, err)
	}

	c.addAuthTokensPairToHeader(ctx, resp.AccessToken, resp.RefreshToken)
	return ctx.Redirect(http.StatusMovedPermanently, resp.RedirectURI)
}

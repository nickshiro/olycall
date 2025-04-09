package rest

import (
	"net/http"

	"olycall-server/internal/core/service/auth"

	"github.com/google/uuid"
)

// @Summary	Refreshes auth token pair
// @Tags	Auth
// @Param	RefresnTokenCookie	header		string		true	"Refresh token cookie (e.g., refresh_token=<token>)"
// @Header	200					{string}	Set-Cookie	"access_token=<token>; HttpOnly; Secure; SameSite=Strict"
// @Header	200					{string}	Set-Cookie	"refresh_token=<token>; HttpOnly; Secure; SameSite=Strict"
// @Success	200					{object}	SuccessResponse[string]
// @Failure	400					{object}	ErrorResponse
// @Failure	500					{object}	ErrorResponse
// @Router	/auth/refresh [get]
func (c Controller) refresh(r *http.Request) handlerResponse {
	refreshToken, err := c.getRefreshToken(r)
	if err != nil {
		return handlerResponse{
			Body:    err.Error(),
			Status:  http.StatusBadRequest,
			IsError: true,
		}
	}

	refreshTokensResp, err := c.authService.RefreshTokens(r.Context(), &auth.RefreshTokensParams{
		RefreshToken: refreshToken,
	})
	if err != nil {
		return c.handleError(err)
	}

	h := c.addAuthTokensPairToHeader(&http.Header{}, refreshTokensResp.AccessToken, refreshTokensResp.RefreshToken)

	return handlerResponse{
		Body:    "ok",
		Status:  http.StatusOK,
		IsError: false,
		Headers: h,
	}
}

// @Summary	Endpoint that redirects to google oauth page
// @Tags	Auth
// @Param	redirect_uri	query		string	true	"URI to redirect to after successful login"
// @Success	302				{string}	string	"Redirect to Google's OAuth2 consent screen"
// @Failure	400				{object}	ErrorResponse
// @Failure	500				{object}	ErrorResponse
// @Router	/auth/google [get]
func (c Controller) google(w http.ResponseWriter, r *http.Request) {
	redirectURI := r.URL.Query().Get("redirect_uri")
	resp, err := c.authService.GetGoogleLoginURL(r.Context(), &auth.GetGoogleLoginURLParams{
		RedirectURI: redirectURI,
	})
	if err != nil {
		c.processHandlerResponse(r.Context(), w, c.handleError(err))
		return
	}

	http.Redirect(w, r, resp.URL, http.StatusFound)
}

// @Summary	Get auth tokens with google oauth token
// @Tags	Auth
// @Param	RefresnTokenCookie	header		string		true	"Refresh token cookie (e.g., refresh_token=<token>)"
// @Header	200					{string}	Set-Cookie	"access_token=<token>; HttpOnly; Secure; SameSite=Strict"
// @Header	200					{string}	Set-Cookie	"refresh_token=<token>; HttpOnly; Secure; SameSite=Strict"
// @Success	200					{object}	SuccessResponse[string]
// @Failure	400					{object}	ErrorResponse
// @Failure	500					{object}	ErrorResponse
// @Router	/auth/google-callback [get]
func (c Controller) googleCallback(r *http.Request) handlerResponse {
	state := r.URL.Query().Get("state")

	stateUUID, err := uuid.Parse(state)
	if err != nil {
		return handlerResponse{
			Body:    err,
			Status:  http.StatusBadRequest,
			IsError: true,
		}
	}

	resp, err := c.authService.HandleGoogleCallback(r.Context(), &auth.HandleGoogleCallbackParams{
		Code:         r.URL.Query().Get("code"),
		OAuthStateID: stateUUID,
	})
	if err != nil {
		return c.handleError(err)
	}

	h := c.addAuthTokensPairToHeader(&http.Header{}, resp.AccessToken, resp.RefreshToken)

	return handlerResponse{
		Body:    "ok",
		Status:  http.StatusOK,
		Headers: h,
		IsError: false,
	}
}

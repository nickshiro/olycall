package rest

import (
	"net/http"

	"olycall-server/internal/service"
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

	refreshTokensResp, err := c.service.RefreshTokens(r.Context(), &service.RefreshTokensParams{
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
// @Success 302				{string}	string	"Redirect to Google's OAuth2 consent screen"
// @Failure 400				{object}	ErrorResponse
// @Failure 500				{object}	ErrorResponse
// @Router	/auth/login [get]
func (c Controller) login(w http.ResponseWriter, r *http.Request) {
	redirectURI := r.URL.Query().Get("redirect_uri")
	resp, err := c.service.GetGoogleLoginURL(r.Context(), &service.GetGoogleLoginURLParams{
		RedirectURI: redirectURI,
	})
	if err != nil {
		c.processHandlerResponse(r.Context(), w, c.handleError(err))
		return
	}

	http.Redirect(w, r, resp.URL, http.StatusFound)
}

func (c Controller) googleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	resp, err := c.service.HandleGoogleCallback(r.Context(), &service.HandleGoogleCallbackParams{
		Code:  code,
		State: state,
	})
	if err != nil {
		c.processHandlerResponse(r.Context(), w, c.handleError(err))
		return
	}

	h := w.Header()
	c.addAuthTokensPairToHeader(&h, resp.AccessToken, resp.RefreshToken)

	http.Redirect(w, r, resp.RedirectURI, http.StatusFound)
}

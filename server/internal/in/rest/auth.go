package rest

import (
	"net/http"
	"olycall-server/internal/core"

	"github.com/google/uuid"
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
func (c Controller) refresh(r *http.Request) handlerResponse {
	refreshToken := c.getRefreshToken(r)

	refreshTokensResp, err := c.service.RefreshTokens(r.Context(), &core.RefreshTokensParams{
		RefreshToken: refreshToken,
	})
	if err != nil {
		return c.handleError(err)
	}

	h := &http.Header{}
	c.addAuthTokensPairToHeader(h, refreshTokensResp.AccessToken, refreshTokensResp.RefreshToken)

	return handlerResponse{
		Body:    "ok",
		Status:  http.StatusOK,
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

	resp, err := c.service.GetGoogleLoginURL(r.Context(), &core.GetGoogleLoginURLParams{
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

	stateUUID, err := uuid.Parse(state)
	if err != nil {
		c.processHandlerResponse(r.Context(), w, handlerResponse{
			Body:    err,
			Status:  http.StatusBadRequest,
			Headers: nil,
		})

		return
	}

	resp, err := c.service.HandleGoogleCallback(r.Context(), &core.HandleGoogleCallbackParams{
		Code:         code,
		OAuthStateID: stateUUID,
	})
	if err != nil {
		c.processHandlerResponse(r.Context(), w, c.handleError(err))

		return
	}

	h := w.Header()
	c.addAuthTokensPairToHeader(&h, resp.AccessToken, resp.RefreshToken)

	http.Redirect(w, r, resp.RedirectURI, http.StatusMovedPermanently)
}

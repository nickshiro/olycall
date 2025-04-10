package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	accessTokenCookieName  = "access_token"
	refreshTokenCookieName = "refresh_token"
)

func (c Controller) addAuthTokensPairToHeader(ctx echo.Context, accessToken string, refreshToken string) {
	ctx.SetCookie(&http.Cookie{
		Name:     accessTokenCookieName,
		Value:    accessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	ctx.SetCookie(&http.Cookie{
		Name:     refreshTokenCookieName,
		Value:    refreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
}

func (c Controller) getAccessToken(ctx echo.Context) (string, error) {
	cookie, err := ctx.Cookie(accessTokenCookieName)
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}

func (c Controller) getRefreshToken(ctx echo.Context) (string, error) {
	cookie, err := ctx.Cookie(refreshTokenCookieName)
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}

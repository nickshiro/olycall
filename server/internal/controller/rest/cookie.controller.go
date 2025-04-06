package rest

import (
	"fmt"
	"net/http"
)

const (
	accessTokenCookieName  = "access_token"
	refreshTokenCookieName = "refresh_token"
)

func (c Controller) addAuthTokensPairToHeader(h *http.Header, accessToken string, refreshToken string) *http.Header {
	accessTokenCookie := &http.Cookie{
		Name:     accessTokenCookieName,
		Value:    accessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	refreshTokenCookie := &http.Cookie{
		Name:     refreshTokenCookieName,
		Value:    refreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	// if h.isDebug {
	// 	cookie.HttpOnly = false
	// 	cookie.Secure = false
	// 	cookie.SameSite = http.SameSiteNoneMode
	// } else {
	// 	cookie.HttpOnly = true
	// 	cookie.Secure = true
	// 	cookie.SameSite = http.SameSiteStrictMode
	// }

	h.Add("Set-Cookie", accessTokenCookie.String())
	h.Add("Set-Cookie", refreshTokenCookie.String())
	return h
}

func (c Controller) getRefreshToken(r *http.Request) (string, error) {
	tokenCookie, err := r.Cookie(refreshTokenCookieName)
	if err != nil {
		return "", fmt.Errorf("failed to get refresh token cookie: %w", err)
	}

	return tokenCookie.Value, nil
}

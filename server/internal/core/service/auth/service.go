package auth

import (
	"olycall-server/internal/core/ports/googleoauthprovider"
	"olycall-server/internal/core/ports/oauthstatestore"
	"olycall-server/internal/core/ports/userstore"
)

type AuthService struct {
	userStore           userstore.UserStore
	oauthStateStore     oauthstatestore.OAuthStateStore
	googleOAuthProvider googleoauthprovider.GoogleOAuthProvider
	secret              string
}

func NewService(
	userStore userstore.UserStore,
	oauthStateStore oauthstatestore.OAuthStateStore,
	googleOAuthProvider googleoauthprovider.GoogleOAuthProvider,
	secret string,
) *AuthService {
	return &AuthService{
		userStore:           userStore,
		oauthStateStore:     oauthStateStore,
		googleOAuthProvider: googleOAuthProvider,
		secret:              secret,
	}
}

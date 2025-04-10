package core

import (
	"olycall-server/internal/core/ports/googleoauthprovider"
	"olycall-server/internal/core/ports/oauthstatestore"
	"olycall-server/internal/core/ports/userstore"
)

type Service struct {
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
) *Service {
	return &Service{
		userStore:           userStore,
		oauthStateStore:     oauthStateStore,
		googleOAuthProvider: googleOAuthProvider,
		secret:              secret,
	}
}

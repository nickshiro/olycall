package core

import (
	"olycall-server/internal/core/ports/chatstore"
	"olycall-server/internal/core/ports/connectionstore"
	"olycall-server/internal/core/ports/googleoauthprovider"
	"olycall-server/internal/core/ports/notificationsprovider"
	"olycall-server/internal/core/ports/oauthstatestore"
	"olycall-server/internal/core/ports/userstore"
)

type Service struct {
	userStore             userstore.UserStore
	chatStore             chatstore.ChatStore
	oauthStateStore       oauthstatestore.OAuthStateStore
	googleOAuthProvider   googleoauthprovider.GoogleOAuthProvider
	notificationsProvider notificationsprovider.NotificationsProvider
	connectionStore       connectionstore.ConnectionStore
	secret                string
}

func NewService(
	userStore userstore.UserStore,
	chatStore chatstore.ChatStore,
	oauthStateStore oauthstatestore.OAuthStateStore,
	googleOAuthProvider googleoauthprovider.GoogleOAuthProvider,
	notificationsProvider notificationsprovider.NotificationsProvider,
	connectionStore connectionstore.ConnectionStore,
	secret string,
) *Service {
	return &Service{
		userStore:             userStore,
		chatStore:             chatStore,
		oauthStateStore:       oauthStateStore,
		googleOAuthProvider:   googleOAuthProvider,
		notificationsProvider: notificationsProvider,
		connectionStore:       connectionStore,
		secret:                secret,
	}
}

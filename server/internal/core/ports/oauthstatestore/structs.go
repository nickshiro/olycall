package oauthstatestore

import (
	"time"

	"github.com/google/uuid"
)

type OAuthState struct {
	RedirectURI string `json:"redirect_uri"`
}

type CreateOAuthStateParams struct {
	ID          uuid.UUID     `json:"id"`
	RedirectURI string        `json:"redirect_uri"`
	TTL         time.Duration `json:"ttl"`
}

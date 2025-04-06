package cache

import (
	"time"

	"github.com/google/uuid"
)

type State struct {
	RedirectURI string `json:"redirect_uri"`
}

type SetStateParams struct {
	ID          uuid.UUID     `json:"id"`
	RedirectURI string        `json:"redirect_uri"`
	TTL         time.Duration `json:"ttl"`
}

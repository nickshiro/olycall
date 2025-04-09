package redis

import (
	"github.com/redis/go-redis/v9"
)

type OAuthStateStore struct {
	rc *redis.Client
}

func NewOAuthStateStore(rc *redis.Client) *OAuthStateStore {
	return &OAuthStateStore{
		rc: rc,
	}
}

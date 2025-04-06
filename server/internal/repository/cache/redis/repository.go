package redis

import (
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	rc *redis.Client
}

func NewRepo(rc *redis.Client) *Repository {
	return &Repository{
		rc: rc,
	}
}

package redis

import (
	"context"
	"fmt"

	r "github.com/redis/go-redis/v9"
)

type Config struct {
	Host     string
	Port     int
	Password string
}

func NewRedisClient(cfg *Config) (*r.Client, error) {
	client := r.NewClient(&r.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       0,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("failed to ping: %w", err)
	}

	return client, nil
}

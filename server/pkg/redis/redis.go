package redis

import (
	"context"
	"fmt"

	rc "github.com/redis/go-redis/v9"
)

func NewClient(
	ctx context.Context,
	host string,
	port int,
	password string,
	db int,
) (*rc.Client, error) {
	client := rc.NewClient(&rc.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
		DB:       db,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to ping: %w", err)
	}

	return client, nil
}

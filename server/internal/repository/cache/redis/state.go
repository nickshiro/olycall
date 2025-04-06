package redis

import (
	"context"
	"errors"

	"olycall-server/internal/repository/cache"

	"github.com/redis/go-redis/v9"
)

func (r Repository) getStateKey(stateID string) string {
	return "auth-state:" + stateID
}

func (r Repository) SetState(ctx context.Context, params *cache.SetStateParams) error {
	return r.rc.Set(ctx, r.getStateKey(params.ID.String()), params.RedirectURI, params.TTL).Err()
}

func (r Repository) GetState(ctx context.Context, id string) (*cache.State, error) {
	stateKey := r.getStateKey(id)
	state, err := r.rc.Get(ctx, stateKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}

		return nil, err
	}

	return &cache.State{
		RedirectURI: state,
	}, nil
}

func (r Repository) RemoveState(ctx context.Context, id string) (bool, error) {
	res, err := r.rc.Del(ctx, r.getStateKey(id)).Result()
	if err != nil {
		return false, err
	}

	if res == 0 {
		return false, nil
	}

	return true, nil
}

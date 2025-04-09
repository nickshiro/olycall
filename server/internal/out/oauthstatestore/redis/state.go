package redis

import (
	"context"
	"errors"

	"olycall-server/internal/core/ports/oauthstatestore"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func (r OAuthStateStore) getStateKey(stateID string) string {
	return "oauth-state:" + stateID
}

func (r OAuthStateStore) CreateOAuthState(ctx context.Context, params *oauthstatestore.CreateOAuthStateParams) error {
	return r.rc.Set(ctx, r.getStateKey(params.ID.String()), params.RedirectURI, params.TTL).Err()
}

func (r OAuthStateStore) GetOAuthState(ctx context.Context, id uuid.UUID) (*oauthstatestore.OAuthState, error) {
	stateKey := r.getStateKey(id.String())
	state, err := r.rc.Get(ctx, stateKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}

		return nil, err
	}

	return &oauthstatestore.OAuthState{
		RedirectURI: state,
	}, nil
}

func (r OAuthStateStore) DeleteOAuthState(ctx context.Context, id uuid.UUID) (bool, error) {
	res, err := r.rc.Del(ctx, r.getStateKey(id.String())).Result()
	if err != nil {
		return false, err
	}

	if res == 0 {
		return false, nil
	}

	return true, nil
}

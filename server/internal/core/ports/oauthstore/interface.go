package oauthstore

import "context"

type OAuthStore interface {
	CreateOauthIdentity(ctx context.Context, params *CreateOauthIdentityParams) error
}

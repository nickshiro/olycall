package googleoauthprovider

import "context"

type GoogleOAuthProvider interface {
	GetUserInfo(ctx context.Context, code string) (UserInfo, error)
	GetLoginURL(ctx context.Context, state string) string
}

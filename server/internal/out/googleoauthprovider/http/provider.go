package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"olycall-server/internal/core/ports/googleoauthprovider"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleOAuthProvider struct {
	config *oauth2.Config
}

func NewGoogleOAuthProvider(
	clientID string,
	clientSecret string,
	redirectURL string,
) *GoogleOAuthProvider {
	config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes: []string{
			"openid",
			"profile",
			"email",
		},
		Endpoint: google.Endpoint,
	}

	return &GoogleOAuthProvider{config: config}
}

func (p GoogleOAuthProvider) GetUserInfo(
	ctx context.Context,
	code string,
) (googleoauthprovider.UserInfo, error) {
	token, err := p.exchangeCodeForToken(ctx, code)
	if err != nil {
		return googleoauthprovider.UserInfo{}, err
	}

	userInfo, err := p.getUserInfoFromToken(ctx, token.AccessToken)
	if err != nil {
		return googleoauthprovider.UserInfo{}, err
	}

	return userInfo, nil
}

func (p GoogleOAuthProvider) exchangeCodeForToken(
	ctx context.Context,
	code string,
) (*oauth2.Token, error) {
	options := []oauth2.AuthCodeOption{}

	token, err := p.config.Exchange(ctx, code, options...)
	if err != nil {
		return nil, fmt.Errorf("token exchange failed: %w", err)
	}

	return token, nil
}

// TODO: move error out
var ErrGetUserInfo = errors.New("failed to get user info")

func (p GoogleOAuthProvider) getUserInfoFromToken(
	ctx context.Context,
	accessToken string,
) (googleoauthprovider.UserInfo, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://www.googleapis.com/oauth2/v3/userinfo", nil)
	if err != nil {
		return googleoauthprovider.UserInfo{}, fmt.Errorf("create user info request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Do(req)
	if err != nil {
		return googleoauthprovider.UserInfo{}, fmt.Errorf("send user info request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)

		return googleoauthprovider.UserInfo{}, fmt.Errorf(
			"%w: status: %d, body: %s",
			ErrGetUserInfo,
			resp.StatusCode,
			body,
		)
	}

	var userInfo googleoauthprovider.UserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return googleoauthprovider.UserInfo{}, fmt.Errorf("decode user info: %w", err)
	}

	return userInfo, nil
}

func (p GoogleOAuthProvider) GetLoginURL(ctx context.Context, state string) string {
	return p.config.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

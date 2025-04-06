package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"olycall-server/internal/repository/cache"
	"olycall-server/internal/repository/domain"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

type GetGoogleLoginURLParams struct {
	RedirectURI string `json:"redirect_uri"`
}

type GetGoogleLoginURLResponse struct {
	URL string `json:"url"`
}

const stateTTL = time.Minute * 5

func (s Service) GetGoogleLoginURL(
	ctx context.Context,
	params *GetGoogleLoginURLParams,
) (*GetGoogleLoginURLResponse, error) {
	if err := validation.ValidateStructWithContext(ctx, params,
		validation.Field(&params.RedirectURI, validation.Required),
	); err != nil {
		return nil, wrapInvalidParamsErr(err)
	}

	stateID := uuid.New()
	if err := s.cacheRepo.SetState(ctx, &cache.SetStateParams{
		ID:          stateID,
		RedirectURI: params.RedirectURI,
		TTL:         stateTTL,
	}); err != nil {
		return nil, fmt.Errorf("failed to set state: %w", err)
	}

	url := s.googleOauth2Config.AuthCodeURL(stateID.String(), oauth2.AccessTypeOffline)
	return &GetGoogleLoginURLResponse{
		URL: url,
	}, nil
}

type HandleGoogleCallbackParams struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

type HandleGoogleCallbackResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	RedirectURI  string `json:"redirect_uri"`
}

func (s Service) HandleGoogleCallback(
	ctx context.Context,
	params *HandleGoogleCallbackParams,
) (*HandleGoogleCallbackResponse, error) {
	if err := validation.ValidateStructWithContext(ctx, params,
		validation.Field(&params.Code, validation.Required),
		validation.Field(&params.State, validation.Required),
	); err != nil {
		return nil, wrapInvalidParamsErr(err)
	}

	state, err := s.cacheRepo.GetState(ctx, params.State)
	if err != nil {
		return nil, fmt.Errorf("failed to get state: %w", err)
	}

	ok, err := s.cacheRepo.RemoveState(ctx, params.State)
	if err != nil {
		return nil, fmt.Errorf("failed to remove state: %w", err)
	}
	if !ok {
		return nil, errors.New("state not found")
	}

	token, err := s.googleOauth2Config.Exchange(ctx, params.Code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %w", err)
	}

	client := s.googleOauth2Config.Client(ctx, token)
	userInfoResp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}
	defer userInfoResp.Body.Close()

	body, _ := io.ReadAll(userInfoResp.Body)
	var userInfo struct {
		ID         string  `json:"id"`
		Email      string  `json:"email"`
		Name       string  `json:"name"`
		GivenName  string  `json:"given_name"`
		FamilyName *string `json:"family_name"`
		Picture    string  `json:"picture"`
		Locale     string  `json:"locale"`
	}
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return nil, fmt.Errorf("body parsing failed: %w", err)
	}

	user, err := s.domainRepo.GetUserByEmail(ctx, userInfo.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to select user by email: %w", err)
	}

	var userID uuid.UUID
	if user != nil {
		ok, err := s.domainRepo.UpdateUser(ctx, &domain.UpdateUserParams{
			ID: userID,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to update user: %w", err)
		}
		if !ok {
			return nil, errors.New("failed to update user: user not found")
		}

		userID = user.ID
	} else {
		userID = uuid.New()
		if err := s.domainRepo.CreateUser(ctx, &domain.CreateUserParams{
			ID:    userID,
			Email: userInfo.Email,
			// Username: *userInfo.FamilyName,
		}); err != nil {
			return nil, fmt.Errorf("failed to create user: %w", err)
		}
	}

	tokenPair := s.generateJWT(userID.String())

	return &HandleGoogleCallbackResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
		RedirectURI:  state.RedirectURI,
	}, nil
}

type RefreshTokensParams struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokensResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (s Service) RefreshTokens(
	_ context.Context,
	params *RefreshTokensParams,
) (*RefreshTokensResponse, error) {
	userID, err := s.getUserIDFromAccessToken(params.RefreshToken)
	if err != nil {
		return nil, ErrInvalidToken
	}

	tokenPair := s.generateJWT(userID.String())

	return &RefreshTokensResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}, nil
}

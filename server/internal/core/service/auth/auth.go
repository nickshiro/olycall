package auth

import (
	"context"
	"fmt"
	"time"

	"olycall-server/internal/core/domain"
	"olycall-server/internal/core/ports/oauthstatestore"
	"olycall-server/internal/core/ports/userstore"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

func wrapInvalidParamsErr(err error) error {
	return fmt.Errorf("invalid params: %w", err)
}

type GetGoogleLoginURLParams struct {
	RedirectURI string
}

type GetGoogleLoginURLResponse struct {
	URL string
}

const stateTTL = time.Minute * 5

func (s AuthService) GetGoogleLoginURL(
	ctx context.Context,
	params *GetGoogleLoginURLParams,
) (*GetGoogleLoginURLResponse, error) {
	if err := validation.ValidateStructWithContext(ctx, params,
		validation.Field(&params.RedirectURI, validation.Required),
	); err != nil {
		return nil, wrapInvalidParamsErr(err)
	}

	stateID := uuid.New()
	if err := s.oauthStateStore.CreateOAuthState(ctx, &oauthstatestore.CreateOAuthStateParams{
		ID:          stateID,
		RedirectURI: params.RedirectURI,
		TTL:         stateTTL,
	}); err != nil {
		return nil, fmt.Errorf("create oauth state: %w", err)
	}

	url := s.googleOAuthProvider.GetLoginURL(ctx, stateID.String())
	return &GetGoogleLoginURLResponse{
		URL: url,
	}, nil
}

type HandleGoogleCallbackParams struct {
	Code         string
	OAuthStateID uuid.UUID
}

type HandleGoogleCallbackResponse struct {
	AccessToken  string
	RefreshToken string
	RedirectURI  string
}

func (s AuthService) HandleGoogleCallback(
	ctx context.Context,
	params *HandleGoogleCallbackParams,
) (*HandleGoogleCallbackResponse, error) {
	if err := validation.ValidateStructWithContext(ctx, params,
		validation.Field(&params.Code, validation.Required),
		validation.Field(&params.OAuthStateID, validation.Required),
	); err != nil {
		return nil, wrapInvalidParamsErr(err)
	}

	oauthState, err := s.oauthStateStore.GetOAuthState(ctx, params.OAuthStateID)
	if err != nil {
		return nil, fmt.Errorf("get oauth state: %w", err)
	}
	if oauthState == nil {
		return nil, domain.ErrOAuthStateNotFound
	}

	userInfo, err := s.googleOAuthProvider.GetUserInfo(ctx, params.Code)
	if err != nil {
		return nil, fmt.Errorf("get user info: %w", err)
	}

	if _, err := s.oauthStateStore.DeleteOAuthState(ctx, params.OAuthStateID); err != nil {
		return nil, fmt.Errorf("delete oauth state: %w", err)
	}

	user, err := s.userStore.GetUserByEmail(ctx, userInfo.Email)
	if err != nil {
		return nil, fmt.Errorf("get user by email: %w", err)
	}

	var userID uuid.UUID
	if user != nil {
		found, err := s.userStore.UpdateUser(ctx, &userstore.UpdateUserParams{
			ID: userID,
		})
		if err != nil {
			return nil, fmt.Errorf("update user: %w", err)
		}
		if !found {
			return nil, domain.ErrUserNotFound
		}

		userID = user.ID
	} else {
		userID = uuid.New()
		if err := s.userStore.CreateUser(ctx, &userstore.CreateUserParams{
			ID:    userID,
			Email: userInfo.Email,
			// Username: *userInfo.FamilyName,
		}); err != nil {
			return nil, fmt.Errorf("create user: %w", err)
		}
	}

	tokenPair := s.generateJWT(userID.String())

	return &HandleGoogleCallbackResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}, nil
}

type RefreshTokensParams struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokensResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (s AuthService) RefreshTokens(
	_ context.Context,
	params *RefreshTokensParams,
) (*RefreshTokensResponse, error) {
	userID, err := s.getUserIDFromAccessToken(params.RefreshToken)
	if err != nil {
		return nil, err
	}

	tokenPair := s.generateJWT(userID.String())

	return &RefreshTokensResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}, nil
}

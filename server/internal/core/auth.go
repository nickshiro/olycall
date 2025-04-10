package core

import (
	"context"
	"fmt"
	"math/rand/v2"
	"strings"
	"time"
	"unicode"

	"olycall-server/internal/core/domain"
	"olycall-server/internal/core/ports/oauthstatestore"
	"olycall-server/internal/core/ports/userstore"

	"github.com/brianvoe/gofakeit/v7"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

func randomName() string {
	n := [7]string{
		gofakeit.MinecraftAnimal(),
		gofakeit.MinecraftFood(),
		gofakeit.MinecraftMobBoss(),
		gofakeit.MinecraftMobHostile(),
		gofakeit.MinecraftMobNeutral(),
		gofakeit.MinecraftMobPassive(),
		"User",
	}

	i := rand.IntN(len(n))
	for strings.Contains(n[i], " ") {
		i = (i + 1) % len(n)
	}

	u := fmt.Sprintf(
		"%s%s%d",
		capitalizeFirst(gofakeit.AdjectiveDescriptive()),
		capitalizeFirst(n[i]),
		rand.IntN(9000)+1000,
	)
	return u
}

func capitalizeFirst(s string) string {
	if len(s) == 0 {
		return ""
	}
	firstChar := []rune(s)[0]
	upperFirst := string(unicode.ToUpper(firstChar))
	restOfString := string([]rune(s)[1:])
	return upperFirst + restOfString
}

func wrapInvalidParamsErr(err error) error {
	return fmt.Errorf("invalid params: %w", err)
}

type GetGoogleLoginURLParams struct {
	RedirectURI string
}

type GetGoogleLoginURLResp struct {
	URL string
}

const stateTTL = time.Minute * 5

func (s Service) GetGoogleLoginURL(
	ctx context.Context,
	params *GetGoogleLoginURLParams,
) (*GetGoogleLoginURLResp, error) {
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
	return &GetGoogleLoginURLResp{
		URL: url,
	}, nil
}

type HandleGoogleCallbackParams struct {
	Code         string
	OAuthStateID uuid.UUID
}

type HandleGoogleCallbackResp struct {
	AccessToken  string
	RefreshToken string
	RedirectURI  string
}

func (s Service) HandleGoogleCallback(
	ctx context.Context,
	params *HandleGoogleCallbackParams,
) (*HandleGoogleCallbackResp, error) {
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
	if user == nil {
		username := randomName()
		userID = uuid.New()
		now := time.Now()

		if err := s.userStore.CreateUser(ctx, &userstore.CreateUserParams{
			ID:        userID,
			Email:     userInfo.Email,
			Username:  username,
			CreatedAt: now,
		}); err != nil {
			return nil, fmt.Errorf("create user: %w", err)
		}
	}

	tokenPair := s.generateJWT(userID.String())

	return &HandleGoogleCallbackResp{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
		RedirectURI:  oauthState.RedirectURI,
	}, nil
}

type RefreshTokensParams struct {
	RefreshToken string
}

type RefreshTokensResp struct {
	AccessToken  string
	RefreshToken string
}

func (s Service) RefreshTokens(
	_ context.Context,
	params *RefreshTokensParams,
) (*RefreshTokensResp, error) {
	userID, err := s.getUserIDFromAccessToken(params.RefreshToken)
	if err != nil {
		return nil, err
	}

	tokenPair := s.generateJWT(userID.String())

	return &RefreshTokensResp{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}, nil
}

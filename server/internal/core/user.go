package core

import (
	"context"
	"fmt"

	"olycall-server/internal/core/domain"

	"github.com/google/uuid"
)

func (s Service) GetUser(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	user, err := s.userStore.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get user by id: %w", err)
	}

	if user == nil {
		return nil, domain.ErrUserNotFound
	}

	return &domain.User{
		ID:        user.ID,
		Username:  user.Username,
		Name:      user.Name,
		AvatarURL: user.AvatarURL,
	}, nil
}

func (s Service) GetMe(ctx context.Context, accessToken string) (*domain.Profile, error) {
	userID, err := s.getUserIDFromJWT(accessToken)
	if err != nil {
		return nil, domain.ErrInvalidToken
	}

	user, err := s.userStore.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get user by id: %w", err)
	}

	if user == nil {
		return nil, domain.ErrUserNotFound
	}

	return &domain.Profile{
		ID:        user.ID,
		Username:  user.Username,
		Name:      user.Name,
		AvatarURL: user.AvatarURL,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}

//	type UpdateMeParams struct {
//		AccessToken string
//		Username    string
//		Name        string
//	}
//
//	func (s Service) UpdateMe(ctx context.Context, params *UpdateMeParams) (*domain.User, error) {
//		userID, err := s.getUserIDFromJWT(params.AccessToken)
//		if err != nil {
//			return nil, domain.ErrInvalidToken
//		}
//
//		if err := validation.ValidateStructWithContext(ctx, params,
//			validation.Field(&params.Username, UserNameRule...),
//		); err != nil {
//			return nil, wrapInvalidParamsErr(err)
//		}
//
//		user, err := s.userStore.GetUserByID(ctx, userID)
//		if err != nil {
//			return nil, fmt.Errorf("get user by id: %w", err)
//		}
//
//		if user == nil {
//			return nil, domain.ErrUserNotFound
//		}
//
//		if user.Username != params.Username {
//			if _, err := s.userStore.UpdateUser(ctx, &userstore.UpdateUserParams{
//				ID:       userID,
//				Username: params.Username,
//			}); err != nil {
//				return nil, fmt.Errorf("update user by id: %w", err)
//			}
//
//			user.Username = params.Username
//		}
//
//		return &domain.User{
//			ID:        user.ID,
//			Email:     user.Email,
//			Username:  user.Username,
//			CreatedAt: user.CreatedAt,
//		}, nil
//	}

func (s Service) SearchUsers(ctx context.Context, query string) ([]domain.User, error) {
	usersResp, err := s.userStore.SearchUsersByUsername(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("search users: %w", err)
	}

	users := make([]domain.User, len(usersResp))

	for i, user := range usersResp {
		users[i] = domain.User{
			ID:        user.ID,
			Username:  user.Username,
			Name:      user.Name,
			AvatarURL: user.AvatarURL,
		}
	}

	return users, nil
}

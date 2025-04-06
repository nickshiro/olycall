package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s Service) GetUser(ctx context.Context, userID uuid.UUID) (*User, error) {
	user, err := s.domainRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	return &User{
		ID:        userID,
		Email:     user.Email,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s Service) GetMe(ctx context.Context, accessToken string) (*User, error) {
	userID, err := s.getUserIDFromAccessToken(accessToken)
	if err != nil {
		return nil, ErrInvalidToken
	}

	user, err := s.domainRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to select user: %w", err)
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	return &User{
		ID:        userID,
		Email:     user.Email,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

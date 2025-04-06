package service

import (
	"fmt"

	"github.com/google/uuid"
)

func (s Service) getUserIDFromAccessToken(accessToken string) (uuid.UUID, error) {
	claims, err := s.parseJWT(accessToken)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to parse jwt: %w", err)
	}

	parsedID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to parse id: %w", err)
	}

	return parsedID, nil
}

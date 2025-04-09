package auth

import (
	"fmt"

	"github.com/google/uuid"
)

func (s AuthService) getUserIDFromAccessToken(accessToken string) (uuid.UUID, error) {
	claims, err := s.parseJWT(accessToken)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("parse jwt: %w", err)
	}

	parsedID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("parse id: %w", err)
	}

	return parsedID, nil
}

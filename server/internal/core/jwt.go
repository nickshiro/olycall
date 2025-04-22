package core

import (
	"errors"
	"fmt"
	"time"

	"olycall-server/internal/core/domain"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	jwt.RegisteredClaims
}

type tokenPair struct {
	AccessToken  string
	RefreshToken string
}

func (s Service) generateJWT(userID string) tokenPair {
	now := time.Now()

	accessTokenExp := now.Add(7 * 24 * time.Hour)
	refreshTokenExp := now.Add(7 * 24 * time.Hour)

	accessToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(accessTokenExp),
		IssuedAt:  jwt.NewNumericDate(now),
	}).SignedString([]byte(s.secret))

	refreshToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(refreshTokenExp),
		IssuedAt:  jwt.NewNumericDate(now),
	}).SignedString([]byte(s.secret))

	return tokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

var ErrUnexpectedSigningMethod = errors.New("unexpected signing method")

func (s Service) parseJWT(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%w: %v", ErrUnexpectedSigningMethod, token.Header["alg"])
		}

		return []byte(s.secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		if claims.ExpiresAt.Before(time.Now()) {
			return nil, domain.ErrTokenExpired
		}

		return claims, nil
	}

	return nil, domain.ErrInvalidToken
}

func (s Service) getUserIDFromJWT(accessToken string) (uuid.UUID, error) {
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

func (s Service) GetUserIDFromAccessToken(accessToken string) (uuid.UUID, error) {
	return s.getUserIDFromJWT(accessToken)
}

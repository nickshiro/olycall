package service

import (
	"context"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"olycall-server/internal/repository/cache"
	"olycall-server/internal/repository/domain"
)

type DomainRepo interface {
	CreateUser(context.Context, *domain.CreateUserParams) error
	GetUserByEmail(context.Context, string) (*domain.User, error)
	GetUserByID(context.Context, uuid.UUID) (*domain.User, error)
	UpdateUser(context.Context, *domain.UpdateUserParams) (bool, error)
}

type CacheRepo interface {
	SetState(context.Context, *cache.SetStateParams) error
	GetState(context.Context, string) (*cache.State, error)
	RemoveState(context.Context, string) (bool, error)
}

type Service struct {
	domainRepo         DomainRepo
	cacheRepo          CacheRepo
	googleOauth2Config oauth2.Config
	secret             string
}

func New(
	domainRepo DomainRepo,
	cacheRepo CacheRepo,
	googleOauth2Config oauth2.Config,
	secret string,
) *Service {
	googleOauth2Config.Scopes = []string{
		"email",
		"profile",
	}
	googleOauth2Config.Endpoint = google.Endpoint

	return &Service{
		domainRepo:         domainRepo,
		cacheRepo:          cacheRepo,
		secret:             secret,
		googleOauth2Config: googleOauth2Config,
	}
}

package postgres

import "olycall-server/pkg/pg"

type Repository struct {
	db pg.Querier
}

func NewOAuthStore(db pg.Querier) *Repository {
	return &Repository{
		db: db,
	}
}

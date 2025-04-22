package postgres

import "olycall-server/pkg/pg"

type UserStore struct {
	db pg.Querier
}

func New(db pg.Querier) *UserStore {
	return &UserStore{
		db: db,
	}
}

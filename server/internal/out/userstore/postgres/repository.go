package postgres

import "olycall-server/pkg/pg"

type UserStore struct {
	db pg.Querier
}

func NewUserStore(db pg.Querier) *UserStore {
	return &UserStore{
		db: db,
	}
}

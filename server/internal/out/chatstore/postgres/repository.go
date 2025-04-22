package postgres

import (
	"context"
	"fmt"

	"olycall-server/internal/core/ports/chatstore"
	"olycall-server/pkg/pg"

	"github.com/jackc/pgx/v5"
)

type ChatStore struct {
	db pg.QuerierWithTx
	tx pgx.Tx
}

func New(db pg.QuerierWithTx) *ChatStore {
	return &ChatStore{
		db: db,
	}
}

func (s ChatStore) getQuerier() pg.Querier {
	var q pg.Querier

	q = s.db
	if s.tx != nil {
		q = s.tx
	}

	return q
}

func (s ChatStore) WithTx(ctx context.Context, fn func(ctx context.Context, store chatstore.ChatStore) error) error {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	c := ChatStore{
		db: s.db,
		tx: tx,
	}

	err = fn(ctx, c)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("rollback failed: %w (original error: %w)", rbErr, err)
		}

		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit: %w", err)
	}

	return nil
}

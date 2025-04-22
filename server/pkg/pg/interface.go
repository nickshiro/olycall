package pg

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Querier interface {
	Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, query string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, query string, args ...any) pgx.Row
}

type QuerierWithTx interface {
	Querier
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

type Pg struct {
	QuerierWithTx
	tx pgx.Tx
}

func New(q QuerierWithTx) *Pg {
	return &Pg{
		QuerierWithTx: q,
	}
}

func (p Pg) GetQuerier() Querier {
	if p.tx != nil {
		return p.tx
	}

	return p
}

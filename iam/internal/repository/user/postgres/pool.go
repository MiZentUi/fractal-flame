package postgres

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Pool struct {
	pool *pgxpool.Pool
}

type Sqlizer interface {
	ToSql() (string, []any, error)
}

func (p *Pool) Get(ctx context.Context, dst any, sqlizer Sqlizer) error {
	query, args, err := sqlizer.ToSql()
	if err != nil {
		slog.Error("Failed to convert select builder query to SQL", "err", err)

		return fmt.Errorf("builder to sql: %w", err)
	}

	return pgxscan.Get(ctx, p.pool, dst, query, args...)
}

func (p *Pool) Exec(ctx context.Context, sqlizer Sqlizer) (pgconn.CommandTag, error) {
	query, args, err := sqlizer.ToSql()
	if err != nil {
		slog.Error("Failed to convert insert builder query to SQL", "err", err)

		return pgconn.CommandTag{}, fmt.Errorf("builder to sql: %w", err)
	}

	return p.pool.Exec(ctx, query, args...)
}

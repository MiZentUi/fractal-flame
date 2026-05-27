package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	errs "github.com/mizentui/fractal-flame/iam/internal/error"
	"github.com/mizentui/fractal-flame/iam/internal/model"
	"github.com/mizentui/fractal-flame/iam/internal/repository/user/postgres/record"
)

type repository struct {
	pool Pool
}

func New(pool *pgxpool.Pool) *repository {
	return &repository{
		pool: Pool{pool: pool},
	}
}

func (r *repository) Save(ctx context.Context, username, password, image string) (int64, error) {
	builder := sq.Insert(record.UsersTable).
		PlaceholderFormat(sq.Dollar).
		Columns(record.UsersTableColumnUsername, record.UsersTableColumnPassword, record.UsersTableColumnImage).
		Values(username, password, image).
		Suffix("ON CONFLICT(username) DO NOTHING RETURNING id")

	var id int64
	err := r.pool.Get(ctx, &id, builder)
	if err != nil {
		slog.Error("Failed to get user id when creating user", "err", err)

		return 0, fmt.Errorf("get user id: %w", err)
	}

	return id, nil
}

func (r *repository) FindByID(ctx context.Context, id int64) (model.User, error) {
	builder := sq.Select(record.UsersTableColumns...).
		PlaceholderFormat(sq.Dollar).
		From(record.UsersTable).
		Where(sq.Eq{record.UsersTableColumnID: id})

	var user record.UserRow
	err := r.pool.Get(ctx, &user, builder)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Warn("Failed to find user with this id", "id", id, "err", errs.ErrUserNotFound)

			return model.User{}, errs.ErrUserNotFound
		}
	}

	return record.RawToModel(user), nil
}

func (r *repository) FindByUsername(ctx context.Context, username string) (model.User, error) {
	builder := sq.Select(record.UsersTableColumns...).
		PlaceholderFormat(sq.Dollar).
		From(record.UsersTable).
		Where(sq.Eq{record.UsersTableColumnUsername: username})

	var user record.UserRow
	err := r.pool.Get(ctx, &user, builder)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Warn("Failed to find user with this username", "username", username, "err", errs.ErrUserNotFound)

			return model.User{}, errs.ErrUserNotFound
		}
	}

	return record.RawToModel(user), nil
}

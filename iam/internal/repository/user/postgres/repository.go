package postgres

import (
	"context"
	"database/sql"
	"errors"
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

func (r *repository) Save(ctx context.Context, username, password string) (int64, error) {
	builder := sq.Insert(record.UsersTable).
		PlaceholderFormat(sq.Dollar).
		Columns(record.UsersTableColumnUsername, record.UsersTableColumnPassword).
		Values(username, password).
		Suffix("ON CONFLICT(username) DO NOTHING RETURNING id")

	var id int64
	err := r.pool.Get(ctx, &id, builder)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Warn("User with this username already exists", "username", username)

			return 0, errs.ErrUserAlreadyExists
		}
		slog.Error("DB error", "err", err)

		return 0, err
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
			slog.Warn("User with this id not found", "id", id)

			return model.User{}, errs.ErrUserNotFound
		}
		slog.Error("DB error", "err", err)

		return model.User{}, err
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
			slog.Warn("User with this username not found", "username", username)

			return model.User{}, errs.ErrUserNotFound
		}
		slog.Error("DB error", "err", err)

		return model.User{}, err
	}

	return record.RawToModel(user), nil
}

func (r *repository) Update(ctx context.Context, id int64, username, password, image string) (model.User, error) {
	builder := sq.Update(record.UsersTable).
		PlaceholderFormat(sq.Dollar).
		Set(record.UsersTableColumnImage, image).
		Where(sq.Eq{record.UsersTableColumnID: id}).
		Suffix("RETURNING *")

	if username != "" {
		builder = builder.Set(record.UsersTableColumnUsername, username)
	}
	if password != "" {
		builder = builder.Set(record.UsersTableColumnPassword, password)
	}

	var user record.UserRow
	err := r.pool.Get(ctx, &user, builder)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Warn("User with this id not found", "id", id)

			return model.User{}, errs.ErrUserNotFound
		}
		slog.Error("DB error", "err", err)

		return model.User{}, err
	}

	return record.RawToModel(user), nil
}

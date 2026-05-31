package auth

import (
	"context"
	"fmt"
	"log/slog"

	errs "github.com/mizentui/fractal-flame/iam/internal/error"
)

type authKey struct{}

var key authKey

func WithUserID(ctx context.Context, id int64) context.Context {
	return context.WithValue(ctx, key, id)
}

func ExtractUserID(ctx context.Context) (int64, error) {
	id, ok := ctx.Value(key).(int64)
	if !ok {
		slog.Warn("Failed to extract user id from ctx")

		return 0, fmt.Errorf("extract user from context: %w", errs.ErrInvalidCtxValue)
	}

	return id, nil
}

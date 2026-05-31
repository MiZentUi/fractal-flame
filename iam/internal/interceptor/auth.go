package interceptor

import (
	"context"
	"log/slog"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/mizentui/fractal-flame/iam/internal/auth"
	errs "github.com/mizentui/fractal-flame/iam/internal/error"
)

type TokenManager interface {
	ValidateAccessToken(accessToken string) (int64, error)
}

func ExtractIdentity(manager TokenManager) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp any, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok || len(md["authorization"]) == 0 {
			return handler(ctx, req)
		}

		header := strings.Split(md["authorization"][0], " ")
		if len(header) != 2 || header[0] != "Bearer" {
			slog.Warn("Failed, invalid token format")
			return nil, errs.ErrInvalidToken
		}

		token := header[1]

		id, err := manager.ValidateAccessToken(token)
		if err != nil {
			slog.Warn("Failed to extract identity from access token", "method", info.FullMethod)

			return nil, err
		}

		ctx = auth.WithUserID(ctx, id)

		return handler(ctx, req)
	}
}

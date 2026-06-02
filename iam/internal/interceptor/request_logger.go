package interceptor

import (
	"context"
	"log/slog"
	"time"

	"google.golang.org/grpc"
)

func RequestLogger() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp any, err error) {
		start := time.Now()

		slog.Debug("Start IAM service method", "method", info.FullMethod)

		resp, err = handler(ctx, req)
		if err != nil {
			slog.Debug("IAM service method finished with error", "method", info.FullMethod, "took", time.Since(start).String())

			return nil, err
		}

		slog.Debug("IAM service successfully finished", "method", info.FullMethod, "took", time.Since(start).Abs().String())

		return resp, nil
	}
}

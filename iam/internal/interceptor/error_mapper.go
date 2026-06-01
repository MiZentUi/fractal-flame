package interceptor

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	errs "github.com/mizentui/fractal-flame/iam/internal/error"
)

func MappingError() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp any, err error) {
		resp, err = handler(ctx, req)
		if err != nil {
			switch {
			case errors.Is(err, errs.ErrWeakPassword), errors.Is(err, errs.ErrInvalidCredentials),
				errors.Is(err, errs.ErrInvalidCtxValue), errors.Is(err, errs.ErrInvalidUserID),
				errors.Is(err, errs.ErrNothingToUpdate), errors.Is(err, errs.ErrInvalidImageName):
				return nil, status.Error(codes.InvalidArgument, err.Error())
			case errors.Is(err, errs.ErrUserAlreadyExists):
				return nil, status.Error(codes.AlreadyExists, err.Error())
			case errors.Is(err, errs.ErrUserNotFound), errors.Is(err, errs.ErrImageNotFound):
				return nil, status.Error(codes.NotFound, err.Error())
			case errors.Is(err, errs.ErrInvalidToken):
				return nil, status.Error(codes.Unauthenticated, err.Error())
			default:
				return nil, status.Error(codes.Internal, err.Error())
			}
		}

		return resp, nil
	}
}

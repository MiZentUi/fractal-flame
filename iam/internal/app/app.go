package app

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	miniogo "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/mizentui/fractal-flame/iam/internal/api/v1"
	"github.com/mizentui/fractal-flame/iam/internal/auth/password"
	"github.com/mizentui/fractal-flame/iam/internal/auth/password/bcrypt"
	"github.com/mizentui/fractal-flame/iam/internal/auth/token/jwt"
	"github.com/mizentui/fractal-flame/iam/internal/config"
	"github.com/mizentui/fractal-flame/iam/internal/interceptor"
	"github.com/mizentui/fractal-flame/iam/internal/repository/image/minio"
	"github.com/mizentui/fractal-flame/iam/internal/repository/user/postgres"
	"github.com/mizentui/fractal-flame/iam/internal/service/auth"
	"github.com/mizentui/fractal-flame/iam/internal/service/image"
	"github.com/mizentui/fractal-flame/iam/internal/service/user"
	"github.com/mizentui/fractal-flame/iam/pkg/closer"
	"github.com/mizentui/fractal-flame/iam/pkg/grpc/health"
	"github.com/mizentui/fractal-flame/iam/pkg/logger"
	iamv1 "github.com/mizentui/fractal-flame/iam/pkg/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	connectionTimeout = 5 * time.Second
)

func Run() {
	err := config.Setup()
	if err != nil {
		panic("Failed to setup config")
	}

	logger.Init(config.App().Logger.Level(), config.App().Logger.AsJSON())

	ctx, cancel := context.WithTimeout(context.Background(), connectionTimeout)
	defer cancel()

	defer closer.CloseAll(ctx)

	listener, err := net.Listen("tcp", config.App().GRPC.Address())
	if err != nil {
		slog.Error("Failed to listen socket", "err", err)

		return
	}
	closer.Add(func(ctx context.Context) error { return listener.Close() })

	pool, err := pgxpool.New(ctx, config.App().Postgres.URI())
	if err != nil {
		slog.Error("Failed to create postgres pool", "uri", config.App().Postgres.URI(), "err", err)

		return
	}
	closer.Add(func(ctx context.Context) error {
		pool.Close()
		return nil
	})

	client, err := miniogo.New(config.App().Minio.Endpoint(), &miniogo.Options{
		Creds:  credentials.NewStaticV4(config.App().Minio.AccessKey(), config.App().Minio.SecretKey(), ""),
		Secure: true,
	})
	if err != nil {
		slog.Error("Failed to create MinIO client", "endpoint", config.App().Minio.Endpoint(), "err", err)

		return
	}

	repo := postgres.New(pool)
	storage, err := minio.New(ctx, client)
	if err != nil {
		slog.Error("Failed to create image repository", "err", err)

		return
	}
	hasher := bcrypt.New()
	validator := password.New(config.App().Auth.PasswordEntropy())
	manager := jwt.New()

	auth := auth.New(repo, validator, hasher, manager)
	user := user.New(repo, storage, validator, hasher)
	image := image.New(storage)

	api := api.New(auth, user, image)

	server := grpc.NewServer(grpc.UnaryInterceptor(interceptor.ExtractIdentity(manager)))

	reflection.Register(server)

	health.Register(server)

	iamv1.RegisterIAMServiceServer(server, api)

	go func() {
		slog.Debug("Start gRPC server", "address", config.App().GRPC.Address())

		err = server.Serve(listener)
		if err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			slog.Error("Failed to start gRPC server", "address", config.App().GRPC.Address())
		}
	}()

	slog.Debug("Shutting down gRPC server...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	server.GracefulStop()
}

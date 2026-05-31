package minio

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"log/slog"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"

	errs "github.com/mizentui/fractal-flame/iam/internal/error"
	"github.com/mizentui/fractal-flame/iam/internal/repository/image/minio/record"
)

type repository struct {
	client *minio.Client
}

func New(ctx context.Context, client *minio.Client) (*repository, error) {
	repository := &repository{
		client: client,
	}

	err := client.MakeBucket(ctx, record.ImageBucketName, minio.MakeBucketOptions{Region: record.ImageBucketLocation})
	if err != nil {
		exists, errBucketExists := client.BucketExists(ctx, record.ImageBucketName)
		if errBucketExists == nil && exists {
			slog.Debug("We already own this bucket", "bucket", record.ImageBucketName)

			return repository, nil
		}

		slog.Error("Failed to create minio bucket", "bucket", record.ImageBucketName, "err", err)

		return nil, fmt.Errorf("create bucket: %w", err)
	}

	return repository, nil
}

func (r *repository) Save(ctx context.Context, imageB64 string) (string, error) {
	image, err := base64.StdEncoding.DecodeString(imageB64)
	if err != nil {
		slog.Error("Failed to decode image", "err", err)

		return "", err
	}

	name := uuid.NewString()
	reader := bytes.NewReader(image)
	size := len(image)

	_, err = r.client.PutObject(ctx, record.ImageBucketName, name, reader, int64(size), minio.PutObjectOptions{ContentType: "image/png"})
	if err != nil {
		slog.Error("Failed to put image in minio storage", "err", err)

		return "", err
	}

	return name, nil
}

func (r *repository) FindByName(ctx context.Context, name string) ([]byte, error) {
	object, err := r.client.GetObject(ctx, record.ImageBucketName, name, minio.GetObjectOptions{})
	if err != nil {
		slog.Error("Failed to get image by name", "name", name, "err", err)

		return nil, err
	}
	defer func() {
		cerr := object.Close()
		if cerr != nil {
			slog.Error("Catch object close error", "err", err)
		}
	}()

	_, err = object.Stat()
	if err != nil {
		code := minio.ToErrorResponse(err).Code

		if code == record.NoSuchKeyCode {
			slog.Warn("Image not found", "bucket", record.ImageBucketName, "name", name, "err", err)
			return nil, errs.ErrImageNotFound
		}

		slog.Error("Failed to get image from MinIO", "bucket", record.ImageBucketName, "name", name, "err", err)

		return nil, err
	}

	bytes, err := io.ReadAll(object)
	if err != nil {
		slog.Error("Failed to read image bytes", "err", err)

		return nil, err
	}

	return bytes, nil
}

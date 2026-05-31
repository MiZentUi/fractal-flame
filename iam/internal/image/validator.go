package image

import (
	"bytes"
	"errors"
	"image"
	"log/slog"

	errs "github.com/mizentui/fractal-flame/iam/internal/error"
)

type validator struct {
	width  int
	height int
}

func New(width, height int) *validator {
	return &validator{
		width:  width,
		height: height,
	}
}

func (v *validator) Validate(img []byte) error {
	reader := bytes.NewReader(img)

	cfg, _, err := image.DecodeConfig(reader)
	if err != nil {
		if errors.Is(err, image.ErrFormat) {
			slog.Warn("Image format error", "err", err)

			return errs.ErrInvalidImage
		}
		slog.Error("Image error", "err", err)

		return err
	}

	if cfg.Width > v.width || cfg.Height > v.height {
		slog.Error("Invalid image width or height", "err", errs.ErrInvalidImage)

		return errs.ErrInvalidImage
	}

	return nil
}

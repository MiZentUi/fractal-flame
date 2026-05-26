package password

import (
	"log/slog"

	passwordvalidator "github.com/wagslane/go-password-validator"
)

type validator struct {
	entropy float64
}

func New(entropy float64) *validator {
	return &validator{
		entropy: entropy,
	}
}

func (v *validator) Validate(password string) error {
	err := passwordvalidator.Validate(password, v.entropy)
	if err != nil {
		slog.Warn("Failed, invalid password", "err", err)
	}

	return nil
}

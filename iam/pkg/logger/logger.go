package logger

import (
	"log/slog"
	"os"
	"strings"
)

func Init(level string, asJSON bool) {
	slog.SetDefault(slog.New(newHandler(parseLevel(level), asJSON)))
}

func newHandler(level slog.Level, asJSON bool) slog.Handler {
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     level,
	}

	if asJSON {
		return slog.NewJSONHandler(os.Stdout, opts)
	}

	return slog.NewTextHandler(os.Stdout, opts)
}

func parseLevel(level string) slog.Level {
	switch strings.ToUpper(level) {
	case "INFO":
		return slog.LevelInfo
	case "WARN", "WARNING":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelDebug
	}
}

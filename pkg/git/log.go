package git

import (
	"log/slog"
	"os"
)

func SetupLogging(logLevel slog.Level) {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     logLevel,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)
}

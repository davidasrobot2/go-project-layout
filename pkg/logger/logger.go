package logger

import (
	"log/slog"
	"os"

	"github.com/google/wire"
)

// Providers is a Wire provider set that provides a new slog.Logger.
var Providers = wire.NewSet(NewLogger)

// NewLogger creates a new slog.Logger instance that writes to standard output in JSON format.
func NewLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

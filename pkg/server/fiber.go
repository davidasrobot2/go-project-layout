package server

import (
	"davidasrobot/project-layout/internal/app/http/middleware"
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

// Providers is a Wire provider set that provides a new Fiber server.
var Providers = wire.NewSet(NewFiberServer)

// NewFiberServer creates a new Fiber application with a custom error handler.
func NewFiberServer(logger *slog.Logger) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler(logger),
	})
	return app
}

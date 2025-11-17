package middleware

import (
	"davidasrobot/project-layout/pkg/response"
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// ErrorHandler returns a custom error handler function for the Fiber application.
func ErrorHandler(logger *slog.Logger) func(c *fiber.Ctx, err error) error {
	return func(c *fiber.Ctx, err error) error {
		// Check for validator.ValidationErrors
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			errors := response.FormatValidationErrors(validationErrs)
			return response.Error(c, 0, errors)
		}

		// Check for fiber.Error (e.g., 404 Not Found)
		if e, ok := err.(*fiber.Error); ok {
			return response.Error(c, 0, e.Message)
		}

		// Log any other unhandled errors
		logger.Error("unhandled error caught in middleware", "error", err, "path", c.Path())

		// Return a generic 500 Internal Server Error for all other cases
		return response.Error(c, 0, nil)
	}
}

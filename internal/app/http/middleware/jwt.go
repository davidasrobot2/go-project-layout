package middleware

import (
	"davidasrobot/project-layout/config"
	"davidasrobot/project-layout/pkg/constant"
	"davidasrobot/project-layout/pkg/response"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

// Protected returns a JWT middleware that protects routes.
func Protected(cfg *config.Config) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(cfg.JWT.Secret),
		},
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return response.Error(c, fiber.StatusBadRequest, constant.ErrMissingCredential)
	}
	return response.Error(c, fiber.StatusUnauthorized, constant.ErrInvalidExpiredToken)
}

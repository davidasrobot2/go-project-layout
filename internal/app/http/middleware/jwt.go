package middleware

import (
	"davidasrobot2/go-boilerplate/config"
	"davidasrobot2/go-boilerplate/pkg/constant"
	"davidasrobot2/go-boilerplate/pkg/response"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Protected returns a JWT middleware that protects routes.
func Protected(cfg *config.Config) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(cfg.JWT.Secret),
		},
		ErrorHandler: jwtError,
		SuccessHandler: func(c *fiber.Ctx) error {
			// set user id from "sub" in jwt token
			c.Locals("user_id", c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["sub"])
			return c.Next()
		},
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	return response.HandleErrors(c, constant.ErrorMessageInvalidToken)
}

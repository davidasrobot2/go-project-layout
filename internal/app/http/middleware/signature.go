package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	"github.com/gofiber/fiber/v2"
)

func Signed(c *fiber.Ctx) error {
	// Get the X-Signature header
	signature := c.Get("X-Signature")

	// Read the request body
	body := c.Body()

	// Hash the request body using SHA-256 with a signature secret
	// In a real application, the secret would come from your config
	secret := "your-signature-secret" // Replace with actual secret from config
	hasher := hmac.New(sha256.New, []byte(secret))
	hasher.Write(body)
	expectedSignature := hex.EncodeToString(hasher.Sum(nil))

	// Compare the received signature with the expected signature
	if signature != expectedSignature {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid signature",
		})
	}

	// If the signature is valid, continue to the next middleware/handler
	return c.Next()
}

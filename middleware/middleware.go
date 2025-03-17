package middleware

import (
	"os"

	// Ensure this path is correct and the package exists
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func Middleware(c *fiber.Ctx) error {
	// Retrieve the token from the cookie
	token := c.Cookies("token")

	// If there's no token in the cookie, return unauthorized
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "error": "Token mueng mai mee wa"})
	}
	jwtMiddleware := jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")), // Secret key for JWT signing and verification
	})

	// Validate the token
	if err := jwtMiddleware(c); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "error": "Token mueng mai mee wa"})
	}

	// Send a new token and refresh the cookie (use your provided `sendNewTokenRespond` function)
	// If everything is fine, allow the request to continue
	return c.Next()
}

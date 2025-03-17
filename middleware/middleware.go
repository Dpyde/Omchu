package middleware

import (
	"os"

	// Ensure this path is correct and the package exists
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Middleware(c *fiber.Ctx) error {
	// Retrieve the token from the cookie
	tokenString := c.Cookies("token")

	// If there's no token in the cookie, return unauthorized
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "Token mueng mai mee wa",
		})
	}

	// Retrieve JWT secret key
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "JWT secret not configured",
		})
	}

	// Parse and validate the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	// Check if the token is valid
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "Token mueng mai mee wa",
		})
	}

	// Send a new token and refresh the cookie (use your provided `sendNewTokenRespond` function)
	// If everything is fine, allow the request to continue
	return c.Next()
}

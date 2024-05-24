package validating

import (
	"strings"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func SetupMiddleware(app *fiber.App) {
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))
}

func JWTAuth() func(*fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		accessToken := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
		if accessToken == "" {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		c.Locals("user_id", claims["userID"])
		c.Locals("user_type", claims["userType"])
		return c.Next()
	})
}

func IsCooker() func(*fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		userType := c.Locals("user_type")
		if userType != "cooker" {
			return c.Status(403).JSON(fiber.Map{
				"message": "Forbidden",
			})
		}

		return c.Next()
	})
}

func IsStaff() func(*fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		userType := c.Locals("user_type")
		if userType != "staff" {
			return c.Status(403).JSON(fiber.Map{
				"message": "Forbidden",
			})
		}

		return c.Next()
	})
}

package middlewares

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/nekowawolf/airdropv2/utils"
)

func AdminMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		log.Println("Authorization Header:", authHeader)

		if authHeader == "" {
			log.Println("Authorization header is missing")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing Authorization header",
			})
		}

		splitToken := strings.SplitN(authHeader, " ", 2)
		if len(splitToken) != 2 || strings.ToLower(splitToken[0]) != "bearer" {
			log.Println("Invalid Authorization format")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid Authorization format. Use 'Bearer <token>'",
			})
		}

		tokenString := strings.TrimSpace(splitToken[1])
		log.Println("Extracted Token:", tokenString)

		adminID, err := utils.ValidateJWT(tokenString, false)
		if err != nil {
			log.Println("Error validating token:", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		c.Locals("admin_id", adminID)

		return c.Next()
	}
}
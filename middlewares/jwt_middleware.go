package middlewares

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/nekowawolf/airdropv2/utils"
)

func AdminMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("token")
		log.Println("token Header:", authHeader)

		if authHeader == "" {
			log.Println("token header is missing")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing token",
			})
		}

		splitToken := strings.SplitN(authHeader, " ", 2)
		if len(splitToken) != 2 || strings.ToLower(splitToken[0]) != "bearer" {
			log.Println("Invalid token format")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token format",
			})
		}

		tokenString := strings.TrimSpace(splitToken[1]) 
		log.Println("Extracted Token:", tokenString) 

		adminID, err := utils.ValidateJWT(tokenString)
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

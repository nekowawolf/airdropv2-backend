package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nekowawolf/airdropv2/middlewares"
	"github.com/nekowawolf/airdropv2/routes"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
)

func main() {
	app := fiber.New()
	
	app.Use(cors.New(middlewares.Cors))

	routes.SetupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	if err := app.Listen(":" + port); err != nil {
		panic(err)
	}
}
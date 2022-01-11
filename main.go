package main

import (
	"fmt"
	"os"

	"fiber-tx/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	database.ConnectDb()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/SKjustSK/url-shortner-go/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

// setupRoutes defines the application endpoints
func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)

	app.Post("/api/shorten", routes.ShotenURL)
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	// Initialize a new Fiber app instance
	app := fiber.New()

	// Use Logger middleware to log HTTP requests/responses
	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("FRONTEND_DOMAIN"),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Register the routes
	setupRoutes(app)

	// Start the server on the specified port (e.g., :3000)
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}

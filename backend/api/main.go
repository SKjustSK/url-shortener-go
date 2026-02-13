package main

import (
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
	// API Endpoints
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/shorten", routes.ShotenURL)
}

func main() {
	// Load .env for local development.
	// On Render, we ignore the error because variables are injected directly.
	_ = godotenv.Load()

	// Initialize a new Fiber app instance
	app := fiber.New()

	// Use Logger middleware to log HTTP requests/responses
	app.Use(logger.New())

	// CORS Configuration
	// Ensure FRONTEND_DOMAIN in Render Dashboard matches your Vercel URL
	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("FRONTEND_DOMAIN"),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Register the routes
	setupRoutes(app)

	// Start the server
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}

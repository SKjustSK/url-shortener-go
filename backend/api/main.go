package main

import (
	"fmt"
	"log"
	"os"
	"strings"

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

	// Health check for Render deployment tracking
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"status": "ok"})
	})
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

	// --- Port Selection Logic ---
	// 1. Check "PORT" (provided by Render automatically)
	// 2. Fallback to "APP_PORT" (from local .env)
	// 3. Final fallback to ":3000"
	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("APP_PORT")
	}
	if port == "" {
		port = "3000"
	}

	// Ensure the port string starts with a colon for Fiber
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	fmt.Printf("Server is starting on port %s\n", port)

	// Start the server
	log.Fatal(app.Listen(port))
}

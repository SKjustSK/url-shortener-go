package routes

import (
	"github.com/SKjustSK/url-shortner-go/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func ResolveURL(c *fiber.Ctx) error {
	// Get the short ID from the URL parameter (e.g., localhost:3000/abc -> "abc")
	shortURL := c.Params("url")

	// Connect to the Main Database (DB 0) where URL mappings are stored
	r := database.CreateClient(0)
	defer r.Close()

	// Search for the original URL in Redis using the short ID
	originalURL, err := r.Get(database.Ctx, shortURL).Result()

	// Handle the case where the ID does not exist
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "short not found in database",
		})
	} else if err != nil {
		// Handle generic connection errors
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot connect to database",
		})
	}

	// Connect to the Analytics Database (DB 1)
	rInr := database.CreateClient(1)
	defer rInr.Close()

	// Increment the global usage counter for analytics
	_ = rInr.Incr(database.Ctx, "counter")

	// Redirect the user to the original URL
	// Status Code 302: Browser hits server (and count analytics) every time.
	return c.Redirect(originalURL, 302)
}

package routes

import (
	"math"
	"os"
	"strconv"
	"time"

	"github.com/SKjustSK/url-shortner-go/database"
	"github.com/SKjustSK/url-shortner-go/helpers"
	"github.com/asaskevich/govalidator"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Request structure for incoming JSON payload
type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

// Response structure for outgoing JSON payload
type response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"short"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"rate_limit"`
	XRateLimitReset time.Duration `json:"rate_limit_reset"`
}

func ShotenURL(c *fiber.Ctx) error {

	body := new(request)

	// Parse the incoming request body into the struct
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	// --- Rate Limiting Logic ---

	// Connect to Redis DB 1 (used for IP tracking/Rate limiting)
	r2 := database.CreateClient(1)
	defer r2.Close()

	// Check if the user's IP exists in the database
	val, err := r2.Get(database.Ctx, c.IP()).Result()

	if err == redis.Nil {
		// If IP is new, set the quota (from .env) with a 30-minute expiration
		_ = r2.Set(database.Ctx, c.IP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err()
	} else {
		// If IP exists, check if they have requests remaining
		valInt, _ := strconv.Atoi(val)
		if valInt <= 0 {
			// Quota exceeded: Calculate time remaining and return error
			limit, _ := r2.TTL(database.Ctx, c.IP()).Result()
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"error":            "rate limit exceeded",
				"rate_limit_reset": math.Ceil(limit.Minutes()),
			})
		}
	}

	// --- Input Validation ---

	// Ensure the input is a valid URL structure
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid URL",
		})
	}

	// Prevent loop attacks (shortening the localhost/domain itself)
	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"error": "Domain error",
		})
	}

	// Ensure URL has http/https prefix
	body.URL = helpers.EnforceHTTP(body.URL)

	// --- ID Generation ---

	var id string

	if body.CustomShort == "" {
		// Generate a random 6-character UUID
		id = uuid.New().String()[:6]
	} else {
		// Use the user-provided custom alias
		id = body.CustomShort
	}

	// --- Database Storage ---

	// Connect to Redis DB 0 (used for storing URL data)
	r := database.CreateClient(0)
	defer r.Close()

	// Check if the ID is already in use
	val, _ = r.Get(database.Ctx, id).Result()
	if val != "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "URL custom short is already in use",
		})
	}

	// Default expiry to 24 hours if not provided
	if body.Expiry == 0 {
		body.Expiry = 24
	}

	// Save the ShortID -> OriginalURL mapping to Redis
	// Note: Expiry is converted from hours to nanoseconds
	err = r.Set(database.Ctx, id, body.URL, body.Expiry*3600*time.Second).Err()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to connect server",
		})
	}

	// --- Response Handling ---

	resp := response{
		URL:             body.URL,
		CustomShort:     "",
		Expiry:          body.Expiry,
		XRateRemaining:  10, // Default placeholders, updated below
		XRateLimitReset: 30,
	}

	// Decrement the user's rate limit quota in DB 1
	r2.Decr(database.Ctx, c.IP())

	// Fetch updated rate limit details to send back to user
	val, _ = r2.Get(database.Ctx, c.IP()).Result()
	resp.XRateRemaining, _ = strconv.Atoi(val)

	ttl, _ := r2.TTL(database.Ctx, c.IP()).Result()
	resp.XRateLimitReset = ttl / time.Nanosecond / time.Minute

	// Construct the full short URL
	resp.CustomShort = os.Getenv("DOMAIN") + "/" + id

	return c.Status(fiber.StatusOK).JSON(resp)
}

package database

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

// global context used for all Redis operations
var Ctx = context.Background()

// initializes a connection to a specific Redis database (0-15)
func CreateClient(dbNo int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		// Load connection details from environment variables
		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASS"),
		DB:       dbNo,
	})

	return rdb
}

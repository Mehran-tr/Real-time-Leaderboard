package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
)

// Create a global context
var ctx = context.Background()

// NewRedisClient creates a new Redis client connection
func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),     // Redis address, e.g., "localhost:6379"
		Password: os.Getenv("REDIS_PASSWORD"), // No password set
		DB:       0,                           // Use default DB
	})

	// Test the connection
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis successfully")
	return client
}

package database

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

// InitRedisFromDotenv take data from .env and create a new client connection
func InitRedisFromDotenv() (*redis.Client, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	addr := os.Getenv("REDIS_ADDR")
	password := os.Getenv("REDIS_PASS")

	c := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})
	return c, nil
}

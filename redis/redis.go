package redis

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func init() {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	client.Set("Test key", "Test value", 0)
}

func TestFunction() {
	fmt.Println("Redis database for logging")
}

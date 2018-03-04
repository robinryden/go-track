package redis

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
)

type Log struct {
	Name       string
	StatusCode int
	Time       time.Time
}

func Connect() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	return client, nil
}

func Logger(url string, statusCode int, timestamp time.Time) {
	client, err := Connect()
	if err != nil {
		log.Fatal(err)
	}

	logged := &Log{
		url,
		statusCode,
		timestamp,
	}

	data, err := json.Marshal(*logged)
	if err != nil {
		log.Fatal(err)
	}

	client.Set(url, data, 0)
	fmt.Println(client.Get(url))
}

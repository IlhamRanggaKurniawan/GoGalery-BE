package database

import (
	"context"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewRedis() *redis.Client {
	host := os.Getenv("REDIS_HOST")
	client := redis.NewClient(&redis.Options{
		Addr:     host + ":6379",
		Password: "",
		DB:       0,
	})

	return client
}

func SetData(rdc *redis.Client, key string, value string, ttl time.Duration) error {
	resp := rdc.Set(context.Background(), key, value, ttl*time.Second)

	return resp.Err()
}

func GetData(rdc *redis.Client, key string) (string, error) {
	resp := rdc.Get(context.Background(), key)

	if resp.Err() != nil {
		return "", resp.Err()
	}

	data, err := resp.Result()

	if err != nil {
		return "", err
	}

	return data, nil
}

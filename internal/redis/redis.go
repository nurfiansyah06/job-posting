package redis

import (
	"context"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

func ConnectRedis() (*redis.Client, error) {
	redisAddr := os.Getenv("REDIS_URL")
	redisPort := os.Getenv("REDIS_PORT")
	redisDb := 0

	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr + ":" + redisPort,
		Password: "",
		DB:       redisDb,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return redisClient, nil
}

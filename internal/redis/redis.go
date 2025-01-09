package redis

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient() *RedisClient {
	redisAddr := os.Getenv("REDIS_URL")
	redisPort := os.Getenv("REDIS_PORT")
	redisDb := 0

	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr + ":" + redisPort,
		Password: "",
		DB:       redisDb,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Connected to Redis successfully")

	return &RedisClient{
		Client: client,
	}
}

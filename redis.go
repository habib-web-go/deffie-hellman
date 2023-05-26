package main

import (
	"os"

	"github.com/go-redis/redis"
)

func runRedis() {
	redisPort := os.Getenv("REDIS_PORT")
	redisUrl := os.Getenv("REDIS_URL")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisUrl + ":" + redisPort,
		Password: redisPassword,
		DB:       0,
	})

}

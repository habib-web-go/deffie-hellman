package main

import (
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

const (
	ExpirationDuration = time.Minute * 20
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

func setInRedis(key string, data []byte, expiration time.Duration) error {
	err := redisClient.Set(key, data, expiration).Err()
	if err != nil {
		return err
	}

	return err
}

func existInRedis(key string) (bool, error) {
	count, err := redisClient.Exists(key).Result()
	if err != nil {
		return false, err
	}
	return count == 1, nil
}

func getFromRedis(key string) ([]byte, error) {
	result, err := redisClient.Get(key).Result()

	if err != nil {
		return nil, err
	}

	return []byte(result), nil
}

func getClientHandShake(nonce string, serverNonce string) (*clientHandShake, error) {
	var handShakeData clientHandShake

	sha := createSHA1(nonce + serverNonce)
	data, err := getFromRedis(sha)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(data), &handShakeData)
	if err != nil {
		return nil, err
	}

	return &handShakeData, nil
}

func getClientData(authKey int) (*client, error) {
	var clientData client

	data, err := getFromRedis(strconv.Itoa(int(authKey)))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(data), &clientData)
	if err != nil {
		return nil, err
	}

	return &clientData, nil
}

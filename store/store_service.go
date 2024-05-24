package store

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

// constant cache duration for 6 hours
const CacheDuration = 6 * time.Hour

func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()

	if err != nil {
		panic(fmt.Sprintf("Error initializing Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}

// saving the mapping between the original url and the generated url
func SaveUrlMapping(shortUrl string, originalUrl string, salt string) {
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()

	if err != nil {
		panic(fmt.Sprintf("Failed saving key url; Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}

	// saving the salt under a key derived from the short url
	saltKey := "salt:" + shortUrl
	err = storeService.redisClient.Set(ctx, saltKey, salt, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving salt; Error: %v - saltKey: %s - salt: %s\n", err, saltKey, salt))
	}
}

// we should be able to retrieve the long URL once the short one is provided.
func RetrieveInitialUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(ctx, shortUrl).Result()

	if err != nil {
		panic(fmt.Sprintf("Failed retrieving initial url; Error: %v - shortUrl: %s\n", err, shortUrl))
	}

	return result
}

package cache

import (
	"github.com/go-redis/redis/v7"
	"log"
)

var Cache *redis.Client

func init() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Panicf("cache init error: %v", err)
	}
	Cache = client
}

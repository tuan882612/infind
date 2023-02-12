package config

import (
	// "os"

	"os"

	"github.com/go-redis/redis/v8"
)

func ConnectRedis() *redis.Client {
	config := &redis.Options{
		Addr: os.Getenv("redis_add"),
		Password: os.Getenv("redis_psw"),
		DB: 0,
	}
	
	return redis.NewClient(config)
}
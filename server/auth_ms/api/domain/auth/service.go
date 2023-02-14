package auth

import "github.com/go-redis/redis/v8"

type Repository struct {
	Client *redis.Client
}
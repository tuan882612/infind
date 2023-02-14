package auth

import "auth_ms/assets/config"

type AuthRepository interface {
	
}

func NewAuthRepo() AuthRepository {
	return &Repository{
		Client: config.ConnectRedis(),
	}
}
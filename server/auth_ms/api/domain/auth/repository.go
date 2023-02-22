package auth

import "authms/assets/config"

type AuthRepository interface {
	Create() []byte
}

func NewAuthRepo() AuthRepository {
	return &Repository{
		Client: config.ConnectRedis(),
	}
}
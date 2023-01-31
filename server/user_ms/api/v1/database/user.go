package database

import (
	"userms/api/v1/model"
	"userms/assets/config"
)

type UserRepository interface {
	GetUser(username string) model.User
	CreateUser(user model.User) (model.User, error)
	UpdateUser(user model.User) (model.User, error)
	DeleteUser(username string) error
}

func NewRepo() UserRepository {
	client := config.ConnectDynamodb()
	return UserRepository(&Repository{Client: client})
}
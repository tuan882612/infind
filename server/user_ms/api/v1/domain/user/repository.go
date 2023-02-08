package user

import (
	"userms/assets/config"
)

type UserRepository interface {
	GetUser(username string, email string) User
	CreateUser(user User) (User, error)
	UpdateUser(user User) (User, error)
	DeleteUser(username string) error
}

func NewRepo() UserRepository {
	client := config.ConnectDynamodb()
	return &Repository{Client: client, TableName: "user"}
}
package login

import (
	"userms/api/v1/domain/user"
	"userms/assets/config"
)

type LoginRepository interface {
	GetUser(username string, email string) user.User
	CreateUser(user user.User) (user.User, error)
}

func NewRepo() LoginRepository {
	client := config.ConnectDynamodb()
	return &Repository{Client: client, TableName: "user"}
}

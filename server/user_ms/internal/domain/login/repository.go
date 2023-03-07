package login

import (
	"userms/internal/domain/user"
	"userms/pkg/config"
)

type LoginRepository interface {
	GetUser(username string, email string) user.User
	CreateUser(user user.User) error
}

func NewRepo() LoginRepository {
	client := config.ConnectDynamodb()
	return &Repository{Client: client, TableName: "user"}
}

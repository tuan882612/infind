package repository

import (
	"userms/api/v1/model"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type MockUserRepository interface {
	GetUser(username string, email string) model.User
	CreateUser(user model.User) (model.User, error)
	UpdateUser(user model.User) (model.User, error)
	DeleteUser(username string) error
}

func MockRepo(dynamo dynamodbiface.DynamoDBAPI) MockUserRepository {
	return MockUserRepository(&MockRepository{Client: dynamo})
}
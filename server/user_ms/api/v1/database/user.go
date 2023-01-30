package database

import (
	"userms/assets/config"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type UserRepository interface {
	GetUser(username string) (*dynamodb.GetItemOutput, error)
	CreateUser(user  map[string]*dynamodb.AttributeValue) error
	// UpdateUser(user io.ReadCloser) (*dynamodb.UpdateItemOutput, error)
	// DeleteUser(username string) (*dynamodb.DeleteItemOutput, error)
}

func NewRepo() UserRepository {
	client := config.ConnectDynamodb()
	return UserRepository(&Repository{Client: client})
}
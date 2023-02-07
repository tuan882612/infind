package repository

import (
	"userms/api/v1/model"
	"userms/api/v1/repository/adapter"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Repository struct {
	Client *dynamodb.DynamoDB
}

func (r *Repository) GetUser(username string, email string) model.User {
	query := adapter.GetInput(username, email, "user")
	data, _ := r.Client.GetItem(query)
	
	user := model.User{}

	if err := dynamodbattribute.UnmarshalMap(data.Item, &user); err != nil {
		return model.User{}
	}
	return user
}

func (r *Repository) CreateUser(user model.User) (model.User, error) {
	User, _ := dynamodbattribute.MarshalMap(user)
	
	query := adapter.PutInput(User, "user")
	_, err := r.Client.PutItem(query)

	return user, err
}

func (r *Repository) UpdateUser(user model.User) (model.User, error) {
	query := adapter.UpdateInput(user, "user")
	_, err := r.Client.UpdateItem(query)

	return user, err
}

func (r *Repository) DeleteUser(username string) error {
	query := adapter.DeleteInput(username, "user")
	_, err := r.Client.DeleteItem(query)

	return err
}

package repository

import (
	"userms/api/v1/model"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Repository struct {
	Client *dynamodb.DynamoDB
}

var TableName string = "user"

func (r *Repository) GetUser(username string, email string) model.User {
	data, _ := r.Client.GetItem(GetInput(username, email))
	user := model.User{}

	if err := dynamodbattribute.UnmarshalMap(data.Item, &user); err != nil {
		return model.User{}
	}
	return user
}

func (r *Repository) CreateUser(user model.User) (model.User, error) {
	User, _ := dynamodbattribute.MarshalMap(user)

	_, err := r.Client.PutItem(PutInput(User))

	return user, err
}

func (r *Repository) UpdateUser(user model.User) (model.User, error) {
	_, err := r.Client.UpdateItem(UpdateInput(user))

	return user, err
}

func (r *Repository) DeleteUser(username string) error {
	_, err := r.Client.DeleteItem(DeleteInput(username))

	return err
}

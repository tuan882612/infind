package repository

import (
	"userms/api/v1/model"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type MockRepository struct {
	Client dynamodbiface.DynamoDBAPI
}

func (m *MockRepository) GetUser(username string, email string) model.User {
	data, _ := m.Client.GetItem(GetInput(username, email))
	user := model.User{}

	if err := dynamodbattribute.UnmarshalMap(data.Item, &user); err != nil {
		return model.User{}
	}
	return user
}

func (m *MockRepository) CreateUser(user model.User) (model.User, error) {
	User, _ := dynamodbattribute.MarshalMap(user)

	_, err := m.Client.PutItem(PutInput(User))

	return user, err
}

func (m *MockRepository) UpdateUser(user model.User) (model.User, error) {
	_, err := m.Client.UpdateItem(UpdateInput(user))

	return user, err
}

func (m *MockRepository) DeleteUser(username string) error {
	_, err := m.Client.DeleteItem(DeleteInput(username))

	return err
}

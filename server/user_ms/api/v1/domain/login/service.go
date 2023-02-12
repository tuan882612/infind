package login

import (
	"userms/utility/security"
	"userms/api/v1/domain/user"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Repository struct {
	Client    *dynamodb.DynamoDB
	TableName string
}

func (r *Repository) GetUser(username string, password string) user.User {
	key := map[string]*dynamodb.AttributeValue{
		"username": {S: &username},
	}
	query := &dynamodb.GetItemInput{
		TableName: aws.String(r.TableName),
		Key:       key,
	}

	out, _ := r.Client.GetItem(query)

	User := user.User{}

	if err := dynamodbattribute.UnmarshalMap(out.Item, &User); err != nil {
		return user.User{}
	}

	if !security.ValidateHash(password, &User.Password) {
		return user.User{}
	}

	return User
}

func (r *Repository) CreateUser(user user.User) error {
	user.Password, _ = security.HashPassword(user.Password)
	User, _ := dynamodbattribute.MarshalMap(&user)

	query := &dynamodb.PutItemInput{
		TableName: aws.String(r.TableName),
		Item:      User,
	}
	_, err := r.Client.PutItem(query)

	return err
}
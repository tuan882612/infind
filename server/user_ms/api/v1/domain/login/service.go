package login

import (
	"userms/api/v1/domain/user"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Repository struct {
	Client    *dynamodb.DynamoDB
	TableName string
}

func (r *Repository) GetUser(username string, email string) user.User {
	key := map[string]*dynamodb.AttributeValue{}

	if email != "" {
		key = map[string]*dynamodb.AttributeValue{
			"username": {S: &username},
			"email":    {S: &email},
		}
	} else {
		key = map[string]*dynamodb.AttributeValue{
			"username": {S: &username},
		}
	}

	query := &dynamodb.GetItemInput{
		TableName: aws.String(r.TableName),
		Key:       key,
	}

	data, _ := r.Client.GetItem(query)

	User := user.User{}

	if err := dynamodbattribute.UnmarshalMap(data.Item, &User); err != nil {
		return user.User{}
	}

	return User
}

func (r *Repository) CreateUser(user user.User) (user.User, error) {
	User, _ := dynamodbattribute.MarshalMap(&user)

	query := &dynamodb.PutItemInput{
		TableName: aws.String(r.TableName),
		Item:      User,
	}
	_, err := r.Client.PutItem(query)

	return user, err
}
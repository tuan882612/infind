package database

import (
	"userms/api/v1/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	// "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Repository struct{
	Client *dynamodb.DynamoDB
}

var TableName string = "user"

func (r *Repository) GetUser(username string) (*dynamodb.GetItemOutput, error) {
	return r.Client.GetItem(&dynamodb.GetItemInput{
		TableName: &TableName,
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(username),
			},
		},
	})
}

func (r *Repository) CreateUser(user model.User) (*dynamodb.PutItemOutput, error) {
	new_user, _ := dynamodbattribute.MarshalMap(user)
	return r.Client.PutItem(&dynamodb.PutItemInput{
		TableName: &TableName,
		ConditionExpression: aws.String("attribute_not_exists(pk)"),
		Item: new_user,
	})
}

// func (r *Repository) UpdateUser(user io.ReadCloser) (*dynamodb.UpdateItemOutput, error) {}

// func (r *Repository) DeleteUser(username string) (*dynamodb.DeleteItemOutput, error) {}
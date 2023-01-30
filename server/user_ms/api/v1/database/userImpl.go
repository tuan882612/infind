package database

import (
	// "userms/api/v1/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	// "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	// "github.com/aws/aws-sdk-go/service/dynamodb/expression"
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

func (r *Repository) CreateUser(user map[string]*dynamodb.AttributeValue) error {
	// pkCondition := expression.
	// 	Name("partion_key").NotEqual(expression.Value(user.Username))

	// exp := expression.AttributeNotExists()

	_, err := r.Client.PutItem(&dynamodb.PutItemInput{
		TableName: &TableName,
		// ConditionExpression: exp,
		// ExpressionAttributeValues: exp.Values(),
		// ExpressionAttributeNames: exp.Names(),
		Item: user,
	})
	return err
}

// func (r *Repository) UpdateUser(user io.ReadCloser) (*dynamodb.UpdateItemOutput, error) {}

// func (r *Repository) DeleteUser(username string) (*dynamodb.DeleteItemOutput, error) {}
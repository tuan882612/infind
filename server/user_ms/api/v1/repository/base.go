package repository

import (
	"userms/api/v1/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func GetInput(username string, email string) *dynamodb.GetItemInput {
	key := map[string]*dynamodb.AttributeValue{}

	if email != "" {
		key = map[string]*dynamodb.AttributeValue{
			"username": {S: &username},
			"email": {S: &email},
		}
	} else {
		key = map[string]*dynamodb.AttributeValue{
			"username": {S: &username},
		}
	}
	
	return &dynamodb.GetItemInput {
		TableName: &TableName,
		Key: key,
	}
}

func PutInput(user map[string]*dynamodb.AttributeValue) *dynamodb.PutItemInput {
	return &dynamodb.PutItemInput{
		TableName: &TableName,
		Item:      user,
	}
}

func UpdateInput(user model.User) *dynamodb.UpdateItemInput {
	SerializedHistory, _ := dynamodbattribute.MarshalList(user.History)

	return &dynamodb.UpdateItemInput{
		TableName: &TableName,
		Key: map[string]*dynamodb.AttributeValue{
			"username": {S: &user.Username},
		},
		ExpressionAttributeNames: map[string]*string{
			"#E": aws.String("email"),
			"#P": aws.String("password"),
			"#C": aws.String("created"),
			"#H": aws.String("history"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":e": {S: &user.Email},
			":p": {S: &user.Password},
			":c": {S: &user.Created},
			":h": {L: SerializedHistory},
		},
		UpdateExpression: aws.String("SET #E = :e, #P = :p, #C = :c, #H = :h"),
		ReturnValues:     aws.String("ALL_NEW"),
	}
}

func DeleteInput(username string) *dynamodb.DeleteItemInput {
	return &dynamodb.DeleteItemInput{
		TableName: &TableName,
		Key: map[string]*dynamodb.AttributeValue{
			"username": {S: &username},
		},
	}
}
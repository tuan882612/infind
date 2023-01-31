package database

import (
	"userms/api/v1/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Repository struct{
	Client *dynamodb.DynamoDB
}

var TableName string = "user"

func (r *Repository) GetUser(username string) model.User {
	data, _ := r.Client.GetItem(&dynamodb.GetItemInput{
		TableName: &TableName,
		Key: map[string]*dynamodb.AttributeValue{
			"username": { S: &username },
		},
	})
	user := model.User{}
	dynamodbattribute.UnmarshalMap(data.Item, &user)
	return user
}

func (r *Repository) CreateUser(user model.User) (model.User, error) {
	User, _ := dynamodbattribute.MarshalMap(user)

	_, err := r.Client.PutItem(&dynamodb.PutItemInput{
		TableName: &TableName,
		Item: User,
		ReturnValues: aws.String("ALL_NEW"),
	})

	return user, err
}

func (r *Repository) UpdateUser(user model.User) (model.User, error) {
	SerializedHistory, _ := dynamodbattribute.MarshalList(user.History)

	_, err := r.Client.UpdateItem(&dynamodb.UpdateItemInput{
		TableName: &TableName,
		Key: map[string]*dynamodb.AttributeValue{
			"username": { S: &user.Username },
		},
		ExpressionAttributeNames: map[string]*string{
			"#E": aws.String("email"),
			"#P": aws.String("password"),
			"#C": aws.String("created"),
			"#H": aws.String("history"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":e": { S: &user.Email },
			":p": { S: &user.Password },
			":c": { S: &user.Created },
			":h": { L: SerializedHistory },
		},
		UpdateExpression: aws.String("SET #E = :e, #P = :p, #C = :c, #H = :h"),
		ReturnValues: aws.String("ALL_NEW"),
	})

	return user, err
}

func (r *Repository) DeleteUser(username string) error {
	_, err := r.Client.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: &TableName,
		Key: map[string]*dynamodb.AttributeValue{
			"username": { S: &username },
		},
	})

	return err
}
package repository

import (
	"userms/api/v1/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Repository struct {
	Client *dynamodb.DynamoDB
	TableName string
}

func (r *Repository) GetUser(username string, email string) model.User {
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

	query := &dynamodb.GetItemInput {
		TableName: aws.String(r.TableName),
		Key: key,
	}

	data, _ := r.Client.GetItem(query)
	
	user := model.User{}

	if err := dynamodbattribute.UnmarshalMap(data.Item, &user); err != nil {
		return model.User{}
	}
	
	return user
}

func (r *Repository) CreateUser(user model.User) (model.User, error) {
	User, _ := dynamodbattribute.MarshalMap(user)
	
	query := &dynamodb.PutItemInput{
		TableName: aws.String(r.TableName),
		Item:      User,
	}
	_, err := r.Client.PutItem(query)

	return user, err
}

func (r *Repository) UpdateUser(user model.User) (model.User, error) {
	SerializedHistory, _ := dynamodbattribute.MarshalList(user.History)

	query := &dynamodb.UpdateItemInput{
		TableName: aws.String(r.TableName),
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
	_, err := r.Client.UpdateItem(query)

	return user, err
}

func (r *Repository) DeleteUser(username string) error {
	query := &dynamodb.DeleteItemInput{
		TableName: aws.String(r.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {S: &username},
		},
	}
	_, err := r.Client.DeleteItem(query)

	return err
}

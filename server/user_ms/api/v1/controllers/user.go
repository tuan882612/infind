package controllers

import (
	"net/http"
	"userms/api/v1/model"
	"userms/assets/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gin-gonic/gin"
)

var TableName string = "user"

func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.Query("username")
		
		db := config.Connect_Dynamodb()
		data, _ := db.GetItem(&dynamodb.GetItemInput{
			TableName: &TableName,
			Key: map[string]*dynamodb.AttributeValue{
				"username": {
					S: aws.String(username),
				},
			},
		})

		user := model.User{}
		dynamodbattribute.UnmarshalMap(data.Item, &user)

		if user.Username == "" {
			body := model.DefaultResponse{
				Code: http.StatusNotFound,
				Message: "User does not exist in the database",
				Body: map[string]string{},
			}
			ctx.JSON(http.StatusNotFound, body)
		} else {
			body := model.UserResponse{
				Code: http.StatusOK,
				Message: "",
				Body: user,
			}
			ctx.JSON(http.StatusOK, body)
		}
	}
}

func CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
	}
}

func UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
	}
}
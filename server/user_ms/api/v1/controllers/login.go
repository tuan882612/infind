package controllers

import (
	"userms/assets/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
)

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.Query("username")
		password := ctx.Query("password")

		db := config.ConnectDynamodb()
		data, _ := db.GetItem(&dynamodb.GetItemInput{
			TableName: &TableName,
			Key: map[string]*dynamodb.AttributeValue{
				"username": {
					S: aws.String(username),
				},
				"password": {
					S: aws.String(password),
				},
			},
		})

		
	}
}
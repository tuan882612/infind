package controllers

import (
	"net/http"
	"userms/api/v1/model"
	"userms/api/validators"
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
			},
		})

		body := model.LoginResponse{}

		if validators.ValidateHash(password, data.Item["password"].S) {
			body.Found = true
		} else {
			body.Found = false
		}

		ctx.JSON(http.StatusOK, body)
	}
}
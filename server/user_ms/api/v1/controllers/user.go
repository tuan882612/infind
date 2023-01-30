package controllers

import (
	"net/http"
	"userms/api/response"
	"userms/api/v1/database"
	"userms/api/v1/model"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Repo database.UserRepository
}

func (u UserController) Find(ctx *gin.Context) {
	username := ctx.Query("username")
	data, _ := u.Repo.GetUser(username)

	user := model.User{}
	dynamodbattribute.UnmarshalMap(data.Item, &user)

	if msg := ""; user.Username == "" {
		msg := "User does not exist in the database"
		body := response.None(msg)

		ctx.JSON(http.StatusNotFound, body)
	} else {
		body := response.FoundUser(msg, user)
		
		ctx.JSON(http.StatusOK, body)
	}
}

func (u UserController) Update(ctx *gin.Context) {}

func (u UserController) Delete(ctx *gin.Context) {
}
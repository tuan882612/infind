package controllers

import (
	"net/http"
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

func (u UserController) Update(ctx *gin.Context) {}

func (u UserController) Delete(ctx *gin.Context) {
}
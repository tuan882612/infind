package controllers

import (
	"log"
	"net/http"
	"userms/api/v1/database"
	"userms/api/response"
	"userms/api/v1/model"
	"userms/api/validators"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	Repo database.UserRepository
}

func (l LoginController) Login(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")

	data, _ := l.Repo.GetUser(username)

	body := response.Login{}

	if validators.ValidateHash(password, data.Item["password"].S) {
		body.Found = true
	} else {
		body.Found = false
	}

	ctx.JSON(http.StatusOK, body)
}

func (l LoginController) Register(ctx *gin.Context) {
	user := model.User{}
	ctx.BindJSON(&user)
	NewUser, _ := dynamodbattribute.MarshalMap(user)

	err := l.Repo.CreateUser(NewUser)

	if err != nil {
		log.Fatalln(err)
	}

	body := response.Base{}

	ctx.JSON(http.StatusOK, body)
}
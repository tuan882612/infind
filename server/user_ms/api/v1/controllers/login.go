package controllers

import (
	"log"
	"net/http"
	"userms/api/response"
	"userms/api/utility"
	"userms/api/v1/database"
	"userms/api/v1/model"
	"userms/api/validators"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	Repo database.UserRepository
}

func (l LoginController) Login(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")

	user := l.Repo.GetUser(username)

	res := response.Login{}

	if validators.ValidateHash(password, &user.Password) {
		res.Found = true
	} else {
		res.Found = false
	}

	ctx.JSON(http.StatusOK, res)
}

func (l LoginController) Register(ctx *gin.Context) {
	user := model.User{}
	ctx.BindJSON(&user)

	check := l.Repo.GetUser(user.Username);

	if check.Username != "" {
		res := response.Exist(user)

		ctx.JSON(http.StatusConflict, res)
	} else {
		user.Password, _ = utility.HashPassword(user.Password)

		if _, err := l.Repo.CreateUser(user); err != nil {
			log.Fatalln(err)
		}

		res := response.Created(user)
	
		ctx.JSON(http.StatusOK, res)
	}
}
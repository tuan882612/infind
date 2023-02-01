package controllers

import (
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
		ctx.SetCookie(
			"user",
			user.Username,
			120,
			"/",
			"localhost",
			false,
			true,
		)
	} else {
		res.Found = false
	}

	ctx.JSON(http.StatusOK, res)
}

func (l LoginController) Register(ctx *gin.Context) {
	user := model.User{}

	if err := ctx.BindJSON(&user); err != nil {
		res := response.Custom(
			map[string]string{},
			http.StatusBadRequest,
			"Invalid body recieved.",
		)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	check := l.Repo.GetUser(user.Username)

	if check.Username != "" {
		res := response.Custom(
			user,
			http.StatusConflict,
			"User already exist.",
		)
		ctx.JSON(http.StatusConflict, res)
		return
	}
	user.Password, _ = utility.HashPassword(user.Password)

	if _, err := l.Repo.CreateUser(user); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}

	res := response.Custom(user, http.StatusCreated, "")

	ctx.JSON(http.StatusCreated, res)
}

package controllers

import (
	"log"
	"net/http"
	"userms/api/response"
	"userms/api/utility"
	"userms/api/v1/database"
	"userms/api/v1/model"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Repo database.UserRepository
}

func (u UserController) Find(ctx *gin.Context) {
	username := ctx.Query("username")
	user := u.Repo.GetUser(username)
	
	if user.Username == "" {
		msg := "User does not exist in the database"
		res := response.None(msg)

		ctx.JSON(http.StatusNotFound, res)
	} else {
		res := response.FoundUser("", user)
		
		ctx.JSON(http.StatusOK, res)
	}
}

func (u UserController) Update(ctx *gin.Context) {
	user := model.User{}
	ctx.BindJSON(&user)

	check := u.Repo.GetUser(user.Username)
	
	if check.Username == "" {
		msg := "The user does not exist or you entered the wrong username."
		res := response.None(msg)

		ctx.JSON(http.StatusNotFound, res)
	} else {
		user.Password, _ = utility.HashPassword(user.Password)
		
		if _, err := u.Repo.UpdateUser(user); err != nil {
			log.Fatalln(err)
		} else {
			res := response.Updated(user)

			ctx.JSON(http.StatusAccepted, res)
		}
	}
}

func (u UserController) Delete(ctx *gin.Context) {
	username := ctx.Query("username")
	user := u.Repo.GetUser(username)

	if user.Username == "" {
		msg := "User does not exist in the database"
		res := response.None(msg)

		ctx.JSON(http.StatusNotFound, res)
	} else {
		if err := u.Repo.DeleteUser(username); err != nil {
			res := response.Error(err)

			ctx.JSON(http.StatusBadRequest, res)
		} else {
			msg := username+" has been removed from the database."
			res := response.Accepted(msg)

			ctx.JSON(http.StatusAccepted, res)
		}
	}
}
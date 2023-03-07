package handlers

import (
	"net/http"
	"userms/pkg/response"
	u "userms/internal/domain/user"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Repo u.UserRepository
}

func (uc UserController) Find(ctx *gin.Context) {
	username := ctx.Query("username")
	user := uc.Repo.GetUser(username, "")

	if user.Username == "" {
		res := response.Custom(
			map[string]string{},
			http.StatusNotFound,
			"User does not exist.",
		)
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	res := response.Custom(user, http.StatusOK, "")

	ctx.JSON(http.StatusOK, res)
}

func (uc UserController) Update(ctx *gin.Context) {
	var user u.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		res := response.Custom(
			map[string]string{},
			http.StatusBadRequest,
			"Invalid body recieved.",
		)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	check := uc.Repo.GetUser(user.Username, user.Email)

	if check.Username == "" {
		res := response.Custom(
			map[string]string{},
			http.StatusNotFound,
			"User does not exist.",
		)
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	if _, err := uc.Repo.UpdateUser(user); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}

	res := response.Custom(user, http.StatusAccepted, "")

	ctx.JSON(http.StatusAccepted, res)
}

func (uc UserController) Delete(ctx *gin.Context) {
	username := ctx.Query("username")
	user := uc.Repo.GetUser(username, "")

	if user.Username == "" {
		res := response.Custom(
			map[string]string{},
			http.StatusNotFound,
			"User does not exist.",
		)
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	if err := uc.Repo.DeleteUser(username); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}

	res := response.Custom(
		map[string]string{},
		http.StatusAccepted,
		"User has been removed.",
	)
	ctx.JSON(http.StatusAccepted, res)
}

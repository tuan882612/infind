package user

import (
	"net/http"
	"userms/api/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Repo UserRepository
}

func (u UserController) Find(ctx *gin.Context) {
	username := ctx.Query("username")
	user := u.Repo.GetUser(username, "")

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

func (u UserController) Update(ctx *gin.Context) {
	user := User{}

	if err := ctx.ShouldBindJSON(user); err != nil {
		res := response.Custom(
			map[string]string{},
			http.StatusBadRequest,
			"Invalid body recieved.",
		)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	check := u.Repo.GetUser(user.Username, user.Email)

	if check.Username == "" {
		res := response.Custom(
			map[string]string{},
			http.StatusNotFound,
			"User does not exist.",
		)
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	if _, err := u.Repo.UpdateUser(user); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}

	res := response.Custom(user, http.StatusAccepted, "")

	ctx.JSON(http.StatusAccepted, res)
}

func (u UserController) Delete(ctx *gin.Context) {
	username := ctx.Query("username")
	user := u.Repo.GetUser(username, "")

	if user.Username == "" {
		res := response.Custom(
			map[string]string{},
			http.StatusNotFound,
			"User does not exist.",
		)
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	if err := u.Repo.DeleteUser(username); err != nil {
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

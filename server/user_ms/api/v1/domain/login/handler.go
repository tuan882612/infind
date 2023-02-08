package login

import (
	"net/http"
	
	"userms/api/response"
	"userms/api/v1/domain/user"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	Repo LoginRepository
}

func (l LoginController) Login(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")

	user := l.Repo.GetUser(username, password)

	if user.Username == "" {
		res := response.Custom(
			response.Login{
				Found: false,
			},
			http.StatusNotFound,
			"Invalid login credentials",
		)

		ctx.JSON(http.StatusNotFound, res)
		return
	}

	res := response.Custom(
		response.Login{
			Found: true,
		},
		http.StatusOK,
		"",
	)

	ctx.JSON(http.StatusOK, res)
}

func (l LoginController) Register(ctx *gin.Context) {
	user := user.User{}

	if err := ctx.ShouldBindJSON(user); err != nil {
		res := response.Custom(
			map[string]string{},
			http.StatusBadRequest,
			"Invalid body recieved.",
		)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	check := l.Repo.GetUser(user.Username, user.Email)

	if check.Username != "" {
		res := response.Custom(
			user,
			http.StatusConflict,
			"User already exist.",
		)
		ctx.JSON(http.StatusConflict, res)
		return
	}
	
	if err := l.Repo.CreateUser(user); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}

	res := response.Custom(user, http.StatusCreated, "")

	ctx.JSON(http.StatusCreated, res)
}

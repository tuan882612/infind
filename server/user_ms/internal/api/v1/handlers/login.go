package handlers

import (
	"net/http"

	"userms/pkg/response"
	l "userms/internal/domain/login"
	u "userms/internal/domain/user"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	Repo l.LoginRepository
}

func (lc LoginController) Login(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")

	user := lc.Repo.GetUser(username, password)

	if user.Username == "" {
		res := response.Custom(
			l.Login{Found: false},
			http.StatusNotFound,
			"Invalid login credentials",
		)

		ctx.JSON(http.StatusNotFound, res)
		return
	}

	res := response.Custom(
		l.Login{Found: true},
		http.StatusOK,
		"",
	)

	ctx.JSON(http.StatusOK, res)
}

func (lc LoginController) Register(ctx *gin.Context) {
	var User u.User

	if err := ctx.ShouldBindJSON(&User); err != nil {
		res := response.Custom(
			map[string]string{},
			http.StatusBadRequest,
			"Invalid body recieved.",
		)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	
	check := lc.Repo.GetUser(User.Username, User.Password)

	if check.Username != "" {
		res := response.Custom(
			User,
			http.StatusConflict,
			"User already exist.",
		)
		ctx.JSON(http.StatusConflict, res)
		return
	}

	if err := lc.Repo.CreateUser(User); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}

	res := response.Custom(User, http.StatusCreated, "")

	ctx.JSON(http.StatusCreated, res)
}

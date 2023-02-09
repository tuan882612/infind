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
	User := user.User{}

	if err := ctx.ShouldBindJSON(&User); err != nil {
		res := response.Custom(
			map[string]string{},
			http.StatusBadRequest,
			"Invalid body recieved.",
		)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	ch := make(chan int)

	go func(ctx *gin.Context) {
		check := l.Repo.GetUser(User.Username, User.Password)
		
		if check.Username != "" {
			res := response.Custom(
				User,
				http.StatusConflict,
				"User already exist.",
			)
			ctx.JSON(http.StatusConflict, res)

		} else {
			if err := l.Repo.CreateUser(User); err != nil {
				ctx.JSON(http.StatusBadRequest, response.Error(err))
			} else {
				res := response.Custom(User, http.StatusCreated, "")

				ctx.JSON(http.StatusCreated, res)
			}
		}
		ch <- 0
	}(ctx)

	<- ch

}

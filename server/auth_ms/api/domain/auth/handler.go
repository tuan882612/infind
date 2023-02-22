package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthRepo AuthRepository
}

func (a AuthController) CreateToken(ctx *gin.Context) {
	out := a.AuthRepo.Create()
	ctx.JSON(http.StatusOK, out)
}
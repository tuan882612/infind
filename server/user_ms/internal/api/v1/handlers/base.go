package handlers

import (
	"net/http"
	"user_ms/pkg/response"

	"github.com/gin-gonic/gin"
)

func Default(ctx *gin.Context) {
	body := map[string]interface{}{
		"information": "infind user service",
		"version":     "0.2.0",
	}
	res := response.Custom(body, http.StatusOK, "")

	ctx.JSON(http.StatusOK, res)
}

func Health(ctx *gin.Context) {
	body := map[string]string{"health":"ok"}
	ctx.JSON(http.StatusOK, body)
}

func None(ctx *gin.Context) {
	res := response.Custom(
		map[string]string{},
		http.StatusNotFound,
		"That endpoint does not exist",
	)

	ctx.JSON(http.StatusNotFound, res)
}
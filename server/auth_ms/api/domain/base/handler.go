package base

import (
	"authms/api/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Default(ctx *gin.Context) {
	body := map[string]interface{}{
		"information": "infind auth service",
		"version":     "0.2.0",
	}
	res := response.Custom(body, http.StatusOK, "")

	ctx.JSON(http.StatusOK, res)
}

func None(ctx *gin.Context) {
	res := response.Custom(
		map[string]string{},
		http.StatusNotFound,
		"That endpoint does not exist",
	)

	ctx.JSON(http.StatusNotFound, res)
}
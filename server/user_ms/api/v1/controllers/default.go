package controllers

import (
	"net/http"
	"userms/api/response"

	"github.com/gin-gonic/gin"
)

func Base(ctx *gin.Context) {
	body := map[string]string{
		"information": "infind user service",
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

package controllers

import (
	"net/http"
	"userms/api/v1/model"

	"github.com/gin-gonic/gin"
)

func Base(ctx *gin.Context) {
	body := model.DefaultResponse{
		Code: http.StatusOK,
		Message: "",
		Body: map[string]string{
			"information":"infind user service",
			"version":"0.0.1",
		},
	}

	ctx.JSON(http.StatusOK, body)
}

func None(ctx *gin.Context) {
	body := model.DefaultResponse{
		Code: http.StatusNotFound,
		Message: "That endpoint does not exist",
		Body: map[string]string{},
	}

	ctx.JSON(http.StatusNotFound, body)
}
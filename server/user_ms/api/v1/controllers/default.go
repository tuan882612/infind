package controllers

import (
	"net/http"
	"userms/api/v1/model"

	"github.com/gin-gonic/gin"
)

func Base() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data := map[string]string{
			"information":"infind user service",
			"version":"0.0.1",
		}

		body := model.DefaultResponse{
			Code: http.StatusOK,
			Message: "",
			Body: data,
		}

		ctx.JSON(http.StatusOK, body)
	}
}

func None() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusNotFound)
	}
}
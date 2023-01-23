package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Base() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data := map[string]string{
			"information":"infind user service",
			"version":"0.0.1",
		}

		ctx.JSON(http.StatusOK, data)
	}
}

func None() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusNotFound)
	}
}
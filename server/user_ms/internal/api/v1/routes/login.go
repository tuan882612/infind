package routes

import (
	"userms/internal/domain/login"
	h "userms/internal/api/v1/handlers"

	"github.com/gin-gonic/gin"
)

func LoginRoutes(router *gin.RouterGroup) {
	handler := &h.LoginController{Repo: login.NewRepo()}

	router.GET("/login", handler.Login)
	router.POST("/register", handler.Register)
}
package routes

import (
	"userms/internal/domain/user"
	h "userms/internal/api/v1/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {
	handler := &h.UserController{Repo: user.NewRepo()}
	
	user := router.Group("/user") 
	user.GET("/find", handler.Find)
	user.PUT("/update", handler.Update)
	user.DELETE("/delete", handler.Delete)
}
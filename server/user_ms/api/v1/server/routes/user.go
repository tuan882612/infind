package routes

import (
	"userms/api/v1/domain/user"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) *gin.Engine {
	handler := &user.UserController{
		Repo: user.NewRepo(),
	}
	
	user := router.Group("/api/v1/user") 
	{
		user.GET("/find", handler.Find)
		user.PUT("/update", handler.Update)
		user.DELETE("/delete", handler.Delete)
	}

	return router
}
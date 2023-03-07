package v1

import (
	"strconv"
	"user_ms/internal/api/v1/handlers"
	"user_ms/internal/api/v1/routes"
	"user_ms/internal/security"

	"github.com/gin-gonic/gin"
)

type InfindServer struct {
	Engine  *gin.Engine
	Address string
	Port	int
}

func InitService(addres string, port int) *InfindServer {
	return &InfindServer{gin.Default(),addres,port}
}

func (is *InfindServer) Run() {
	url := is.Address+":"+strconv.Itoa(is.Port)
	is.Engine.Run(url)
}

func (is *InfindServer) LoadConfig() {
	is.LoadRoutes()
	security.SetTrustedProxies(is.Engine)
	security.MiddleWare(is.Engine)
}

func (is *InfindServer) LoadRoutes() {
	v1 := is.Engine.Group("/api/v1")
	routes.Default(v1)
	routes.LoginRoutes(v1)
	routes.UserRoutes(v1)
	is.Engine.NoRoute(handlers.None)
}
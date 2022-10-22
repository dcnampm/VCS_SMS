package routes

import (
	"github.com/dcnampm/VCS_SMS.git/controllers"
	"github.com/dcnampm/VCS_SMS.git/middleware"
	"github.com/gin-gonic/gin"
)

type ServerRouteController struct {
	serverController controllers.ServerController
}

func NewRouteServerController(serverController controllers.ServerController) ServerRouteController {
	return ServerRouteController{serverController}
}

func (rc *ServerRouteController) ServerRoute(rg *gin.RouterGroup) {

	router := rg.Group("servers")
	router.Use(middleware.DeserializeUser())
	router.POST("/", rc.serverController.CreateServer)
	router.GET("/:serverID", rc.serverController.ViewAServer)
	router.GET("/", rc.serverController.ViewAllServers)
	router.GET("/", rc.serverController.FilterServers)
	router.PUT("/:serverId", rc.serverController.UpdateServer)
	router.DELETE("/:serverId", rc.serverController.DeleteServer)
}

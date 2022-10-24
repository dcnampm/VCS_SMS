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
	router.GET("/view/sort", rc.serverController.ViewAndSortServers)
	router.GET("/view/filter", rc.serverController.ViewAndFilterServers)
	router.PUT("/:serverID", rc.serverController.UpdateServer)
	router.DELETE("/:serverID", rc.serverController.DeleteServer)
}

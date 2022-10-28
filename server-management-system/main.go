package main

import (
	"log"
	"net/http"

	"github.com/dcnampm/VCS_SMS.git/controllers"
	"github.com/dcnampm/VCS_SMS.git/initializers"
	"github.com/dcnampm/VCS_SMS.git/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine

	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	ServerController      controllers.ServerController
	ServerRouteController routes.ServerRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	ServerController = controllers.NewServerController(initializers.DB)
	ServerRouteController = routes.NewRouteServerController(ServerController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "VCS_SMS"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	go func() {
		ServerController.DailyReport()
	}()

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	ServerRouteController.ServerRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}

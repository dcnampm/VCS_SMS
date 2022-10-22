package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dcnampm/VCS_SMS.git/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ServerController struct {
	DB *gorm.DB
}

func NewServerController(DB *gorm.DB) ServerController {
	return ServerController{DB}
}

//Create Server
func (sc *ServerController) CreateServer(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User) //need to reference the User_id creating server
	var payload *models.CreateNewServer

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newServer := models.Server{
		Server_name:  payload.Server_name,
		User_id:      currentUser.User_id,
		Status:       payload.Status,
		Created_time: now,
		Last_updated: now,
		Ipv4:         payload.Ipv4,
	}

	result := sc.DB.Create(&newServer)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Server_name already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}
}

//View A Server
func (sc *ServerController) ViewAServer(ctx *gin.Context) {
	serverID := ctx.Param("serverID")

	var server models.Server
	result := sc.DB.First(&server, "server_id = ?", serverID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Server not exist"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": server})
}

//View All Servers
func (sc *ServerController) ViewAllServers(ctx *gin.Context) {
	//func (c *Context) DefaultQuery(key, defaultValue string) string
	// returns the keyed url query value if it exists, otherwise it returns the specified defaultValue string
	var page = ctx.DefaultQuery("page", "1") //phân trang, default 1 trang có 10 servers
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	//Sort
	var sortField = ctx.DefaultQuery("sortField", "server_name") //default sort theo name

	var servers []models.Server
	results := sc.DB.Order(sortField).Limit(intLimit).Offset(offset).Find(&servers)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(servers), "data": servers})
}

//Filter Servers
func (sc *ServerController) FilterServers(ctx *gin.Context) {
	var filterBy = ctx.DefaultQuery("filterBy", "status")
	var filterResponse = ctx.DefaultQuery("filterResponse", "Off")

	var servers []models.Server
	results := sc.DB.Where(filterBy, filterResponse).Find(&servers)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(servers), "data": servers})
}

//Update Server
func (sc *ServerController) UpdateServer(ctx *gin.Context) {
	serverID := ctx.Param("serverID") //Param returns the value of the URL param
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.UpdateServer //nhap

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedServer models.Server
	result := sc.DB.First(&updatedServer, "server_id = ?", serverID)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Server not exist"})
		return
	}
	now := time.Now()
	serverUpdate := models.Server{
		Server_name:  payload.Server_name,
		User_id:      currentUser.User_id,
		Status:       payload.Status,
		Created_time: now,
		Last_updated: now,
		Ipv4:         payload.Ipv4,
	}

	sc.DB.Model(&updatedServer).Updates(serverUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedServer})
}

//Delete Server
func (sc *ServerController) DeleteServer(ctx *gin.Context) {
	serverID := ctx.Param("serverID")

	result := sc.DB.Delete(&models.Server{}, "server_id = ?", serverID)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Server not exist"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Delete Completed"})
}

package controllers

import

type ServerController struct {
	DB *gorm.DB
}

func NewServerController(DB *gorm.DB) ServerController {
	return ServerController{DB}
}

//Create Server 
func (sc *ServerController) CreateServer(ctx *gin.Context) {
	currentServer := ctx.MustGet("currentServer").(models.Server)
	var payload *models.CreateServerRequest
}
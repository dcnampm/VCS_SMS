package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dcnampm/VCS_SMS.git/models"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
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

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newServer})
}

//View And Sort Servers
func (sc *ServerController) ViewAndSortServers(ctx *gin.Context) {
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

//View And Filter Servers
func (sc *ServerController) ViewAndFilterServers(ctx *gin.Context) {
	var filterBy = ctx.DefaultQuery("filterBy", "status")
	var filterRequest = ctx.DefaultQuery("filterRequest", "Off")

	var servers []models.Server
	results := sc.DB.Where(filterBy, filterRequest).Find(&servers)
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
	var payload *models.UpdateServer

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var server models.Server
	result := sc.DB.First(&server, "server_id = ?", serverID)

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

	sc.DB.Model(&server).Updates(serverUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": serverUpdate})
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

//ExportExcel
func (sc *ServerController) ExportExcel(ctx *gin.Context) {
	f := excelize.NewFile()

	// Create a new sheet.
	// index := f.NewSheet("Sheet1")

	// Set value of a cell
	f.SetCellValue("Sheet1", "A1", "Server_id")
	f.SetCellValue("Sheet1", "B1", "Server_name")
	f.SetCellValue("Sheet1", "C1", "User_id")
	f.SetCellValue("Sheet1", "D1", "Status")
	f.SetCellValue("Sheet1", "E1", "Created_time")
	f.SetCellValue("Sheet1", "F1", "Last_updated")
	f.SetCellValue("Sheet1", "G1", "Ipv4")

	var servers []models.Server
	sc.DB.Find(&servers)
	for i, r := range servers {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+2), r.Server_id)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+2), r.Server_name)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+2), r.User_id)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(i+2), r.Status)
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(i+2), r.Created_time)
		f.SetCellValue("Sheet1", "G"+strconv.Itoa(i+2), r.Last_updated)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(i+2), r.Ipv4)

	}

	if err := f.SaveAs("ExportServer.xlsx"); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "Failed to export DB to the excel", "error": err})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success"})
}

//ImportExcel
// func (sc *ServerController) ImportExcel(ctx *gin.Context) {
// 	var servers []models.Server
// 	sc.DB.Offset(0).Find(&servers)

// 	f, err := excelize.OpenFile("ImportServer.xlsx")
// 	if err != nil {
// 		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "Failed to import Db to excel", "error": err})
// 		return
// 	}

// 	rows, err := f.GetRows("Sheet1")
// 	if err != nil {
// 		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
// 		return
// 	}

// 	now := time.Now()

// 	serversImport := make([]models.Server, 0)

// 	serversAccept := make([]models.ImportExcel, 0)
// 	serversFail := make([]models.ImportExcel, 0)

// 	if len(servers) != 0 {
// 		for _, r := range servers {
// 			for _, row := range rows {
// 				if len(row) != 0 {
// 					if r.Server_id == row[0] || r.Server_name == row[1] {
// 						newServerFail := models.ImportExcel{
// 							Server_id:   row[0],
// 							Server_name: row[1],
// 						}
// 						serversFail = append(serversFail, newServerFail)
// 						continue
// 					}
// 					user_id, _ := strconv.Atoi(row[4])
// 					newServer := models.Server{
// 						Server_id:    row[0],
// 						Server_name:  row[1],
// 						Status:       row[2],
// 						Ipv4:         row[3],
// 						User_id:      user_id,
// 						Created_time: now,
// 						Last_updated: now,
// 					}
// 					serversImport = append(serversImport, newServer)

// 					newServerAccept := models.ImportExcel{
// 						Server_id:   row[0],
// 						Server_name: row[1],
// 					}
// 					serversAccept = append(serversAccept, newServerAccept)
// 				}
// 			}
// 		}
// 	} else {
// 		for _, row := range rows {
// 			if len(row) != 0 {
// 				user_id, _ := strconv.Atoi(row[4])
// 				newServer := models.Server{
// 					Server_id:    row[0],
// 					Server_name:  row[1],
// 					Status:       row[2],
// 					Ipv4:         row[3],
// 					User_id:      user_id,
// 					Created_time: now,
// 					Last_updated: now,
// 				}
// 				serversImport = append(serversImport, newServer)

// 				newServerAccept := models.ImportExcel{
// 					Server_id:   row[0],
// 					Server_name: row[1],
// 				}
// 				serversAccept = append(serversAccept, newServerAccept)
// 			}
// 		}
// 	}
// 	results := sc.DB.Create(&serversImport)

// 	if results.Error != nil {
// 		ctx.JSON(http.StatusOK, gin.H{"status": "fail", "message": results.Error.Error()})
// 		return
// 	}
// 	ctx.JSON(http.StatusCreated, gin.H{"status": gin.H{"ImportEccept": gin.H{"CountAccept": len(serversAccept), "data": serversAccept}, "ImportFail": gin.H{"CountFail": len(serversFail), "data": serversFail}}})
// }

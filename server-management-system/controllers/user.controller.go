package controllers

import (
	"net/http"

	"github.com/dcnampm/VCS_SMS.git/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}

func (uc *UserController) GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	userResponse := &models.UserResponse{
		User_id:    currentUser.User_id,
		User_name:  currentUser.User_name,
		User_email: currentUser.User_email,
		CreatedAt:  currentUser.CreatedAt,
		UpdatedAt:  currentUser.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}

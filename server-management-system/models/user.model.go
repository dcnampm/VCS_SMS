package models

import (
	"time"
)

type User struct {
	User_id    uint   `gorm:"type:serial;primary_key"`
	User_name  string `gorm:"type:varchar(255);not null"`
	User_email string `gorm:"uniqueIndex;not null"`
	Password   string `gorm:"not null"`
	Verified   bool   `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type SignUp struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}

type SignIn struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`
}

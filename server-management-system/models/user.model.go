package models

import (
	"time"
)

type User struct {
	User_id    int    `gorm:"type:serial;primary_key"`
	User_name  string `gorm:"type:varchar(255);not null"`
	User_email string `gorm:"uniqueIndex;not null"`
	Password   string `gorm:"not null"`
	Verified   bool   `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type SignUp struct {
	User_name       string `json:"user_name" binding:"required"`
	User_email      string `json:"user_email" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}

type SignIn struct {
	User_email string `json:"user_email"  binding:"required"`
	Password   string `json:"password"  binding:"required"`
}

type UserResponse struct {
	User_id    int       `json:"user_id"`
	User_name  string    `json:"user_name"`
	User_email string    `json:"user_email"`
	CreatedAt  time.Time `json:"create_at"`
	UpdatedAt  time.Time `json:"update_at"`
}

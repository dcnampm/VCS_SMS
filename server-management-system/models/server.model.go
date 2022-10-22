package models

import "time"

type Server struct {
	Server_id    uint      `gorm:"type:serial" json:"server_id,omitempty"`
	Server_name  string    `gorm:"uniqueIndex;not null" json:"server_name,omitempty"`
	User_id      uint      `gorm:"not null" json:"user_id,omitempty"`
	Status       string    `gorm:"not null" json:"server_status,omitempty"`
	Created_time time.Time `gorm:"not null" json:"created_time,omitempty"`
	Last_updated time.Time `gorm:"not null" json:"last_updated,omitempty"`
	Ipv4         string    `gorm:"not null" json:"ipv4,omitempty"`
}

type CreateNewServer struct {
	Server_name  string    `json:"server_name" binding:"required"`
	User_id      uint      `json:"user_id" binding:"required"`
	Status       string    `json:"status" binding:"required"`
	Created_time time.Time `json:"created_time,omitempty"`
	Last_updated time.Time `json:"last_updated,omitempty"`
	Ipv4         string    `json:"ipv4,omitempty"`
}

type UpdateServer struct {
	Server_name  string    `json:"server_name,omitempty"`
	User_id      uint      `json:"user_id,omitempty"`
	Status       string    `json:"status,omitempty"`
	Created_time time.Time `json:"created_time,omitempty"`
	Last_updated time.Time `json:"last_updated,omitempty"`
	Ipv4         string    `json:"ipv4,omitempty"`
}

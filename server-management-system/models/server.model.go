package models

import "time"

type Server struct {
	Server_id    string    `gorm:"type:serial;primary_key" json:"server_id,omitempty"`
	Server_name  string    `gorm:"uniqueIndex;not null" json:"server_name,omitempty"`
	User_id      int       `gorm:"not null" json:"user_id,omitempty"`
	Status       string    `gorm:"not null" json:"status,omitempty"`
	Created_time time.Time `gorm:"not null" json:"created_time,omitempty"`
	Last_updated time.Time `gorm:"not null" json:"last_updated,omitempty"`
	Ipv4         string    `gorm:"not null" json:"ipv4,omitempty"`
	Uptime       float64   `gorm:"not null" json:"uptime,omitempty"`
}

type CreateNewServer struct {
	Server_name  string    `json:"server_name" binding:"required"`
	User_id      int       `json:"user_id"`
	Status       string    `json:"status" binding:"required"`
	Created_time time.Time `json:"created_time,omitempty"`
	Last_updated time.Time `json:"last_updated,omitempty"`
	Ipv4         string    `json:"ipv4,omitempty" binding:"required"`
	Uptime       float64   `json:"uptime,omitempty"`
}

type UpdateServer struct {
	Server_name  string    `json:"server_name,omitempty"`
	User_id      int       `json:"user_id,omitempty"`
	Status       string    `json:"status,omitempty"`
	Created_time time.Time `json:"created_time,omitempty"`
	Last_updated time.Time `json:"last_updated,omitempty"`
	Ipv4         string    `json:"ipv4,omitempty"`
	Uptime       float64   `json:"uptime,omitempty"`
}

type ImportExcel struct {
	Server_id   string ` json:"server_id,omitempty"`
	Server_name string `json:"server_name,omitempty"`
}

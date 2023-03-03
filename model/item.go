package model

import (
	"gorm.io/gorm"
)

type LostItem struct {
	gorm.Model
	Uuid        string `json:"uuid"`
	User_id     int    `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image_url   string `json:"Image_url"`
	Deadline    string `json:"Deadline"`
}

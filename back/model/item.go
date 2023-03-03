package model

import (
	"time"

	"gorm.io/gorm"
)

type LostItem struct {
	gorm.Model
	Name      string    `json:"name"`
	Kind      string    `json:"kind"`
	Feature   string    `json:"feature"`
	Comment   string    `json:"comment"`
	ImageUrl string    `json:"image_url"`
	// todo Locationを追加する
	FindTime time.Time `json:"find_time"`
}

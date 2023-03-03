package model

import (
	"time"

	"gorm.io/gorm"
)

type LostItem struct {
	gorm.Model
	Name      string    `json:"name"`
	Kind      string    `json:"kind"`
	feature   string    `json:"feature"`
	Comment   string    `json:"comment"`
	Image_url string    `json:"image_url"`
	// todo Locationを追加する
	find_time time.Time `json:"find_time"`
}

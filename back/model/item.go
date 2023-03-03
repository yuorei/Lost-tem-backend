package model

import (
	"time"

	"gorm.io/gorm"
)

type LostItem struct {
	gorm.Model
	KindID       uint      `json:"kindID"`
	Feature      string    `json:"feature"`
	Comment      string    `json:"comment"`
	ImageURL     string    `json:"imageURL"`
	Location     Location  `json:"location"`
	FindTime     time.Time `json:"find_time"`
	CompleteTime time.Time `json:"completeTime`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

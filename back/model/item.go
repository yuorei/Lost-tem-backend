package model

import (
	"time"

	"gorm.io/gorm"
)

type LostItem struct {
	gorm.Model
	KindID       uint   `json:"kindID"`
	Feature      string `json:"feature"`
	Comment      string `json:"comment"`
	ImageUrl     string `json:"imageURL"`
	Location     Location
	FindTime     time.Time `json:"find_time"`
	CompleteTime time.Time `json:"completeTime`
}

type Location struct {
	lat float64 `json:"lat`
	lng float64 `json:"lat`
}

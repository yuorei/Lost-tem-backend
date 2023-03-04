package database

import (
	"time"

	"gorm.io/gorm"
)

type LostItem struct {
	gorm.Model
	Kinds        string
	Feature      string
	Comment      string
	ImageURL     string
	Lat          float64
	Lng          float64
	FindTime     time.Time
	CompleteTime time.Time
}

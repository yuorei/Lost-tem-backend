package database

import (
	"time"

	"gorm.io/gorm"
)

type LostItem struct {
	gorm.Model
	Kinds    string    `gorm:"not null"`
	Comment  string    `gorm:"not null"`
	ImageURL string    `gorm:"not null"`
	Lat      float64   `gorm:"not null"`
	Lng      float64   `gorm:"not null"`
	FindTime time.Time `gorm:"not null"`
}

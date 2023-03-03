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
	FindTime     time.Time `json:"findTime"`
	CompleteTime time.Time `json:"completeTime"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type AreaSearchQuery struct {
	Location1 Location `json:"location1"`
	Location2 Location `json:"location2"`
}

type SearchResult struct {
	Count uint `json:"count"`
	Items []struct {
		Id  uint    `json:"id"`
		Lat float64 `json:"lat"`
		Lng float64 `json:"Lng"`
	} `json:"items"`
}

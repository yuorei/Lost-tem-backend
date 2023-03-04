package model

import "time"

type LostItem struct {
	ID       uint      `json:"id"`
	Kinds    []string  `json:"tags"`
	Comment  string    `json:"note"`
	ImageURL string    `json:"pic"`
	Location Location  `json:"location"`
	FindTime time.Time `json:"date" time_format:"2006-01-02T15:04:05Z"`
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
	Count uint       `json:"count"`
	Items []LostItem `json:"items"`
}

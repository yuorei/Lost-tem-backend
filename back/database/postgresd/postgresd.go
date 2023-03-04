package postgresd

import (
	"fmt"
	"lost-item/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgresd struct {
	conn  *gorm.DB
	limit uint
}

func NewPostgresd() (*Postgresd, error) {
	POSTGRES_HOST := os.Getenv("POSTGRES_HOST")
	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_DB := os.Getenv("POSTGRES_DB")
	TZ := os.Getenv("TZ")
	dsn := "host=" + POSTGRES_HOST + " user=" + POSTGRES_USER + " password=" + POSTGRES_PASSWORD + " dbname=" + POSTGRES_DB + " port=5432 sslmode=disable TimeZone=" + TZ

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	var limit uint = 100

	return &Postgresd{
		conn:  db,
		limit: limit,
	}, err
}

func (d *Postgresd) SearchItemsFor(query string) (model.SearchResult, error) {
	items := make([]model.LostItem, d.limit)
	err := d.conn.Where("Kinds LIKE", fmt.Sprintf("%%%s%%", query)).Limit(int(d.limit)).Find(&items).Error

	if err != nil {
		return model.SearchResult{}, err
	}

	return model.SearchResult{
		Count: uint(len(items)),
		Items: items,
	}, nil
}

func (d *Postgresd) SearchItemsArea(left_upper model.Location, right_bottom model.Location) (model.SearchResult, error) {
	items := make([]model.LostItem, d.limit)
	err := d.conn.Preload("Location").Where(
		"Lat <= ? AND Lat >= ? AND Lng <= ? AND Lng >= ?",
		left_upper.Lat, right_bottom.Lat, right_bottom.Lng, left_upper.Lng,
	).Limit(int(d.limit)).Find(&items).Error

	if err != nil {
		return model.SearchResult{}, err
	}

	return model.SearchResult{
		Count: uint(len(items)),
		Items: items,
	}, nil
}

func (d *Postgresd) ItemDetail(id uint64) (model.LostItem, error) {
	item := model.LostItem{}
	err := d.conn.Where("id = ?", id).First(&item).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.LostItem{
				Model: gorm.Model{ID: 0},
			}, nil
		}
		return model.LostItem{}, err
	}

	return item, nil
}

package postgresd

import (
	"log"
	"lost-item/database"

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

func toModelLostItem(item database.LostItem) model.LostItem {
	return model.LostItem{
		ID:       item.ID,
		Kinds:    item.Kinds,
		Comment:  *item.Comment,
		ImageURL: item.ImageURL,
		Location: model.Location{
			Lat: item.Lat,
			Lng: item.Lng,
		},
		FindTime: item.FindTime,
	}
}

func mapToModelLostItem(items []database.LostItem) []model.LostItem {
	mapped_items := make([]model.LostItem, len(items))
	for idx, item := range items {
		mapped_items[idx] = toModelLostItem(item)
	}

	return mapped_items
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

func (d *Postgresd) CreateTable() {
	if err := d.conn.AutoMigrate(&database.LostItem{}); err != nil {
		log.Fatalf("Database create table failed")
	}
}

func (d *Postgresd) Search(left_upper model.Location, right_bottom model.Location, query string, tags []string) (model.SearchResult, error) {
	query_string := "Lat <= ? AND Lat >= ? AND Lng <= ? AND Lng >= ?"
	if query != "" {
		query_string = fmt.Sprintf("%s AND Comment LIKE %%%s%%", query_string, query)
	}
	for _, tag := range tags {
		query_string = fmt.Sprintf("%s AND Kinds LIKE %%%s%%", query_string, tag)
	}

	items := make([]database.LostItem, d.limit)
	err := d.conn.Where(
		query_string,
		left_upper.Lat, right_bottom.Lat, right_bottom.Lng, left_upper.Lng,
	).Limit(int(d.limit)).Find(&items).Error

	if err != nil {
		return model.SearchResult{}, err
	}

	return model.SearchResult{
		Count: uint(len(items)),
		Items: mapToModelLostItem(items),
	}, nil
}

func (d *Postgresd) ItemDetail(id uint64) (model.LostItem, error) {
	item := database.LostItem{}
	err := d.conn.Where("id = ?", id).First(&item).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.LostItem{
				ID: 0,
			}, nil
		}
		return model.LostItem{}, err
	}

	return toModelLostItem(item), nil
}

func (d *Postgresd) CompleteItem(id uint64) error {
	err := d.conn.Where("id = ?", id).Delete(&database.LostItem{}).Error
	return err
}

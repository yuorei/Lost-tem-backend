package postgresd

import (
	"log"
	"lost-item/database"
	"lost-item/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgresd struct {
	conn *gorm.DB
}

func NewPostgresd() (*Postgresd, error) {
	POSTGRES_HOST := os.Getenv("POSTGRES_HOST")
	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_DB := os.Getenv("POSTGRES_DB")
	TZ := os.Getenv("TZ")
	dsn := "host=" + POSTGRES_HOST + " user=" + POSTGRES_USER + " password=" + POSTGRES_PASSWORD + " dbname=" + POSTGRES_DB + " port=5432 sslmode=disable TimeZone=" + TZ

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return &Postgresd{
		conn: db,
	}, err
}

func (d *Postgresd) CreateTable() {
	if err := d.conn.AutoMigrate(&database.LostItem{}); err != nil {
		log.Fatalf("Database create table failed")
	}
}

func (d *Postgresd) SearchItemsFor(query string) (model.SearchResult, error) {
	return model.SearchResult{}, nil
}

func (d *Postgresd) SearchItemsArea(left_upper model.Location, right_bottom model.Location) (model.SearchResult, error) {
	return model.SearchResult{}, nil
}

func (d *Postgresd) ItemDetail(id uint64) (model.LostItem, error) {
	return model.LostItem{}, nil
}

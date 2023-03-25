package postgresd

import (
	"log"
	"lost-item/database"
	"strings"
	"time"

	"lost-item/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Postgresd struct {
	conn  *gorm.DB
	limit uint
}

func toModelLostItem(item database.LostItem) model.LostItem {
	return model.LostItem{
		ID:       item.ID,
		Kinds:    strings.Split(item.Kinds, ","),
		Comment:  item.Comment,
		ImageURL: item.ImageURL,
		Location: model.Location{
			Lat: item.Lat,
			Lng: item.Lng,
		},
		FindTime:  item.FindTime,
		ItemName:  item.ItemName,
		Colour:    item.Colour,
		Situation: item.Situation,
		Others:    item.Others,
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

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})

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
	left := left_upper.Lng
	right := right_bottom.Lng
	upper := left_upper.Lat
	bottom := right_bottom.Lat

	swap := func(x, y float64) (float64, float64) { return y, x }
	if left > right {
		left, right = swap(left, right)
	}
	if bottom > upper {
		bottom, upper = swap(bottom, upper)
	}

	db := d.conn.Where("? <= Lat AND Lat <= ? AND ? <= Lng AND Lng <= ?", bottom, upper, left, right)
	if query != "" {
		db = d.conn.Where("Comment LIKE ?", "%"+query+"%")
		db = d.conn.Where("Kinds LIKE ?", "%"+query+"%")
		db = d.conn.Where("colour LIKE ?", "%"+query+"%")
		db = d.conn.Where("situation LIKE ?", "%"+query+"%")
		db = d.conn.Where("others LIKE ?", "%"+query+"%")
	}
	for _, tag := range tags {
		db = d.conn.Where("Kinds LIKE ?", "%"+tag+"%")
	}

	items := make([]database.LostItem, d.limit)
	db = db.Limit(int(d.limit)).Find(&items)
	if db.RowsAffected == 0 {
		return model.SearchResult{
			Count: 0,
			Items: []model.LostItem{},
		}, nil
	}
	err := db.Error

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

func (d *Postgresd) InsertItem(item model.LostItem) (model.LostItem, error) {
	item_db := database.LostItem{
		Model:    gorm.Model{ID: item.ID},
		Kinds:    strings.Join(item.Kinds, ","),
		Comment:  item.Comment,
		ImageURL: item.ImageURL,
		Lat:      item.Location.Lat,
		Lng:      item.Location.Lng,
		FindTime: item.FindTime,
	}

	if err := d.conn.Create(&item_db).Error; err != nil {
		return item, err
	}
	item.ID = item_db.ID
	return item, nil
}

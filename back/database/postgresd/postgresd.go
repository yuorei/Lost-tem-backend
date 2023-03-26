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

	var db *gorm.DB
	// queryが存在する場合
	if query != "" {
		// 空白分け
		queries := strings.Fields(query)
		db = d.conn
		for _, q := range queries {
			db = db.Where("comment LIKE ? OR others LIKE ? OR Situation LIKE ? OR Kinds LIKE ?", "%"+q+"%", "%"+q+"%", "%"+q+"%", "%"+q+"%")
		}
	}

	items := make([]database.LostItem, d.limit)
	if query != "" {
		db = db.Where("? <= Lat AND Lat <= ? AND ? <= Lng AND Lng <= ?", bottom, upper, left, right).Limit(int(d.limit)).Find(&items)
	} else {
		db = d.conn.Where("? <= Lat AND Lat <= ? AND ? <= Lng AND Lng <= ?", bottom, upper, left, right).Limit(int(d.limit)).Find(&items)
	}
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
		Model:     gorm.Model{ID: item.ID},
		Kinds:     strings.Join(item.Kinds, ","),
		Comment:   item.Comment,
		ImageURL:  item.ImageURL,
		Lat:       item.Location.Lat,
		Lng:       item.Location.Lng,
		FindTime:  item.FindTime,
		ItemName:  item.ItemName,
		Colour:    item.Colour,
		Situation: item.Situation,
		Others:    item.Others,
	}

	if err := d.conn.Create(&item_db).Error; err != nil {
		return item, err
	}
	item.ID = item_db.ID
	return item, nil
}

func (d *Postgresd) UpdateItem(id uint64, item model.UpdateLostItem) (model.LostItem, error) {
	var item_db database.LostItem
	if err := d.conn.Where("id = ?", id).First(&item_db).Error; err != nil {
		return model.LostItem{}, err
	}

	if len(item.Kinds) > 0 {
		item_db.Kinds = strings.Join(item.Kinds, ",")
	}

	if item.Comment != "" {
		item_db.Comment = item.Comment
	}

	if item.ItemName != "" {
		item_db.ItemName = item.ItemName
	}

	if item.Colour != "" {
		item_db.Colour = item.Colour
	}

	if item.Situation != "" {
		item_db.Situation = item.Situation
	}

	if item.Others != "" {
		item_db.Others = item.Others
	}

	if item.Location != nil {
		item_db.Lat = item.Location.Lat
		item_db.Lng = item.Location.Lng
	}

	if item.FindTime != nil {
		item_db.FindTime = *item.FindTime
	}

	if err := d.conn.Save(&item_db).Error; err != nil {
		return model.LostItem{}, err
	}

	return toModelLostItem(item_db), nil
}

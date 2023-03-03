package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Association struct {
	DB    *gorm.DB
	Error error
}

// DBと接続をします
func (a Association) Open(){
	POSTGRES_HOST := os.Getenv("POSTGRES_HOST")
	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_DB := os.Getenv("POSTGRES_DB")
	TZ := os.Getenv("TZ")
	dsn := "host=" + POSTGRES_HOST + " user=" + POSTGRES_USER + " password=" + POSTGRES_PASSWORD + " dbname=" + POSTGRES_DB + " port=5432 sslmode=disable TimeZone=" + TZ
	a.DB, a.Error = gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
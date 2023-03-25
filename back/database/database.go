package database

import "lost-item/model"

type DBConn interface {
	CreateTable()
	Search(model.Location, model.Location, string, []string) (model.SearchResult, error)
	ItemDetail(uint64) (model.LostItem, error)
	InsertItem(model.LostItem) (model.LostItem, error)
	CompleteItem(uint64) error
}

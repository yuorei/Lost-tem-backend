package database

import "lost-item/model"

type DBConn interface {
	CreateTable()
	SearchItemsFor(string) (model.SearchResult, error)
	SearchItemsArea(model.Location, model.Location) (model.SearchResult, error)
	ItemDetail(uint64) (model.LostItem, error)
	InsertItem(model.LostItem) (LostItem, error)
}

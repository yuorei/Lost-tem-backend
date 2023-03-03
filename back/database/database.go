package database

import "lost-item/model"

type DBConn struct{}

func NewDBConn(url string, port uint16) (error, *DBConn) {
	return nil, &DBConn{}
}

func (d *DBConn) SearchItemsFor(query string) (error, model.SearchResult) {

}

package database

import "lost-item/model"

type DBConn struct{}

func NewDBConn(url string, port uint16) (*DBConn, error) {
	return nil, &DBConn{}
}

func (d *DBConn) SearchItemsFor(query string) (model.SearchResult, error) {

}

func (d *DBConn) SearchItemsArea(left_upper model.Location, right_bottom model.Location) (model.SearchResult, error) {

}

func (d *DBConn) ItemDetail(id uint64) (model.LostItem, error) {

}

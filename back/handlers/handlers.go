package handlers

import (
	"log"
	"lost-item/database"
	"lost-item/database/postgresd"
	"lost-item/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	db database.DBConn
}

func (h Handler) Init() {
	var err error
	if h.db, err = postgresd.NewPostgresd(); err != nil {
		log.Fatalf("Database connection failed")
	}
}

func (h Handler) Search(c *gin.Context) {
	search_query := c.Param("q")
	search_result, err := h.db.SearchItemsFor(search_query)

	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusOK, search_result)
}

func (h Handler) ItemList(c *gin.Context) {
	search_query := model.AreaSearchQuery{}
	err := c.Bind(&search_query)

	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}

	search_result, err := h.db.SearchItemsArea(search_query.Location1, search_query.Location2)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusOK, search_result)
}

func (h Handler) ItemDetail(c *gin.Context) {
	item_id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}

	item_detail, err := h.db.ItemDetail(uint64(item_id))
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusOK, item_detail)
}

func (h Handler) RegisterItem(c *gin.Context) {
	register_item:=model.LostItem{}
	err := c.Bind(&register_item)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}
	// TODO insert
	c.JSON(http.StatusOK, register_item)

}

func (h Handler) DeleteItem(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	// TODO delete
	
	delete_item, err := h.db.ItemDetail(id)
	c.JSON(http.StatusOK, delete_item)
}

func (h Handler) parse(c *gin.Context) {

}

func (h Handler) RegisterImage(c *gin.Context) {

}

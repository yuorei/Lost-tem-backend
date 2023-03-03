package handlers

import (
	"log"
	"lost-item/database"
	"lost-item/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context, db *database.DBConn) {
	search_query := c.Param("q")
	search_result, err := db.SearchItemsFor(search_query)

	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, search_result)
}

func ItemList(c *gin.Context, db *database.DBConn) {
	search_query := model.AreaSearchQuery{}
	err := c.Bind(&search_query)

	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}

	search_result, err := db.SearchItemsArea(search_query.Location1, search_query.Location2)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, search_result)
}

func ItemDetail(c *gin.Context, db *database.DBConn) {
	item_id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}

	item_detail, err := db.ItemDetail(uint64(item_id))
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, item_detail)
}

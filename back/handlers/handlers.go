package handlers

import (
	"log"
	"lost-item/database"
	"lost-item/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context, db *database.DBConn) {
	search_query := c.Param("q")
	err, search_result := db.SearchItemsFor(search_query)

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

	err, search_result := db.SearchItemsArea(search_query.Location1, search_query.Location2)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, search_result)
}

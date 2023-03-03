package handlers

import (
	"log"
	"lost-item/database"
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

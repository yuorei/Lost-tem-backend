package handlers

import "github.com/gin-gonic/gin"

func Search(c *gin.Context) {
	search_query := c.Param("q")

}

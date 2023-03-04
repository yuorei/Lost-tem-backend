package router

import (
	"lost-item/handlers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	h := handlers.Handler{}
	h.Init()
	r.GET("/item_list", h.ItemList)
	r.GET("/item", h.ItemDetail)
	r.POST("/item", h.RegisterItem)
	r.DELETE("/item", h.DeleteItem)
	r.POST("/parse", h.Search)
	r.POST(" /image", h.RegisterImage)
	return r
}

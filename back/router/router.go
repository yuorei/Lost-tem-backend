package router

import (
	"lost-item/handlers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	h := handlers.Handler{}
	h.Init()
	r.GET("/serch",h.Search)
	return r
}

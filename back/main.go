package main

import (
	"lost-item/db"

	"github.com/gin-gonic/gin"
)
// サンプル
type ss struct {
	i int
}
func main() {
	d := &db.Association{}
	d.Open()
	//サンプル
	// テーブル
	d.DB.AutoMigrate(&ss{})
	// insert
	d.DB.Create(&ss{
		i:1,
	})

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":3000")
}

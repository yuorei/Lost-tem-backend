package main

import (
	"lost-item/router"
)

func main() {
	r := router.Router()
	r.Run(":3000")
}

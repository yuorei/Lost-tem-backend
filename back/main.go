package main

import (
	"log"
	"lost-item/router"
	"os"
)

func main() {
	r := router.Router()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	r.Run(":"+port)
}

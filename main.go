package main

import (
	"log"

	db "github.com/nictes1/career-path-server/database"
	"github.com/nictes1/career-path-server/handlers"
)

func main() {

	if db.ConnectionCheck() == 0 {
		log.Fatal("No connection to the database")
		return
	}
	handlers.LaunchingServer()
}

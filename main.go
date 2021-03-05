package main

import (
	"fmt"
	"log"

	"github.com/nictes1/career-path-server/bd"
	"github.com/nictes1/career-path-server/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
	fmt.Println("hola mundo!")
}

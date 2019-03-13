package main

import (
	"log"
	"net/http"

	"github.com/DiscoFighter47/goingTDD/application/data/memory"
	"github.com/DiscoFighter47/goingTDD/application/server"
)

func main() {
	svr := server.NewPlayerServer(memory.NewPlayerStore())
	log.Println("Server starting on port: 8080")
	if err := http.ListenAndServe(":8080", svr); err != nil {
		log.Fatal("unable to serve", err)
	}
}

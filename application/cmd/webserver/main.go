package main

import (
	"log"
	"net/http"

	"github.com/DiscoFighter47/goingTDD/application/data/file"
	"github.com/DiscoFighter47/goingTDD/application/server"
)

const fileName = "game.db.json"

func main() {
	store, _ := file.NewPlayerStoreByPath(fileName)
	svr := server.NewPlayerServer(store)
	log.Println("Server starting on port: 8080")
	if err := http.ListenAndServe(":8080", svr); err != nil {
		log.Fatal("unable to serve", err)
	}
}

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/DiscoFighter47/goingTDD/application/data/file"
	"github.com/DiscoFighter47/goingTDD/application/server"
)

const fileName = "game.db.json"

func main() {
	db, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("error opening file", fileName)
	}

	store, _ := file.NewPlayerStore(db)
	svr := server.NewPlayerServer(store)
	log.Println("Server starting on port: 8080")
	if err := http.ListenAndServe(":8080", svr); err != nil {
		log.Fatal("unable to serve", err)
	}
}

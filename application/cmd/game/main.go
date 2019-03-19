package main

import (
	"os"

	"github.com/DiscoFighter47/goingTDD/application/cli"
	"github.com/DiscoFighter47/goingTDD/application/data/file"
)

const fileName = "game.db.json"

func main() {
	store, _ := file.NewPlayerStoreByPath(fileName)
	cli := cli.NewCLI(store, os.Stdin)
	cli.PlayPoker()
}

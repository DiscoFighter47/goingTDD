package main

import (
	"os"

	"github.com/DiscoFighter47/goingTDD/application/alerter"
	"github.com/DiscoFighter47/goingTDD/application/game"

	"github.com/DiscoFighter47/goingTDD/application/cli"
	"github.com/DiscoFighter47/goingTDD/application/data/file"
)

const fileName = "game.db.json"

func main() {
	store, _ := file.NewPlayerStoreByPath(fileName)
	game := game.NewPoker(alerter.BlindAlerterFunc(alerter.StdOutAlerter), store)
	cli := cli.NewCLI(os.Stdin, os.Stdout, game)
	cli.PlayPoker()
}

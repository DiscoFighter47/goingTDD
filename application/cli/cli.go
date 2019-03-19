package cli

import (
	"bufio"
	"io"
	"strings"

	"github.com/DiscoFighter47/goingTDD/application/data"
)

// CLI struct
type CLI struct {
	playerStore data.PlayerStore
	input       io.Reader
}

// NewCLI returns a new CLI
func NewCLI(store data.PlayerStore, input io.Reader) *CLI {
	return &CLI{
		playerStore: store,
		input:       input,
	}
}

// PlayPoker playes poker
func (c *CLI) PlayPoker() {
	reader := bufio.NewScanner(c.input)
	reader.Scan()
	c.playerStore.RegisterWin(exractWinner(reader.Text()))
}

func exractWinner(input string) string {
	return strings.Replace(input, " scores", "", 1)
}

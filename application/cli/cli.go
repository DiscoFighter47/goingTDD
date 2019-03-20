package cli

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/DiscoFighter47/goingTDD/application/game"
)

// CLI struct
type CLI struct {
	input  *bufio.Scanner
	output io.Writer
	game   game.Game
}

// NewCLI returns a new CLI
func NewCLI(input io.Reader, output io.Writer, game game.Game) *CLI {
	return &CLI{
		input:  bufio.NewScanner(input),
		output: output,
		game:   game,
	}
}

// PlayPoker playes poker
func (c *CLI) PlayPoker() {
	fmt.Fprint(c.output, "Please enter the number of players: ")
	numPlayer, err := strconv.Atoi(c.readLine())
	if err != nil {
		return
	}
	c.game.Start(numPlayer)
	input := c.readLine()
	c.game.Finish(exractWinner(input))
}

func (c *CLI) readLine() string {
	c.input.Scan()
	return c.input.Text()
}

func exractWinner(input string) string {
	return strings.Replace(input, " scores", "", 1)
}

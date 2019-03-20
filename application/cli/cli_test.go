package cli_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/DiscoFighter47/goingTDD/application/cli"
)

type spyGame struct {
	start       int
	startCalled bool
	finish      string
}

func (spy *spyGame) Start(numPlayer int) {
	spy.start = numPlayer
	spy.startCalled = true
}

func (spy *spyGame) Finish(player string) {
	spy.finish = player
}

func TestCLI(t *testing.T) {
	t.Run("Zahid wins", func(t *testing.T) {
		input := strings.NewReader("7\nzahid scores")
		output := &bytes.Buffer{}
		game := &spyGame{}
		cli := cli.NewCLI(input, output, game)
		cli.PlayPoker()
		assertOutput(t, output.String(), "Please enter the number of players: ")
		assertStart(t, game.start, 7)
	})

	t.Run("Non numeric value", func(t *testing.T) {
		input := strings.NewReader("zahid scores")
		output := &bytes.Buffer{}
		game := &spyGame{}
		cli := cli.NewCLI(input, output, game)
		cli.PlayPoker()
		assertStartCalled(t, game.startCalled)
	})
}

func assertOutput(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got output '%s' want '%s'", got, want)
	}
}

func assertStart(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got start '%d' want '%d'", got, want)
	}
}

func assertStartCalled(t *testing.T, called bool) {
	t.Helper()
	if called {
		t.Error("game start called")
	}
}

func assertFinish(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got finish '%s' want '%s'", got, want)
	}
}

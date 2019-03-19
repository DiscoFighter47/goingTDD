package cli_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/DiscoFighter47/goingTDD/application/cli"
	"github.com/DiscoFighter47/goingTDD/application/model"
)

type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	league   []model.Player
}

func (stub *StubPlayerStore) GetPlayerScore(name string) (int, bool) {
	val, found := stub.Scores[name]
	return val, found
}

func (stub *StubPlayerStore) RegisterWin(name string) {
	stub.WinCalls = append(stub.WinCalls, name)
}

func (stub *StubPlayerStore) GetLeagueTable() []model.Player {
	return stub.league
}

func assertWinCalls(t *testing.T, got, want []string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}

func TestCLI(t *testing.T) {
	t.Run("Zahid wins", func(t *testing.T) {
		input := strings.NewReader("zahid scores")
		playerStore := &StubPlayerStore{}
		cli := cli.NewCLI(playerStore, input)
		cli.PlayPoker()
		assertWinCalls(t, playerStore.WinCalls, []string{"zahid"})
	})

	t.Run("Tair wins", func(t *testing.T) {
		input := strings.NewReader("tair scores")
		playerStore := &StubPlayerStore{}
		cli := cli.NewCLI(playerStore, input)
		cli.PlayPoker()
		assertWinCalls(t, playerStore.WinCalls, []string{"tair"})
	})
}

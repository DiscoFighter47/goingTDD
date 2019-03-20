package game

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/DiscoFighter47/goingTDD/application/alerter"
	"github.com/DiscoFighter47/goingTDD/application/model"

	"github.com/DiscoFighter47/goingTDD/application/data"
)

type alert struct {
	scheduledAt time.Duration
	amount      int
}

type spyBlindAlerter struct {
	alerts []alert
}

func (spy *spyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	spy.alerts = append(spy.alerts, alert{duration, amount})
}

type stubPlayerStore struct {
	winCalls []string
}

func (stub *stubPlayerStore) GetPlayerScore(name string) (int, bool) {
	return 0, false
}

func (stub *stubPlayerStore) RegisterWin(name string) {
	stub.winCalls = append(stub.winCalls, name)
}

func (stub *stubPlayerStore) GetLeagueTable() []model.Player {
	return nil
}

var dummyStore data.PlayerStore
var dummyBlindAlerter alerter.BlindAlerter

func TestPoker_Start(t *testing.T) {
	t.Run("Game for 5 players", func(t *testing.T) {
		alerter := &spyBlindAlerter{}
		poker := NewPoker(alerter, dummyStore)
		poker.Start(5)

		cases := []alert{
			{0 * time.Second, 100},
			{10 * time.Second, 200},
			{20 * time.Second, 400},
			{30 * time.Second, 600},
			{40 * time.Second, 1000},
			{50 * time.Second, 2000},
			{60 * time.Second, 4000},
			{70 * time.Second, 8000},
		}

		for i, c := range cases {
			t.Run(fmt.Sprintf("%d scheduled for %v", c.amount, c.scheduledAt), func(t *testing.T) {
				if len(alerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, alerter.alerts)
				}
				assertAlert(t, alerter.alerts[i], c)
			})
		}
	})

	t.Run("Game for 7 players", func(t *testing.T) {
		alerter := &spyBlindAlerter{}
		poker := NewPoker(alerter, dummyStore)
		poker.Start(7)

		cases := []alert{
			{0 * time.Second, 100},
			{12 * time.Second, 200},
			{24 * time.Second, 400},
			{36 * time.Second, 600},
			{48 * time.Second, 1000},
			{60 * time.Second, 2000},
			{72 * time.Second, 4000},
			{84 * time.Second, 8000},
		}

		for i, c := range cases {
			t.Run(fmt.Sprintf("%d scheduled for %v", c.amount, c.scheduledAt), func(t *testing.T) {
				if len(alerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, alerter.alerts)
				}
				assertAlert(t, alerter.alerts[i], c)
			})
		}
	})
}

func TestPoker_Finish(t *testing.T) {
	t.Run("Game finish", func(t *testing.T) {
		store := &stubPlayerStore{}
		poker := NewPoker(dummyBlindAlerter, store)
		poker.Finish("zahid")
		assertWinCalls(t, store.winCalls, []string{"zahid"})
	})
}

func assertAlert(t *testing.T, got, want alert) {
	t.Helper()
	if got.amount != want.amount {
		t.Errorf("got amount '%d' want '%d'", got.amount, want.amount)
	}
	if got.scheduledAt != want.scheduledAt {
		t.Errorf("got schedule time '%v' want '%v'", got.scheduledAt, want.scheduledAt)
	}
}

func assertWinCalls(t *testing.T, got, want []string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}

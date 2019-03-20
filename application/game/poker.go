package game

import (
	"time"

	"github.com/DiscoFighter47/goingTDD/application/alerter"
	"github.com/DiscoFighter47/goingTDD/application/data"
)

// Poker struct
type Poker struct {
	alerter alerter.BlindAlerter
	store   data.PlayerStore
}

// NewPoker returns a new game
func NewPoker(alerter alerter.BlindAlerter, store data.PlayerStore) *Poker {
	return &Poker{alerter, store}
}

// Start starts the game
func (p *Poker) Start(numPlayer int) {
	blinds := []int{100, 200, 400, 600, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	blindInc := time.Duration(5+numPlayer) * time.Second
	for _, blind := range blinds {
		p.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime += blindInc
	}
}

// Finish finishes a game
func (p *Poker) Finish(player string) {
	p.store.RegisterWin(player)
}

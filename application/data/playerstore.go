package data

import "github.com/DiscoFighter47/goingTDD/application/model"

// PlayerStore stores players information
type PlayerStore interface {
	GetPlayerScore(string) (int, bool)
	RegisterWin(string)
	GetLeagueTable() []*model.Player
}

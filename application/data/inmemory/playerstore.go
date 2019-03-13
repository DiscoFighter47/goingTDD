package inmemory

import "github.com/DiscoFighter47/goingTDD/application/model"

// PlayerStore for fixed storage
type PlayerStore struct {
	store map[string]int
}

// NewPlayerStore returns a new In memory store
func NewPlayerStore() *PlayerStore {
	return &PlayerStore{map[string]int{}}
}

// GetPlayerScore returns fixed score
func (s *PlayerStore) GetPlayerScore(name string) (int, bool) {
	val, found := s.store[name]
	return val, found
}

// RegisterWin regiters score
func (s *PlayerStore) RegisterWin(name string) {
	s.store[name]++
}

// GetLeagueTable returns league table
func (s *PlayerStore) GetLeagueTable() []*model.Player {
	league := []*model.Player{}
	for name, score := range s.store {
		league = append(league, &model.Player{
			Name:  name,
			Score: score,
		})
	}
	return league
}

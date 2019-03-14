package file

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/DiscoFighter47/goingTDD/application/model"
)

type league []model.Player

func (l league) find(name string) *model.Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}

// PlayerStore storese palyer's information
type PlayerStore struct {
	database *json.Encoder
	league   league
}

func initFile(file *os.File) error {
	file.Seek(0, 0)
	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("unable to read file %s %v", file.Name(), err)
	}
	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}
	return nil
}

// NewPlayerStore returns a new player store
func NewPlayerStore(file *os.File) (*PlayerStore, error) {
	if err := initFile(file); err != nil {
		return nil, fmt.Errorf("unable to init file %s %v", file.Name(), err)
	}
	l := league{}
	err := json.NewDecoder(file).Decode(&l)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s %v", file.Name(), err)
	}
	return &PlayerStore{
		database: json.NewEncoder(&tape{file}),
		league:   l,
	}, nil
}

// GetPlayerScore returns fixed score
func (s *PlayerStore) GetPlayerScore(name string) (int, bool) {
	player := s.league.find(name)
	if player == nil {
		return 0, false
	}
	return player.Score, true
}

// RegisterWin registers score
func (s *PlayerStore) RegisterWin(name string) {
	player := s.league.find(name)
	if player != nil {
		player.Score++
	} else {
		s.league = append(s.league, model.Player{
			Name:  name,
			Score: 1,
		})
	}
	s.database.Encode(&s.league)
}

// GetLeagueTable returns league table
func (s *PlayerStore) GetLeagueTable() []model.Player {
	sort.Slice(s.league, func(i, j int) bool {
		return s.league[i].Score > s.league[j].Score
	})
	return s.league
}

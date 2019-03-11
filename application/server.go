package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Player represents a player's score information
type Player struct {
	Name  string
	Score int
}

// PlayerStore stores players information
type PlayerStore interface {
	GetPlayerScore(string) (int, bool)
	RegisterWin(string)
	GetLeagueTable() []*Player
}

// PlayerServer serves players information
type PlayerServer struct {
	store PlayerStore
	http.Handler
}

// NewPlayerServer returns a new player server
func NewPlayerServer(store PlayerStore) *PlayerServer {
	svr := &PlayerServer{
		store: store,
	}

	router := http.NewServeMux()
	router.HandleFunc("/league", svr.leagueHandler)
	router.HandleFunc("/players/", svr.scoreHandler)
	svr.Handler = router

	return svr
}

func (svr *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(svr.store.GetLeagueTable())
	w.WriteHeader(http.StatusOK)
}

func (svr *PlayerServer) scoreHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		svr.serveScore(w, r)
	case http.MethodPost:
		svr.registerScore(w, r)
	}
}

func (svr *PlayerServer) serveScore(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	score, found := svr.store.GetPlayerScore(player)
	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Fprint(w, score)
}

func (svr *PlayerServer) registerScore(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	svr.store.RegisterWin(player)
	w.WriteHeader(http.StatusAccepted)
}

// InMemoryStore for fixed storage
type InMemoryStore struct {
	store map[string]int
}

// NewInMemoryStore returns a new In memory store
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{map[string]int{}}
}

// GetPlayerScore returns fixed score
func (s *InMemoryStore) GetPlayerScore(name string) (int, bool) {
	val, found := s.store[name]
	return val, found
}

// RegisterWin regiters score
func (s *InMemoryStore) RegisterWin(name string) {
	s.store[name]++
}

// GetLeagueTable returns league table
func (s *InMemoryStore) GetLeagueTable() []*Player {
	league := []*Player{}
	for name, score := range s.store {
		league = append(league, &Player{name, score})
	}
	return league
}

func main() {
	svr := NewPlayerServer(NewInMemoryStore())
	log.Println("Server starting on port: 8080")
	if err := http.ListenAndServe(":8080", svr); err != nil {
		log.Fatal("unable to serve", err)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

// PlayerStore stores players information
type PlayerStore interface {
	GetPlayerScore(string) (int, bool)
	RegisterWin(string)
}

// PlayerServer serves players information
type PlayerServer struct {
	Store PlayerStore
}

func (svr *PlayerServer) serveScore(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	score, found := svr.Store.GetPlayerScore(player)
	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Fprint(w, score)
}

func (svr *PlayerServer) registerScore(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	svr.Store.RegisterWin(player)
	w.WriteHeader(http.StatusAccepted)
}

// ServeHTTP serves players score
func (svr *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		svr.serveScore(w, r)
	case http.MethodPost:
		svr.registerScore(w, r)
	}
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

func main() {
	svr := &PlayerServer{NewInMemoryStore()}
	log.Println("Server starting on port: 8080")
	if err := http.ListenAndServe(":8080", svr); err != nil {
		log.Fatal("unable to serve", err)
	}
}

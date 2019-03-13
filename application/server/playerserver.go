package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DiscoFighter47/goingTDD/application/data"
)

// PlayerServer serves players information
type PlayerServer struct {
	store data.PlayerStore
	http.Handler
}

// NewPlayerServer returns a new player server
func NewPlayerServer(store data.PlayerStore) *PlayerServer {
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

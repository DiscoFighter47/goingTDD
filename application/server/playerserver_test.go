package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/DiscoFighter47/goingTDD/application/data/memory"
	"github.com/DiscoFighter47/goingTDD/application/model"
)

type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	league   []*model.Player
}

func (stub *StubPlayerStore) GetPlayerScore(name string) (int, bool) {
	val, found := stub.Scores[name]
	return val, found
}

func (stub *StubPlayerStore) RegisterWin(name string) {
	stub.WinCalls = append(stub.WinCalls, name)
}

func (stub *StubPlayerStore) GetLeagueTable() []*model.Player {
	return stub.league
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}
func assertWinCalls(t *testing.T, got, want []string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}
func assertLeague(t *testing.T, got, want []*model.Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}

func TestServeScore(t *testing.T) {
	stub := &StubPlayerStore{
		Scores: map[string]int{
			"pepper": 20,
			"floyd":  10,
		},
	}
	svr := NewPlayerServer(stub)

	t.Run("Return Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/pepper", nil)
		response := httptest.NewRecorder()
		svr.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")

	})

	t.Run("Return Floyd's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/floyd", nil)
		response := httptest.NewRecorder()
		svr.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("Return 404 for missing player", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/zahid", nil)
		response := httptest.NewRecorder()
		svr.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusNotFound)
		assertResponseBody(t, response.Body.String(), "")
	})
}

func TestRegisterScore(t *testing.T) {
	stub := &StubPlayerStore{
		Scores: map[string]int{
			"pepper": 20,
			"floyd":  10,
		},
	}
	svr := NewPlayerServer(stub)
	t.Run("Post Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/pepper", nil)
		response := httptest.NewRecorder()
		svr.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusAccepted)
		assertWinCalls(t, stub.WinCalls, []string{"pepper"})
	})
}

func TestLeague(t *testing.T) {
	league := []*model.Player{
		{
			Name:  "Zahid",
			Score: 20,
		},
		{
			Name:  "Al",
			Score: 15,
		},
		{
			Name:  "Tair",
			Score: 10,
		},
	}
	stub := &StubPlayerStore{nil, nil, league}
	svr := NewPlayerServer(stub)

	t.Run("Retursn 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/league", nil)
		response := httptest.NewRecorder()
		svr.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)

		var got []*model.Player
		if err := json.NewDecoder(response.Body).Decode(&got); err != nil {
			t.Fatal("unable to parse response body")
		}
		assertLeague(t, got, league)
	})
}

func TestIntegration(t *testing.T) {
	svr := NewPlayerServer(memory.NewPlayerStore())
	t.Run("Post Pepper's score and get score", func(t *testing.T) {
		requestPost, _ := http.NewRequest(http.MethodPost, "/players/pepper", nil)
		requestGet, _ := http.NewRequest(http.MethodGet, "/players/pepper", nil)

		svr.ServeHTTP(httptest.NewRecorder(), requestPost)
		svr.ServeHTTP(httptest.NewRecorder(), requestPost)
		svr.ServeHTTP(httptest.NewRecorder(), requestPost)

		response1 := httptest.NewRecorder()
		svr.ServeHTTP(response1, requestGet)
		assertResponseBody(t, response1.Body.String(), "3")

		svr.ServeHTTP(httptest.NewRecorder(), requestPost)
		svr.ServeHTTP(httptest.NewRecorder(), requestPost)

		response2 := httptest.NewRecorder()
		svr.ServeHTTP(response2, requestGet)
		assertResponseBody(t, response2.Body.String(), "5")

		requestLeague, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response3 := httptest.NewRecorder()
		svr.ServeHTTP(response3, requestLeague)
		var got []*model.Player
		if err := json.NewDecoder(response3.Body).Decode(&got); err != nil {
			t.Fatal("unable to parse response body")
		}
		assertLeague(t, got, []*model.Player{
			{
				Name:  "pepper",
				Score: 5,
			},
		})
	})
}

package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
}

func (stub *StubPlayerStore) GetPlayerScore(name string) (int, bool) {
	val, found := stub.Scores[name]
	return val, found
}

func (stub *StubPlayerStore) RegisterWin(name string) {
	stub.WinCalls = append(stub.WinCalls, name)
}

func TestGetPlayers(t *testing.T) {
	stub := &StubPlayerStore{
		Scores: map[string]int{
			"pepper": 20,
			"floyd":  10,
		},
	}
	svr := &PlayerServer{stub}

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
	svr := &PlayerServer{stub}
	t.Run("Post Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/pepper", nil)
		response := httptest.NewRecorder()
		svr.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusAccepted)
		assertWinCalls(t, stub.WinCalls, []string{"pepper"})
	})
}

func TestIntegration(t *testing.T) {
	svr := PlayerServer{NewInMemoryStore()}
	t.Run("Post Pepper's score and get score", func(t *testing.T) {
		requestPost, _ := http.NewRequest(http.MethodPost, "/players/pepper", nil)
		requestGet, _ := http.NewRequest(http.MethodGet, "/players/pepper", nil)
		response1 := httptest.NewRecorder()
		response2 := httptest.NewRecorder()

		svr.ServeHTTP(httptest.NewRecorder(), requestPost)
		svr.ServeHTTP(httptest.NewRecorder(), requestPost)
		svr.ServeHTTP(httptest.NewRecorder(), requestPost)
		svr.ServeHTTP(response1, requestGet)
		assertResponseBody(t, response1.Body.String(), "3")

		svr.ServeHTTP(httptest.NewRecorder(), requestPost)
		svr.ServeHTTP(httptest.NewRecorder(), requestPost)
		svr.ServeHTTP(response2, requestGet)
		assertResponseBody(t, response2.Body.String(), "5")
	})
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

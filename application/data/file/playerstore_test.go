package file

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/DiscoFighter47/goingTDD/application/model"
)

func assertScore(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}

func assertLeague(t *testing.T, got, want []model.Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}

func assertNoErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("unexpected error '%v'", err)
	}
}

func createTempFile(t *testing.T, data string) (*os.File, func()) {
	t.Helper()
	tempFile, err := ioutil.TempFile("", "db")
	if err != nil {
		t.Fatal("could not crete a temp file", err)
	}
	tempFile.WriteString(data)
	removeFile := func() {
		os.Remove(tempFile.Name())
	}
	return tempFile, removeFile
}

func TestPlayerStore(t *testing.T) {
	t.Run("League from a reader", func(t *testing.T) {
		db, cleanDB := createTempFile(t, `[
			{"name": "Zahid", "score": 20},
			{"name": "Al", "score": 15},
			{"name": "Tair", "score": 10}
		]`)
		defer cleanDB()
		store, _ := NewPlayerStore(db)
		want := []model.Player{
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

		got := store.GetLeagueTable()
		assertLeague(t, got, want)

		got = store.GetLeagueTable()
		assertLeague(t, got, want)
	})

	t.Run("Get player score", func(t *testing.T) {
		db, cleanDB := createTempFile(t, `[
			{"name": "Zahid", "score": 20},
			{"name": "Al", "score": 15},
			{"name": "Tair", "score": 10}
		]`)
		defer cleanDB()
		store, _ := NewPlayerStore(db)

		got, _ := store.GetPlayerScore("Zahid")
		want := 20
		assertScore(t, got, want)

		got, _ = store.GetPlayerScore("Tair")
		want = 10
		assertScore(t, got, want)
	})

	t.Run("Store win for exiting player", func(t *testing.T) {
		db, cleanDB := createTempFile(t, `[
			{"name": "Zahid", "score": 20},
			{"name": "Al", "score": 15},
			{"name": "Tair", "score": 10}
		]`)
		defer cleanDB()
		store, _ := NewPlayerStore(db)

		store.RegisterWin("Zahid")
		got, _ := store.GetPlayerScore("Zahid")
		want := 21
		assertScore(t, got, want)
	})

	t.Run("Store win for new player", func(t *testing.T) {
		db, cleanDB := createTempFile(t, `[
			{"name": "Zahid", "score": 20},
			{"name": "Al", "score": 15}
		]`)
		defer cleanDB()
		store, _ := NewPlayerStore(db)

		store.RegisterWin("Tair")
		got, _ := store.GetPlayerScore("Tair")
		want := 1
		assertScore(t, got, want)
	})

	t.Run("Working with empty file", func(t *testing.T) {
		db, cleanDB := createTempFile(t, "")
		defer cleanDB()
		_, err := NewPlayerStore(db)
		assertNoErr(t, err)
	})

	t.Run("Sorted league", func(t *testing.T) {
		db, cleanDB := createTempFile(t, `[
			{"name": "Zahid", "score": 15},
			{"name": "Al", "score": 10},
			{"name": "Tair", "score": 20}
		]`)
		defer cleanDB()
		store, _ := NewPlayerStore(db)
		want := []model.Player{
			{
				Name:  "Tair",
				Score: 20,
			},
			{
				Name:  "Zahid",
				Score: 15,
			},
			{
				Name:  "Al",
				Score: 10,
			},
		}
		got := store.GetLeagueTable()
		assertLeague(t, got, want)

		got = store.GetLeagueTable()
		assertLeague(t, got, want)
	})
}

package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDealyedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {
	slowServer := makeDealyedServer(2 * time.Second)
	defer slowServer.Close()

	fastServer := makeDealyedServer(0 * time.Second)
	defer fastServer.Close()

	got, err := Racer(slowServer.URL, fastServer.URL)
	want := fastServer.URL

	if err != nil {
		t.Errorf("unexpected error '%s'", err)
	}

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

	t.Run("timeout", func(t *testing.T) {
		slowServer := makeDealyedServer(3 * time.Second)
		defer slowServer.Close()

		fastServer := makeDealyedServer(2 * time.Second)
		defer fastServer.Close()

		_, err := configurableRacer(slowServer.URL, fastServer.URL, 1*time.Second)
		if err == nil {
			t.Errorf("expected error but didn't get any")
		}
	})
}

package racer

import (
	"errors"
	"net/http"
	"time"
)

var defaultTimeout = 10 * time.Second

// Racer returns faster url among two
func Racer(urlA, urlB string) (string, error) {
	return configurableRacer(urlA, urlB, defaultTimeout)
}

func configurableRacer(urlA, urlB string, timeout time.Duration) (string, error) {
	select {
	case <-ping(urlA):
		return urlA, nil
	case <-ping(urlB):
		return urlB, nil
	case <-time.After(timeout):
		return "", errors.New("timeout")
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}

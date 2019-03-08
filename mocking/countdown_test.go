package main

import (
	"bytes"
	"reflect"
	"testing"
)

type mockSleeper struct {
	calls int
}

func (s *mockSleeper) Sleep() {
	s.calls++
}

const write = "write"
const sleep = "sleep"

type countdownSpy struct {
	calls []string
}

func (spy *countdownSpy) Sleep() {
	spy.calls = append(spy.calls, sleep)
}

func (spy *countdownSpy) Write(p []byte) (n int, err error) {
	spy.calls = append(spy.calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("Print 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &mockSleeper{}
		Countdown(buffer, sleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}

		if sleeper.calls != 4 {
			t.Errorf("not enough calls to sleeprs, want '4', got '%d'", sleeper.calls)
		}
	})

	t.Run("Sleep after every print", func(t *testing.T) {
		spy := &countdownSpy{}
		Countdown(spy, spy)
		want := []string{sleep, write, sleep, write, sleep, write, sleep, write}
		got := spy.calls
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got calls '%v' want '%v'", got, want)
		}
	})
}

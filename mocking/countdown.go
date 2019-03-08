package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Sleeper sleeps for a moment
type Sleeper interface {
	Sleep()
}

// DefaultSleeper sleeps for 1 second
type DefaultSleeper struct{}

// Sleep sleeps for 1 second
func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const count = 3
const finalWord = "Go!"

// Countdown countdowns from 3 to 1
func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := count; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(writer, i)
	}
	sleeper.Sleep()
	fmt.Fprint(writer, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

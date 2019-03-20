package alerter

import (
	"fmt"
	"os"
	"time"
)

// BlindAlerter interface
type BlindAlerter interface {
	ScheduleAlertAt(time.Duration, int)
}

// BlindAlerterFunc implementation
type BlindAlerterFunc func(time.Duration, int)

// ScheduleAlertAt implementation
func (a BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int) {
	a(duration, amount)
}

// StdOutAlerter implementation
func StdOutAlerter(duration time.Duration, amount int) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(os.Stdout, "blind is now %d\n", amount)
	})
}

package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	assertCounter := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	}

	t.Run("Count to 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()
		assertCounter(t, counter.Value, 3)
	})

	t.Run("concurrent counter", func(t *testing.T) {
		want := 10000
		counter := NewCounter()
		var wg sync.WaitGroup
		wg.Add(want)
		for i := 0; i < want; i++ {
			go func(wg *sync.WaitGroup) {
				counter.Inc()
				wg.Done()
			}(&wg)
		}
		wg.Wait()
		assertCounter(t, counter.Value, want)
	})
}

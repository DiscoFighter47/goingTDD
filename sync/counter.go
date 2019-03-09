package counter

import "sync"

// Counter counts
type Counter struct {
	mu    sync.Mutex
	Value int
}

// NewCounter returns a counter
func NewCounter() *Counter {
	return &Counter{}
}

// Inc increments counter value
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Value++
}

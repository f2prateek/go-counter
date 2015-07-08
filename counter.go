package counter

import "sync/atomic"

type Counter struct {
	values map[string]*int64
}

// Create a new Counter
func NewCounter() *Counter {
	m := make(map[string]*int64)
	return &Counter{m}
}

// Increment the value for the given key
func (c *Counter) Increment(key string) {
	if _, ok := c.values[key]; !ok {
		var n int64 = 0
		c.values[key] = &n
	}

	atomic.AddInt64(c.values[key], 1)
}

// Return a copy of the counter values
func (c *Counter) Values() map[string]int64 {
	v := make(map[string]int64)
	for key, value := range c.values {
		v[key] = *value
	}
	return v
}

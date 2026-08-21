// main package helpers shared by the concurrency demos
package main

import "sync"

// sumOf drains a channel and returns the sum of its values
func sumOf(ch <-chan int) int {
	total := 0
	for v := range ch {
		total += v
	}
	return total
}

// counter is a mutex-protected integer used in the race-free demo
type counter struct {
	mu    sync.Mutex
	value int
}

func (c *counter) add(delta int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value += delta
}

func (c *counter) get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

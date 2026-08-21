// main package demonstrates fan-in and fan-out channel patterns
package main

import (
	"fmt"
	"sync"
)

// FanIn merges multiple channels into one; output closes when all inputs close
func FanIn(inputs ...<-chan int) <-chan int {
	merged := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(inputs))

	for _, in := range inputs {
		go func(ch <-chan int) {
			defer wg.Done()
			for v := range ch {
				merged <- v
			}
		}(in)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()
	return merged
}

// Distribute sends every item to exactly one of the worker channels (fan-out)
func Distribute(items []int, workers int) []<-chan int {
	chans := make([]chan int, workers)
	outs := make([]<-chan int, workers)
	for i := range chans {
		chans[i] = make(chan int)
		outs[i] = chans[i]
	}

	go func() {
		for idx, item := range items {
			chans[idx%workers] <- item
		}
		for _, ch := range chans {
			close(ch)
		}
	}()
	return outs
}

// Collect drains a channel into a slice
func Collect(ch <-chan int) []int {
	var out []int
	for v := range ch {
		out = append(out, v)
	}
	return out
}

// FanInMain demonstrates merging channels
func FanInMain() {
	a := make(chan int, 2)
	b := make(chan int, 2)
	a <- 1
	a <- 3
	b <- 2
	b <- 4
	close(a)
	close(b)

	fmt.Println("merged sum:", sumOf(FanIn(a, b)))
}

// FanOutMain demonstrates distributing work across channels
func FanOutMain() {
	outs := Distribute([]int{1, 2, 3, 4}, 2)
	c := &counter{}
	var wg sync.WaitGroup
	wg.Add(len(outs))
	for _, ch := range outs {
		go func(cch <-chan int) {
			defer wg.Done()
			for v := range cch {
				c.add(v) // race-free accumulation
			}
		}(ch)
	}
	wg.Wait()
	fmt.Println("fan-out total:", c.get())
}

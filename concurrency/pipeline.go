// main package demonstrates pipelines, semaphores and graceful shutdown
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// PipelineSum builds generator -> doubler -> sum and returns the result
func PipelineSum(values []int) int {
	gen := make(chan int)
	doubled := make(chan int)

	go func() {
		defer close(gen)
		for _, v := range values {
			gen <- v
		}
	}()

	go func() {
		defer close(doubled)
		for v := range gen {
			doubled <- v * 2
		}
	}()

	total := 0
	for v := range doubled {
		total += v
	}
	return total
}

// RunWithLimit runs every task but never more than limit concurrently
func RunWithLimit(tasks []func(), limit int) {
	sem := make(chan struct{}, limit)
	var wg sync.WaitGroup
	for _, task := range tasks {
		wg.Add(1)
		sem <- struct{}{} // acquire slot; blocks when full
		go func(task func()) {
			defer wg.Done()
			defer func() { <-sem }()
			task()
		}(task)
	}
	wg.Wait()
}

// ServeUntilCancelled simulates a server loop that shuts down gracefully:
// it stops accepting work on ctx.Done and waits for in-flight work to finish
func ServeUntilCancelled(ctx context.Context, requests <-chan int) int {
	var wg sync.WaitGroup
	handled := 0
	var mu sync.Mutex

	accepting := true
	for accepting {
		select {
		case req, ok := <-requests:
			if !ok {
				accepting = false
				break
			}
			wg.Add(1)
			go func(r int) {
				defer wg.Done()
				time.Sleep(time.Millisecond) // simulated handling
				mu.Lock()
				handled += r
				mu.Unlock()
			}(req)
		case <-ctx.Done():
			accepting = false
		}
	}

	wg.Wait() // graceful: wait for in-flight handlers
	return handled
}

// PipelineMain demonstrates the three-stage pipeline
func PipelineMain() {
	fmt.Println("pipeline sum of 1..4 doubled:", PipelineSum([]int{1, 2, 3, 4}))
}

// SemaphoreMain demonstrates bounded concurrency
func SemaphoreMain() {
	c := &counter{}
	RunWithLimit([]func(){
		func() { c.add(1) },
		func() { c.add(2) },
		func() { c.add(3) },
	}, 2)
	fmt.Println("semaphore total:", c.get())
}

// ShutdownMain demonstrates graceful shutdown
func ShutdownMain() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	requests := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			requests <- i
			time.Sleep(5 * time.Millisecond)
		}
		close(requests)
	}()

	fmt.Println("handled total:", ServeUntilCancelled(ctx, requests))
}

// main runs the demo entry points of this package
func main() {
	FanInMain()
	FanOutMain()
	PipelineMain()
	SemaphoreMain()
	ShutdownMain()
	WorkerPoolMain()
}

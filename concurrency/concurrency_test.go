package main

import (
	"context"
	"reflect"
	"sort"
	"sync"
	"testing"
	"time"
)

func TestWorkerPoolPreservesOrder(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got := RunWorkerPool(input, 4)
	want := []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("RunWorkerPool = %v, want %v", got, want)
	}
}

func TestWorkerPoolEmpty(t *testing.T) {
	if got := RunWorkerPool([]int{}, 3); len(got) != 0 {
		t.Errorf("expected empty result, got %v", got)
	}
}

func TestFanInMergesAndCloses(t *testing.T) {
	a := make(chan int, 2)
	b := make(chan int, 2)
	a <- 1
	a <- 2
	close(a)
	b <- 3
	close(b)

	got := Collect(FanIn(a, b))
	sort.Ints(got)
	if !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("FanIn collected %v", got)
	}
}

func TestFanInWithNoInputs(t *testing.T) {
	got := Collect(FanIn())
	if len(got) != 0 {
		t.Errorf("expected empty, got %v", got)
	}
}

func TestDistributeSendsEverythingExactlyOnce(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	outs := Distribute(items, 3)

	var mu sync.Mutex
	var got []int
	var wg sync.WaitGroup
	wg.Add(len(outs))
	for _, ch := range outs {
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				mu.Lock()
				got = append(got, v)
				mu.Unlock()
			}
		}(ch)
	}
	wg.Wait()

	sort.Ints(got)
	if !reflect.DeepEqual(got, items) {
		t.Errorf("distributed = %v, want %v", got, items)
	}
}

func TestPipelineSum(t *testing.T) {
	if got := PipelineSum([]int{1, 2, 3, 4}); got != 20 {
		t.Errorf("PipelineSum = %d, want 20", got)
	}
	if got := PipelineSum(nil); got != 0 {
		t.Errorf("PipelineSum(nil) = %d, want 0", got)
	}
}

func TestRunWithLimitRunsAllTasks(t *testing.T) {
	c := &counter{}
	RunWithLimit([]func(){
		func() { c.add(1) },
		func() { c.add(1) },
		func() { c.add(1) },
	}, 2)
	if c.get() != 3 {
		t.Errorf("counter = %d, want 3", c.get())
	}
}

func TestServeUntilCancelledGracefulShutdown(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancel()

	requests := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			requests <- i
			time.Sleep(5 * time.Millisecond)
		}
		close(requests)
	}()

	handled := ServeUntilCancelled(ctx, requests)
	if handled == 0 {
		t.Error("server handled nothing before shutdown")
	}
}

func TestConcurrencyMainSmoke(t *testing.T) {
	FanInMain()
	FanOutMain()
	PipelineMain()
	SemaphoreMain()
	ShutdownMain()
	WorkerPoolMain()
}

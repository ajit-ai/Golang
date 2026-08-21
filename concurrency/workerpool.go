// main package demonstrates a fixed-size worker pool
package main

import (
	"fmt"
	"sort"
	"sync"
)

// job carries an index so results can be reassembled in order
type job struct {
	index int
	value int
}

// RunWorkerPool processes values with n workers, returning results
// in the original input order regardless of completion order
func RunWorkerPool(values []int, workers int) []int {
	jobs := make(chan job)
	results := make([]int, len(values))

	var wg sync.WaitGroup
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func() {
			defer wg.Done()
			for j := range jobs {
				results[j.index] = j.value * j.value // simulated work
			}
		}()
	}

	for i, v := range values {
		jobs <- job{index: i, value: v}
	}
	close(jobs)
	wg.Wait()
	return results
}

// WorkerPoolMain demonstrates the pool with 3 workers
func WorkerPoolMain() {
	got := RunWorkerPool([]int{1, 2, 3, 4, 5}, 3)
	fmt.Println("squares:", got)
}

var _ = sort.Ints // placeholder to keep sort available for future demos

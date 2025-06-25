package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"sync"
	"sync/atomic"
	"time"
)

// 1. Goroutines and Channels with Select
func goroutinesAndChannels() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "from ch1"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "from ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
		}
	}
}

// 2. Context for Cancellation and Deadlines
func contextExample() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func() {
		<-ctx.Done()
		fmt.Println("Task cancelled or timed out")
	}()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Task completed")
	case <-ctx.Done():
		fmt.Println("Context cancelled:", ctx.Err())
	}
}

// 3. Sync.Pool for Object Reuse
func syncPoolExample() {
	var pool = sync.Pool{
		New: func() interface{} {
			return make([]byte, 1024)
		},
	}

	// Get from pool
	buf := pool.Get().([]byte)
	fmt.Printf("Got buffer of size: %d\n", len(buf))

	// Return to pool
	pool.Put(buf)
}

// 4. Custom HTTP Server with Middleware
func customHTTPServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, Advanced Go!")
	})

	// Middleware for logging
	logMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			log.Printf("Started %s %s", r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
			log.Printf("Completed in %v", time.Since(start))
		})
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: logMiddleware(mux),
	}

	fmt.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

// 5. Rate Limiting with Token Bucket
func rateLimitingExample() {
	const rate = 2 // 2 requests per second
	bucket := make(chan struct{}, rate)

	// Token bucket filler
	go func() {
		for range time.Tick(time.Second / rate) {
			select {
			case bucket <- struct{}{}:
			default:
			}
		}
	}()

	// Simulate requests
	for i := 0; i < 5; i++ {
		<-bucket
		fmt.Printf("Request %d processed at %v\n", i+1, time.Now())
	}
}

// 6. Reflection for Dynamic Type Handling
func reflectionExample(v interface{}) {
	typ := reflect.TypeOf(v)
	val := reflect.ValueOf(v)

	fmt.Printf("Type: %v\n", typ)
	fmt.Printf("Value: %v\n", val)

	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			fmt.Printf("Element %d: %v\n", i, val.Index(i))
		}
	}
}

// 7. Atomic Operations for Lock-Free Concurrency
func atomicExample() {
	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter)
}

func main() {
	fmt.Println("1. Goroutines and Channels")
	goroutinesAndChannels()
	fmt.Println("\n2. Context Example")
	contextExample()
	fmt.Println("\n3. Sync.Pool Example")
	syncPoolExample()
	fmt.Println("\n4. Rate Limiting Example")
	rateLimitingExample()
	fmt.Println("\n5. Reflection Example")
	slice := []int{1, 2, 3}
	reflectionExample(slice)
	fmt.Println("\n6. Atomic Operations Example")
	atomicExample()

	// Run HTTP server in a separate goroutine
	go customHTTPServer()
	time.Sleep(1 * time.Second) // Allow server to start
}

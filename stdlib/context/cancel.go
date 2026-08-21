// main package demonstrates context: cancellation, timeouts and
// cancellation propagation through call chains
package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// ErrCancelled is returned when work is interrupted
var ErrCancelled = errors.New("work cancelled")

// DoWork simulates work that respects cancellation.
// It completes when done fires before ctx is cancelled.
func DoWork(ctx context.Context, workTime time.Duration) error {
	select {
	case <-time.After(workTime):
		return nil
	case <-ctx.Done():
		return ErrCancelled
	}
}

// RunWithTimeout runs DoWork under a timeout; returns the resulting error
func RunWithTimeout(workTime, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return DoWork(ctx, workTime)
}

// Child inherits its parent's cancellation: cancelling the parent
// interrupts the child even though the child has its own longer deadline
func Child(parent context.Context, workTime time.Duration) error {
	ctx, cancel := context.WithTimeout(parent, 10*time.Second)
	defer cancel()
	return DoWork(ctx, workTime)
}

// Values demonstrates carrying request-scoped values through a context
func Values(ctx context.Context) string {
	user, ok := ctx.Value(userKey{}).(string)
	if !ok {
		return "anonymous"
	}
	return user
}

type userKey struct{}

// ContextMain demonstrates timeouts and propagation
func ContextMain() {
	fmt.Println("fast work:", RunWithTimeout(10*time.Millisecond, time.Second))
	fmt.Println("slow work:", RunWithTimeout(time.Second, 10*time.Millisecond))

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(10 * time.Millisecond)
		cancel() // parent cancels the child chain
	}()
	fmt.Println("child after parent cancel:", Child(ctx, time.Second))

	authed := context.WithValue(context.Background(), userKey{}, "ajit")
	fmt.Println("user from context:", Values(authed))
}

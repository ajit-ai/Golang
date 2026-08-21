package main

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestDoWorkCompletesBeforeDeadline(t *testing.T) {
	if err := RunWithTimeout(5*time.Millisecond, time.Second); err != nil {
		t.Errorf("expected success, got %v", err)
	}
}

func TestDoWorkTimesOut(t *testing.T) {
	err := RunWithTimeout(time.Second, 5*time.Millisecond)
	if err == nil {
		t.Fatal("expected timeout error")
	}
	if !errors.Is(err, ErrCancelled) {
		t.Errorf("error = %v, want cancellation", err)
	}
}

func TestParentCancelPropagatesToChild(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(5 * time.Millisecond)
		cancel()
	}()
	if err := Child(ctx, time.Second); err == nil {
		t.Error("child should be cancelled by parent")
	}
}

func TestChildSurvivesWhenParentLives(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := Child(ctx, 5*time.Millisecond); err != nil {
		t.Errorf("child should finish: %v", err)
	}
}

func TestContextValues(t *testing.T) {
	ctx := context.WithValue(context.Background(), userKey{}, "ajit")
	if got := Values(ctx); got != "ajit" {
		t.Errorf("Values = %q", got)
	}
	if got := Values(context.Background()); got != "anonymous" {
		t.Errorf("missing value = %q, want anonymous", got)
	}
}

func TestContextMainSmoke(t *testing.T) {
	ContextMain()
}

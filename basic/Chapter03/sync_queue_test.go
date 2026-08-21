package main

import (
	"sync"
	"testing"
	"time"
)

func TestSyncQueueHandshake(t *testing.T) {
	q := &SyncQueue{}
	q.New()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(20 * time.Millisecond)
		q.StartPass()
		q.EndPass()
	}()

	q.StartTicketIssue()
	q.EndTicketIssue()

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("sync queue handshake deadlocked")
	}
}

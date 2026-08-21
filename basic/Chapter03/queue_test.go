package main

import "testing"

func TestQueueAdd(t *testing.T) {
	var queue Queue
	order := &Order{}
	order.New(2, 10, "widget", "alice")

	queue.Add(order)

	if len(queue) != 1 || queue[0] != order {
		t.Errorf("queue length/content wrong after Add: %v", queue)
	}
}

func TestQueueMainSmoke(t *testing.T) {
	QueueMain()
}

// SyncQueueMain is intentionally not smoke-tested: it ends with select{}
// and never returns (endless concurrency demo). See sync_queue_test.go
// for a bounded test of its handshake logic.

func TestTuples(t *testing.T) {
	if got := h(3, 4); got != 12 {
		t.Errorf("h(3, 4) = %d, want 12", got)
	}
	x, y := g(1, 2)
	if x != 2 || y != 8 {
		t.Errorf("g(1, 2) = %d, %d; want 2, 8", x, y)
	}
}

func TestTuplesMainSmoke(t *testing.T) {
	TuplesMain()
}

package main

import "testing"

func TestReferenceCounterAddSubtract(t *testing.T) {
	rc := newReferenceCounter()
	rc.Add()
	rc.Add()
	if got := *rc.num; got != 2 {
		t.Errorf("after two Adds, num = %d, want 2", got)
	}
	rc.Subtract()
	if got := *rc.num; got != 1 {
		t.Errorf("after Subtract, num = %d, want 1", got)
	}
	if got := *rc.removed; got != 0 {
		t.Errorf("removed = %d, want 0 while count above zero", got)
	}
}

func TestReferenceCounterSubtractToZeroCountsRemoved(t *testing.T) {
	rc := newReferenceCounter()
	rc.Add()
	rc.Subtract()
	if got := *rc.num; got != 0 {
		t.Errorf("num = %d, want 0", got)
	}
	if got := *rc.removed; got != 1 {
		t.Errorf("removed = %d, want 1 after 1->0 transition", got)
	}
}

func TestNewReferenceCounterInitializesFields(t *testing.T) {
	rc := newReferenceCounter()
	if rc.num == nil || rc.pool == nil || rc.removed == nil {
		t.Fatal("newReferenceCounter returned uninitialized fields")
	}
}

func TestReferenceCountingMainSmoke(t *testing.T) {
	ReferenceCountingMain()
}

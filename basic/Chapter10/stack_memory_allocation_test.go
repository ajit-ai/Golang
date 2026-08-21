package main

import "testing"

func TestAddOneValueSemantics(t *testing.T) {
	if got := addOne(17); got != 18 {
		t.Errorf("addOne(17) = %d, want 18", got)
	}
	if got := addOne(0); got != 1 {
		t.Errorf("addOne(0) = %d, want 1", got)
	}
	if got := addOne(-5); got != -4 {
		t.Errorf("addOne(-5) = %d, want -4", got)
	}
}

func TestStackMemoryAllocationMainSmoke(t *testing.T) {
	StackMemoryAllocationMain()
}

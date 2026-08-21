package main

import "testing"

func TestAddOnePointerMutatesCallerValue(t *testing.T) {
	number := 17
	addOnePointer(&number)
	if number != 18 {
		t.Errorf("after addOnePointer, number = %d, want 18", number)
	}
	number = 0
	addOnePointer(&number)
	if number != 1 {
		t.Errorf("after second addOnePointer, number = %d, want 1", number)
	}
}

func TestStackMemoryPointerMainSmoke(t *testing.T) {
	StackMemoryPointerMain()
}

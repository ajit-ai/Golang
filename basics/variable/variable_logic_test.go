package main

import "testing"

func TestAddReturnsSumAndDifference(t *testing.T) {
	sum, diff := add(7, 2)
	if sum != 9 || diff != 5 {
		t.Errorf("add(7, 2) = %d, %d; want 9, 5", sum, diff)
	}
}

func TestMultiReturn(t *testing.T) {
	n, s := multiReturn()
	if n != 1 || s != "two" {
		t.Errorf("multiReturn() = %d, %q; want 1, two", n, s)
	}
}

func TestGlobalVar(t *testing.T) {
	if globalVar != "I am global" {
		t.Errorf("globalVar = %q", globalVar)
	}
}

package main

import "testing"

func TestOffsetArray(t *testing.T) {
	got := OffsetArray()
	if len(got) != 10 {
		t.Fatalf("len = %d, want 10", len(got))
	}
	for k, v := range got {
		if v != k+200 {
			t.Errorf("got[%d] = %d, want %d", k, v, k+200)
		}
	}
}

func TestComplexityMain(t *testing.T) {
	ComplexityMain()
}

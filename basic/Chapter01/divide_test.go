package main

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		k    int
		want int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{6, 13},
		{7, 21},
		{10, 89},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("k%d", tt.k), func(t *testing.T) {
			if got := Fibonacci(tt.k); got != tt.want {
				t.Errorf("Fibonacci(%d) = %d, want %d", tt.k, got, tt.want)
			}
		})
	}
}

func TestDivideMain(t *testing.T) {
	DivideMain()
}

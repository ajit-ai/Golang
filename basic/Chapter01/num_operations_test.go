package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"demo values", 42, 13, 55},
		{"zeros", 0, 0, 0},
		{"negatives", -3, -4, -7},
		{"mixed signs", -5, 5, 0},
		{"positive and negative", 1, -1, 0},
		{"large values", 1000000, 2000000, 3000000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.a, tt.b); got != tt.want {
				t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestNumOperationsMain(t *testing.T) {
	NumOperationsMain()
}

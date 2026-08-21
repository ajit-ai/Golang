package main

import "testing"

func TestFindElementsWithSum(t *testing.T) {
	tests := []struct {
		name string
		arr  [10]int
		size int
		k    int
		want int
	}{
		{"compositions of 3 from 1,2,3", [10]int{1, 2, 3}, 3, 3, 3},
		{"twos and fours to six", [10]int{2, 4}, 2, 6, 2},
		{"target zero counts empty combination", [10]int{1, 2, 3}, 3, 0, 1},
		{"unreachable target", [10]int{1, 2, 3}, 3, 100, 0},
		{"negative target", [10]int{1, 2, 3}, 3, -5, 0},
		{"single element match", [10]int{5}, 1, 5, 1},
		{"single element mismatch", [10]int{5}, 1, 3, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var combinations [19]int
			if got := FindElementsWithSum(tt.arr, combinations, tt.size, tt.k, 0, 0, 0); got != tt.want {
				t.Errorf("FindElementsWithSum(size=%d, k=%d) = %d, want %d", tt.size, tt.k, got, tt.want)
			}
		})
	}
}

func TestBacktrackingMain(t *testing.T) {
	BacktrackingMain()
}

package main

import "testing"

func TestFindElement(t *testing.T) {
	arr := [10]int{1, 4, 7, 8, 3, 9, 2, 4, 1, 8}
	tests := []struct {
		name string
		k    int
		want bool
	}{
		{"present middle", 9, true},
		{"absent", 10, false},
		{"first element", 1, true},
		{"last element", 8, true},
		{"duplicate present", 4, true},
		{"zero absent", 0, false},
		{"negative absent", -3, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindElement(arr, tt.k); got != tt.want {
				t.Errorf("FindElement(arr, %d) = %v, want %v", tt.k, got, tt.want)
			}
		})
	}
}

func TestBruteforceMain(t *testing.T) {
	BruteforceMain()
}

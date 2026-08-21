package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTableRow(t *testing.T) {
	tests := []struct {
		k    int
		want []int
	}{
		{1, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{10, []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}},
		{0, make([]int, 10)},
		{-2, []int{-2, -4, -6, -8, -10, -12, -14, -16, -18, -20}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("k%d", tt.k), func(t *testing.T) {
			got := TableRow(tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TableRow(%d) = %v, want %v", tt.k, got, tt.want)
			}
		})
	}
}

func TestQuadraticComplexityMain(t *testing.T) {
	QuadraticComplexityMain()
}

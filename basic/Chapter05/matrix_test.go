package main

import (
	"reflect"
	"testing"
)

func TestChangeMatrixSetsRowAndColumn(t *testing.T) {
	matrix := [3][3]int{
		{1, 0, 0},
		{0, 0, 0},
		{0, 1, 0},
	}
	got := changeMatrix(matrix)
	want := [3][3]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 1, 1},
	}
	if got != want {
		t.Errorf("changeMatrix = %v, want %v", got, want)
	}
}

func TestChangeMatrixAllZeros(t *testing.T) {
	var matrix [3][3]int
	if got := changeMatrix(matrix); got != matrix {
		t.Errorf("changeMatrix of zero matrix = %v, want unchanged", got)
	}
}

func TestIdentity(t *testing.T) {
	id := Identity(3)
	if len(id) != 3 {
		t.Fatalf("Identity(3) has %d rows, want 3", len(id))
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			want := 0.0
			if i == j {
				want = 1.0
			}
			if id[i][j] != want {
				t.Errorf("Identity[%d][%d] = %v, want %v", i, j, id[i][j], want)
			}
		}
	}
}

func TestPrintSpiralShapeAndCompleteness(t *testing.T) {
	got := PrintSpiral(4)
	if len(got) != 16 {
		t.Fatalf("PrintSpiral(4) length = %d, want 16", len(got))
	}
	seen := map[int]bool{}
	for _, v := range got {
		seen[v] = true
	}
	for i := 0; i < 16; i++ {
		if !seen[i] {
			t.Errorf("spiral missing sequence value %d", i)
		}
	}
}

func TestPrintSpiralKnownLayout(t *testing.T) {
	got := PrintSpiral(3)
	want := []int{0, 1, 2, 7, 8, 3, 6, 5, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("PrintSpiral(3) = %v, want %v", got, want)
	}
}

func TestPrintZigZagKnownLayout(t *testing.T) {
	got := PrintZigZag(3)
	want := []int{0, 1, 5, 2, 4, 6, 3, 7, 8}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("PrintZigZag(3) = %v, want %v", got, want)
	}
}

func TestMatrixOps(t *testing.T) {
	a := [2][2]int{{1, 2}, {3, 4}}
	b := [2][2]int{{5, 6}, {7, 8}}

	if got := add(a, b); got != ([2][2]int{{6, 8}, {10, 12}}) {
		t.Errorf("add = %v", got)
	}
	if got := subtract(b, a); got != ([2][2]int{{4, 4}, {4, 4}}) {
		t.Errorf("subtract = %v", got)
	}
	if got := multiply(a, b); got != ([2][2]int{{19, 22}, {43, 50}}) {
		t.Errorf("multiply = %v", got)
	}
	if got := transpose(a); got != ([2][2]int{{1, 3}, {2, 4}}) {
		t.Errorf("transpose = %v", got)
	}
	if got := determinant(a); got != -2 {
		t.Errorf("determinant = %v, want -2", got)
	}
}

func TestInverse(t *testing.T) {
	a := [2][2]int{{4, 7}, {2, 6}}
	inv := inverse(a)
	if inv == nil {
		t.Fatal("inverse returned nil for invertible matrix")
	}
	want := [][]float64{{0.6, -0.7}, {-0.2, 0.4}}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if inv[i][j] != want[i][j] {
				t.Errorf("inverse[%d][%d] = %v, want %v", i, j, inv[i][j], want[i][j])
			}
		}
	}
}

func TestInverseSingularReturnsNil(t *testing.T) {
	singular := [2][2]int{{1, 2}, {2, 4}}
	if inv := inverse(singular); inv != nil {
		t.Errorf("inverse of singular matrix = %v, want nil", inv)
	}
}

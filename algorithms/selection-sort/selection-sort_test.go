package main

import "testing"

// --- Testing code ---

func TestSelectionSortInt(t *testing.T) {
	arr := []int{5, 2, 9, 1, 5}
	SelectionSortInt(arr)
	expected := []int{1, 2, 5, 5, 9}
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("SelectionSortInt failed: got %v, want %v", arr, expected)
		}
	}
}

func TestSelectionSortFloat(t *testing.T) {
	arr := []float64{3.1, 2.2, 5.5, 1.0}
	SelectionSortFloat(arr)
	expected := []float64{1.0, 2.2, 3.1, 5.5}
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("SelectionSortFloat failed: got %v, want %v", arr, expected)
		}
	}
}

func TestSelectionSortString(t *testing.T) {
	arr := []string{"banana", "apple", "cherry"}
	SelectionSortString(arr)
	expected := []string{"apple", "banana", "cherry"}
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("SelectionSortString failed: got %v, want %v", arr, expected)
		}
	}
}

func TestSelectionSortIntDesc(t *testing.T) {
	arr := []int{3, 1, 4, 2}
	SelectionSortIntDesc(arr)
	expected := []int{4, 3, 2, 1}
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("SelectionSortIntDesc failed: got %v, want %v", arr, expected)
		}
	}
}

func TestSelectionSortPersonByAge(t *testing.T) {
	arr := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}
	SelectionSortPersonByAge(arr)
	expected := []Person{
		{"Bob", 25},
		{"Alice", 30},
		{"Charlie", 35},
	}
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("SelectionSortPersonByAge failed: got %v, want %v", arr, expected)
		}
	}
}

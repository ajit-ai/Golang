package main

import "testing"

// --- Testing code ---

func TestBubbleSortInt(t *testing.T) {
	arr := []int{5, 2, 9, 1, 5}
	BubbleSortInt(arr)
	expected := []int{1, 2, 5, 5, 9}
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("BubbleSortInt failed: got %v, want %v", arr, expected)
		}
	}
}

func TestBubbleSortFloat(t *testing.T) {
	arr := []float64{3.1, 2.2, 5.5, 1.0}
	BubbleSortFloat(arr)
	expected := []float64{1.0, 2.2, 3.1, 5.5}
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("BubbleSortFloat failed: got %v, want %v", arr, expected)
		}
	}
}

func TestBubbleSortString(t *testing.T) {
	arr := []string{"banana", "apple", "cherry"}
	BubbleSortString(arr)
	expected := []string{"apple", "banana", "cherry"}
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("BubbleSortString failed: got %v, want %v", arr, expected)
		}
	}
}

func TestBubbleSortIntDesc(t *testing.T) {
	arr := []int{3, 1, 4, 2}
	BubbleSortIntDesc(arr)
	expected := []int{4, 3, 2, 1}
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("BubbleSortIntDesc failed: got %v, want %v", arr, expected)
		}
	}
}

func TestBubbleSortPersonByAge(t *testing.T) {
	arr := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}
	BubbleSortPersonByAge(arr)
	expected := []Person{
		{"Bob", 25},
		{"Alice", 30},
		{"Charlie", 35},
	}
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("BubbleSortPersonByAge failed: got %v, want %v", arr, expected)
		}
	}
}

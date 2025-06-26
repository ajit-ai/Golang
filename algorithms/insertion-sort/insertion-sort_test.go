package main

import "testing"

// --- Testing code ---

func TestInsertionSortInt(t *testing.T) {
	arr := []int{5, 2, 9, 1, 5}
	InsertionSortInt(arr)
	expected := []int{1, 2, 5, 5, 9}
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("InsertionSortInt failed: got %v, want %v", arr, expected)
		}
	}
}

func TestInsertionSortIntDesc(t *testing.T) {
	arr := []int{3, 1, 4, 2}
	InsertionSortIntDesc(arr)
	expected := []int{4, 3, 2, 1}
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("InsertionSortIntDesc failed: got %v, want %v", arr, expected)
		}
	}
}

func TestInsertionSortFloat(t *testing.T) {
	arr := []float64{3.1, 2.2, 5.5, 1.0}
	InsertionSortFloat(arr)
	expected := []float64{1.0, 2.2, 3.1, 5.5}
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("InsertionSortFloat failed: got %v, want %v", arr, expected)
		}
	}
}

func TestInsertionSortFloatDesc(t *testing.T) {
	arr := []float64{3.1, 2.2, 5.5, 1.0}
	InsertionSortFloatDesc(arr)
	expected := []float64{5.5, 3.1, 2.2, 1.0}
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("InsertionSortFloatDesc failed: got %v, want %v", arr, expected)
		}
	}
}

func TestInsertionSortString(t *testing.T) {
	arr := []string{"banana", "apple", "cherry"}
	InsertionSortString(arr)
	expected := []string{"apple", "banana", "cherry"}
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("InsertionSortString failed: got %v, want %v", arr, expected)
		}
	}
}

func TestInsertionSortStringDesc(t *testing.T) {
	arr := []string{"banana", "apple", "cherry"}
	InsertionSortStringDesc(arr)
	expected := []string{"cherry", "banana", "apple"}
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("InsertionSortStringDesc failed: got %v, want %v", arr, expected)
		}
	}
}

func TestInsertionSortPersonByAge(t *testing.T) {
	arr := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}
	InsertionSortPersonByAge(arr)
	expected := []Person{
		{"Bob", 25},
		{"Alice", 30},
		{"Charlie", 35},
	}
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("InsertionSortPersonByAge failed: got %v, want %v", arr, expected)
		}
	}
}

func TestInsertionSortPersonByName(t *testing.T) {
	arr := []Person{
		{"Charlie", 35},
		{"Alice", 30},
		{"Bob", 25},
	}
	InsertionSortPersonByName(arr)
	expected := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("InsertionSortPersonByName failed: got %v, want %v", arr, expected)
		}
	}
}

func TestInsertionSortRune(t *testing.T) {
	arr := []rune{'z', 'a', 'm'}
	InsertionSortRune(arr)
	expected := []rune{'a', 'm', 'z'}
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("InsertionSortRune failed: got %v, want %v", arr, expected)
		}
	}
}

func TestInsertionSortByte(t *testing.T) {
	arr := []byte{'z', 'a', 'm'}
	InsertionSortByte(arr)
	expected := []byte{'a', 'm', 'z'}
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("InsertionSortByte failed: got %v, want %v", arr, expected)
		}
	}
}

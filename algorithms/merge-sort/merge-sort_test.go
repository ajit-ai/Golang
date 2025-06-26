package main

import "testing"

func TestMergeSortInt(t *testing.T) {
	arr := []int{5, 2, 9, 1, 5}
	sorted := MergeSortInt(arr)
	expected := []int{1, 2, 5, 5, 9}
	for i := range sorted {
		if sorted[i] != expected[i] {
			t.Errorf("MergeSortInt failed: got %v, want %v", sorted, expected)
		}
	}
}

func TestMergeSortIntDesc(t *testing.T) {
	arr := []int{3, 1, 4, 2}
	sorted := MergeSortIntDesc(arr)
	expected := []int{4, 3, 2, 1}
	for i := range sorted {
		if sorted[i] != expected[i] {
			t.Errorf("MergeSortIntDesc failed: got %v, want %v", sorted, expected)
		}
	}
}

func TestMergeSortFloat(t *testing.T) {
	arr := []float64{3.1, 2.2, 5.5, 1.0}
	sorted := MergeSortFloat(arr)
	expected := []float64{1.0, 2.2, 3.1, 5.5}
	for i := range sorted {
		if sorted[i] != expected[i] {
			t.Errorf("MergeSortFloat failed: got %v, want %v", sorted, expected)
		}
	}
}

func TestMergeSortFloatDesc(t *testing.T) {
	arr := []float64{3.1, 2.2, 5.5, 1.0}
	sorted := MergeSortFloatDesc(arr)
	expected := []float64{5.5, 3.1, 2.2, 1.0}
	for i := range sorted {
		if sorted[i] != expected[i] {
			t.Errorf("MergeSortFloatDesc failed: got %v, want %v", sorted, expected)
		}
	}
}

func TestMergeSortString(t *testing.T) {
	arr := []string{"banana", "apple", "cherry"}
	sorted := MergeSortString(arr)
	expected := []string{"apple", "banana", "cherry"}
	for i := range sorted {
		if sorted[i] != expected[i] {
			t.Errorf("MergeSortString failed: got %v, want %v", sorted, expected)
		}
	}
}

func TestMergeSortStringCI(t *testing.T) {
	arr := []string{"Banana", "apple", "Cherry"}
	sorted := MergeSortStringCI(arr)
	expected := []string{"apple", "Banana", "Cherry"}
	for i := range sorted {
		if sorted[i] != expected[i] {
			t.Errorf("MergeSortStringCI failed: got %v, want %v", sorted, expected)
		}
	}
}

func TestMergeSortRune(t *testing.T) {
	arr := []rune{'z', 'a', 'm'}
	sorted := MergeSortRune(arr)
	expected := []rune{'a', 'm', 'z'}
	for i := range sorted {
		if sorted[i] != expected[i] {
			t.Errorf("MergeSortRune failed: got %v, want %v", sorted, expected)
		}
	}
}

func TestMergeSortByte(t *testing.T) {
	arr := []byte{'z', 'a', 'm'}
	sorted := MergeSortByte(arr)
	expected := []byte{'a', 'm', 'z'}
	for i := range sorted {
		if sorted[i] != expected[i] {
			t.Errorf("MergeSortByte failed: got %v, want %v", sorted, expected)
		}
	}
}

func TestMergeSortPersonByAge(t *testing.T) {
	arr := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}
	sorted := MergeSortPersonByAge(arr)
	expected := []Person{
		{"Bob", 25},
		{"Alice", 30},
		{"Charlie", 35},
	}
	for i := range sorted {
		if sorted[i] != expected[i] {
			t.Errorf("MergeSortPersonByAge failed: got %v, want %v", sorted, expected)
		}
	}
}

func TestMergeSortPersonByName(t *testing.T) {
	arr := []Person{
		{"Charlie", 35},
		{"Alice", 30},
		{"Bob", 25},
	}
	sorted := MergeSortPersonByName(arr)
	expected := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}
	for i := range sorted {
		if sorted[i] != expected[i] {
			t.Errorf("MergeSortPersonByName failed: got %v, want %v", sorted, expected)
		}
	}
}

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

func TestMergeSortEdgeCases(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"empty", []int{}, []int{}},
		{"single", []int{7}, []int{7}},
		{"left exhausted first drops right tail", []int{1, 3, 2}, []int{1, 2, 3}},
		{"reverse sorted", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"duplicates", []int{2, 2, 1, 1}, []int{1, 1, 2, 2}},
		{"negatives", []int{-3, 5, -1, 0}, []int{-3, -1, 0, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeSortInt(tt.in)
			if len(got) != len(tt.want) {
				t.Fatalf("MergeSortInt(%v) = %v, want %v", tt.in, got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("MergeSortInt(%v) = %v, want %v", tt.in, got, tt.want)
					break
				}
			}
		})
	}
}

func TestMergeSortStringCILeftExhausted(t *testing.T) {
	sorted := MergeSortStringCI([]string{"apple", "Zebra", "monkey"})
	expected := []string{"apple", "monkey", "Zebra"}
	for i := range sorted {
		if sorted[i] != expected[i] {
			t.Fatalf("MergeSortStringCI left-exhausted case failed: got %v, want %v", sorted, expected)
		}
	}
}

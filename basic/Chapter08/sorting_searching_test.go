package main

import (
	"reflect"
	"testing"
)

func TestInsertionSorter(t *testing.T) {
	elements := []int{5, 2, 9, 1, 5}
	InsertionSorter(elements)
	want := []int{1, 2, 5, 5, 9}
	if !reflect.DeepEqual(elements, want) {
		t.Errorf("InsertionSorter = %v, want %v", elements, want)
	}
}

func TestMergeSorter(t *testing.T) {
	got := MergeSorter([]int{5, 2, 9, 1})
	want := []int{1, 2, 5, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MergeSorter = %v, want %v", got, want)
	}
}

func TestJoinArrays(t *testing.T) {
	got := JoinArrays([]int{1, 3}, []int{2, 4})
	want := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("JoinArrays = %v, want %v", got, want)
	}
}

func TestQuickSorter(t *testing.T) {
	elements := []int{5, 2, 9, 1, 5, 6}
	QuickSorter(elements, 0, len(elements)-1)
	want := []int{1, 2, 5, 5, 6, 9}
	if !reflect.DeepEqual(elements, want) {
		t.Errorf("QuickSorter = %v, want %v", elements, want)
	}
}

func TestShellSorter(t *testing.T) {
	elements := []int{5, 2, 9, 1}
	ShellSorter(elements)
	want := []int{1, 2, 5, 9}
	if !reflect.DeepEqual(elements, want) {
		t.Errorf("ShellSorter = %v, want %v", elements, want)
	}
}

func TestLinearSearch(t *testing.T) {
	elements := []int{10, 20, 30, 40}
	if !LinearSearch(elements, 30) {
		t.Error("LinearSearch should find 30")
	}
	if LinearSearch(elements, 99) {
		t.Error("LinearSearch should not find 99")
	}
}

func TestInterpolationSearch(t *testing.T) {
	elements := []int{10, 20, 30, 40, 50}
	found, index := InterpolationSearch(elements, 40)
	if !found || index != 3 {
		t.Errorf("InterpolationSearch(40) = %v, %d; want true, 3", found, index)
	}
	found, _ = InterpolationSearch(elements, 35)
	if found {
		t.Error("InterpolationSearch should not find 35")
	}
}

func TestFactor(t *testing.T) {
	tests := map[int]int{0: 1, 1: 1, 5: 120, 7: 5040}
	for n, want := range tests {
		if got := Factor(n); got != want {
			t.Errorf("Factor(%d) = %d, want %d", n, got, want)
		}
	}
}

func TestStringHashDeterministic(t *testing.T) {
	if stringHash("abc") != stringHash("abc") {
		t.Error("stringHash should be deterministic")
	}
	if stringHash("abc") == stringHash("abd") {
		t.Error("stringHash of different inputs should differ")
	}
}

func TestCreateHashAndXor(t *testing.T) {
	h1 := CreateHash([]byte("hello"))
	h2 := CreateHash([]byte("hello"))
	if !reflect.DeepEqual(h1, h2) {
		t.Error("CreateHash should be deterministic")
	}
	combined := CreateHashMultiple([]byte("a"), []byte("b"))
	if len(combined) == 0 {
		t.Error("CreateHashMultiple returned empty hash")
	}
	x := xor([]byte{0xF0, 0x0F}, []byte{0xFF, 0xFF})
	want := []byte{0x0F, 0xF0}
	if !reflect.DeepEqual(x, want) {
		t.Errorf("xor = %v, want %v", x, want)
	}
}

func TestChapter08MainSmoke(t *testing.T) {
	main()
}

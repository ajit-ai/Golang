package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestGenericStack(t *testing.T) {
	s := &GenericStack[int]{}
	if _, ok := s.Pop(); ok {
		t.Error("Pop on empty stack should report not-ok")
	}
	s.Push(1)
	s.Push(2)
	if v, ok := s.Pop(); !ok || v != 2 {
		t.Errorf("Pop = %d, %v; want 2, true", v, ok)
	}
	if v, ok := s.Pop(); !ok || v != 1 {
		t.Errorf("Pop = %d, %v; want 1, true", v, ok)
	}
	if s.Len() != 0 {
		t.Errorf("Len = %d, want 0", s.Len())
	}
}

func TestGenericStackOfString(t *testing.T) {
	s := &GenericStack[string]{}
	s.Push("go")
	if v, _ := s.Pop(); v != "go" {
		t.Errorf("Pop = %q, want go", v)
	}
}

func TestMapSliceAndFilter(t *testing.T) {
	got := MapSlice([]int{1, 2, 3}, func(v int) string {
		return string(rune('a' + v - 1))
	})
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MapSlice = %v, want %v", got, want)
	}

	evens := Filter([]int{1, 2, 3, 4, 5, 6}, func(v int) bool { return v%2 == 0 })
	if !reflect.DeepEqual(evens, []int{2, 4, 6}) {
		t.Errorf("Filter evens = %v", evens)
	}
	if got := Filter([]int{}, func(v int) bool { return true }); len(got) != 0 {
		t.Error("Filter of empty slice should be empty")
	}
}

func TestMinMaxAndSum(t *testing.T) {
	if Min(3, 7) != 3 || Max(3, 7) != 7 {
		t.Error("Min/Max wrong for ints")
	}
	if Min("apple", "banana") != "apple" {
		t.Error("Min wrong for strings")
	}
	if got := Sum([]int{1, 2, 3}); got != 6 {
		t.Errorf("Sum ints = %d, want 6", got)
	}
	if got := Sum([]float64{0.5, 0.25}); got != 0.75 {
		t.Errorf("Sum floats = %v, want 0.75", got)
	}
	var empty []int
	if Sum(empty) != 0 {
		t.Error("Sum of empty should be 0")
	}
}

func TestKeys(t *testing.T) {
	got := Keys(map[int]string{1: "a", 2: "b"})
	sort.Ints(got)
	if !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("Keys = %v", got)
	}
}

func TestGenericsMainSmoke(t *testing.T) {
	GenericStackMain()
	GenericAlgorithmsMain()
}

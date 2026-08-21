package main

import "testing"

func TestBuildIntList(t *testing.T) {
	intList := BuildIntList(11, 23, 34)
	if intList.Len() != 3 {
		t.Fatalf("Len() = %d, want 3", intList.Len())
	}
	var got []int
	for element := intList.Front(); element != nil; element = element.Next() {
		got = append(got, element.Value.(int))
	}
	want := []int{11, 23, 34}
	if len(got) != len(want) {
		t.Fatalf("values = %v, want %v", got, want)
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("values[%d] = %d, want %d", i, got[i], want[i])
		}
	}
}

func TestBuildIntListEmpty(t *testing.T) {
	intList := BuildIntList()
	if intList.Len() != 0 {
		t.Errorf("Len() = %d, want 0", intList.Len())
	}
	if intList.Front() != nil || intList.Back() != nil {
		t.Error("empty list should have no front or back element")
	}
}

func TestListMain(t *testing.T) {
	ListMain()
}

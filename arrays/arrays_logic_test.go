package main

import "testing"

func TestMakeArray(t *testing.T) {
	got := makeArray()
	want := [3]int{7, 8, 9}
	if got != want {
		t.Errorf("makeArray() = %v, want %v", got, want)
	}
}

func TestArraysAreValues(t *testing.T) {
	original := [3]int{1, 2, 3}
	copied := original
	copied[0] = 99
	if original[0] != 1 {
		t.Error("arrays should be copied by value; original was mutated")
	}
}

func TestArrayExMainSmoke(t *testing.T) {
	ArrayExMain()
}

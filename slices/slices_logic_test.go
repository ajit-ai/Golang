package main

import (
	"reflect"
	"testing"
)

func TestMakeSlice(t *testing.T) {
	got := makeSlice()
	want := []int{7, 8, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("makeSlice() = %v, want %v", got, want)
	}
}

func TestSlicesUseMainSmoke(t *testing.T) {
	SlicesUseMain()
}

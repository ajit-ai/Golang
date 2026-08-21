package gostack

import (
	"reflect"
	"testing"
)

func TestZeroValueStack(t *testing.T) {
	var s Stack[int]
	if _, ok := s.Pop(); ok {
		t.Error("Pop on zero-value stack should report not-ok")
	}
	if _, ok := s.Peek(); ok {
		t.Error("Peek on zero-value stack should report not-ok")
	}
	if s.Len() != 0 {
		t.Errorf("Len = %d, want 0", s.Len())
	}
}

func TestPushPopLIFO(t *testing.T) {
	s := New[string](2)
	s.Push("a")
	s.Push("b")

	if v, _ := s.Peek(); v != "b" {
		t.Errorf("Peek = %q, want b", v)
	}
	if v, ok := s.Pop(); !ok || v != "b" {
		t.Errorf("Pop = %q, %v", v, ok)
	}
	if v, ok := s.Pop(); !ok || v != "a" {
		t.Errorf("Pop = %q, %v", v, ok)
	}
	if _, ok := s.Pop(); ok {
		t.Error("stack should be empty")
	}
}

func TestDrainOrderAndEmptiness(t *testing.T) {
	s := New[int](3)
	for _, v := range []int{1, 2, 3} {
		s.Push(v)
	}
	got := s.Drain()
	if !reflect.DeepEqual(got, []int{3, 2, 1}) {
		t.Errorf("Drain = %v, want [3 2 1]", got)
	}
	if s.Len() != 0 {
		t.Errorf("after Drain Len = %d, want 0", s.Len())
	}
}

func TestStructTypeElement(t *testing.T) {
	type point struct{ X, Y int }
	s := New[point](1)
	s.Push(point{1, 2})
	got, _ := s.Pop()
	if got != (point{1, 2}) {
		t.Errorf("Pop = %+v", got)
	}
}

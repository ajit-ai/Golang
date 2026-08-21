package main

import "testing"

func newTestSet(elements ...int) *Set {
	s := &Set{}
	s.New()
	for _, e := range elements {
		s.AddElement(e)
	}
	return s
}

func TestSetAddContainsDelete(t *testing.T) {
	s := newTestSet(1, 2)
	if !s.ContainsElement(1) || !s.ContainsElement(2) {
		t.Error("set should contain added elements")
	}
	s.DeleteElement(1)
	if s.ContainsElement(1) {
		t.Error("element should be deleted")
	}
}

func TestSetNoDuplicates(t *testing.T) {
	s := newTestSet(1, 1, 1)
	count := 0
	for range s.integerMap {
		count++
	}
	if count != 1 {
		t.Errorf("set size = %d, want 1 after duplicate adds", count)
	}
}

func TestSetIntersect(t *testing.T) {
	a := newTestSet(1, 2, 3)
	b := newTestSet(2, 3, 4)
	inter := a.Intersect(b)
	if !inter.ContainsElement(2) || !inter.ContainsElement(3) {
		t.Error("intersection missing common elements")
	}
	if inter.ContainsElement(1) || inter.ContainsElement(4) {
		t.Error("intersection contains non-common element")
	}
}

func TestSetUnion(t *testing.T) {
	a := newTestSet(1, 2)
	b := newTestSet(2, 3)
	un := a.Union(b)
	for _, e := range []int{1, 2, 3} {
		if !un.ContainsElement(e) {
			t.Errorf("union missing element %d", e)
		}
	}
}

func TestSetMainSmoke(t *testing.T) {
	SetMain()
}

package main

import "testing"

func TestMarkMarksReachableGraph(t *testing.T) {
	c := newObject(3, 0)
	b := newObject(2, 0, c)
	a := newObject(1, 0, b)

	Mark(a)

	for _, o := range []*object{a, b, c} {
		if !IfMarked(o) {
			t.Errorf("object %d reachable from root was not marked", o.id)
		}
	}
}

func TestMarkIsIdempotent(t *testing.T) {
	o := newObject(1, 0)
	Mark(o)
	Mark(o)
	if !IfMarked(o) {
		t.Error("object lost its mark after repeated Mark calls")
	}
}

func TestGetReferences(t *testing.T) {
	child := newObject(2, 0)
	parent := newObject(1, 0, child)
	refs := GetReferences(parent)
	if len(refs) != 1 || refs[0] != child {
		t.Errorf("GetReferences = %v, want [%v]", refs, child)
	}
}

func TestMarkMainSmoke(t *testing.T) {
	MarkMain()
}

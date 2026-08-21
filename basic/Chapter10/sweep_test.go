package main

import "testing"

func TestSweepReleasesOnlyUnmarked(t *testing.T) {
	kept := newObject(1, 0)
	freed := newObject(2, 0)
	registerObject(kept)
	registerObject(freed)
	SetMarked(kept)

	Sweep()

	for _, o := range GetObjects() {
		if o == freed {
			t.Error("unmarked object survived sweep")
		}
	}
	foundKept := false
	for _, o := range GetObjects() {
		if o == kept {
			foundKept = true
		}
	}
	if !foundKept {
		t.Error("marked object was released")
	}
	if !IfMarked(kept) {
		t.Error("marked object lost its mark")
	}
}

func TestSweepWithNothingMarkedReleasesAllNewcomers(t *testing.T) {
	o := newObject(9, 0)
	registerObject(o)

	Sweep()

	for _, live := range GetObjects() {
		if live == o {
			t.Error("unmarked object was not released")
		}
	}
}

func TestSweepMainSmoke(t *testing.T) {
	SweepMain()
}

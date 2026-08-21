package main

import "testing"

func TestGenerationCollectMarksOnlyOldGeneration(t *testing.T) {
	old1 := newObject(1, 3)
	old2 := newObject(2, 3)
	young := newObject(3, 1)
	registerObject(old1)
	registerObject(old2)
	registerObject(young)

	GenerationCollect()

	if !IfMarked(old1) || !IfMarked(old2) {
		t.Error("generation-3 objects were not marked")
	}
	if IfMarked(young) {
		t.Error("generation-1 object should not be marked by GenerationCollect")
	}
}

func TestGetObjectsFromOldGeneration(t *testing.T) {
	old := newObject(1, 2)
	young := newObject(2, 0)
	registerObject(old)
	registerObject(young)

	got := GetObjectsFromOldGeneration(2)
	if len(got) != 1 || got[0] != old {
		t.Errorf("GetObjectsFromOldGeneration(2) = %v, want [%v]", got, old)
	}
}

func TestGenerationCollectMainSmoke(t *testing.T) {
	GenerationCollectMain()
}

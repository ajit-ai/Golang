package main

import "testing"

func TestWeightedReference(t *testing.T) {
	if got := WeightedReference(); got != 42 {
		t.Errorf("WeightedReference() = %d, want 42", got)
	}
}

func TestGetWeightedReferences(t *testing.T) {
	refs := GetWeightedReferences()
	if len(refs) != 3 {
		t.Fatalf("len(GetWeightedReferences()) = %d, want 3", len(refs))
	}
	sum := 0
	for _, r := range refs {
		sum += r.weight
	}
	if sum != 42 {
		t.Errorf("sum of weights = %d, want 42", sum)
	}
}

func TestWeightedReferenceMainSmoke(t *testing.T) {
	WeightedReferenceMain()
}

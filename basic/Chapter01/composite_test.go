package main

import "testing"

var _ IComposite = &Leaflet{}
var _ IComposite = &Branch{}

func TestBranchAdd(t *testing.T) {
	branch := &Branch{name: "branch"}
	branch.add(Leaflet{name: "leaf 1"})
	branch.add(Leaflet{name: "leaf 2"})
	leaflets := branch.getLeaflets()
	if len(leaflets) != 2 {
		t.Fatalf("len(leaflets) = %d, want 2", len(leaflets))
	}
	want := []string{"leaf 1", "leaf 2"}
	for i, leaf := range leaflets {
		if leaf.name != want[i] {
			t.Errorf("leaflets[%d].name = %q, want %q", i, leaf.name, want[i])
		}
	}
}

func TestBranchAddBranch(t *testing.T) {
	branch := &Branch{name: "branch"}
	branch.addBranch(Branch{name: "branch 2"})
	if len(branch.branches) != 1 {
		t.Fatalf("len(branches) = %d, want 1", len(branch.branches))
	}
	if branch.branches[0].name != "branch 2" {
		t.Errorf("branches[0].name = %q, want %q", branch.branches[0].name, "branch 2")
	}
}

func TestPerform(t *testing.T) {
	branch := &Branch{name: "branch"}
	branch.add(Leaflet{name: "leaf"})
	branch.addBranch(Branch{name: "branch 2"})
	branch.perform()
}

func TestCompositeMain(t *testing.T) {
	CompositeMain()
}

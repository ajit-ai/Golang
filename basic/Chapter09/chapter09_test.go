package main

import "testing"

func TestNewKnowledgeGraphConstructor(t *testing.T) {
	if kg := NewKnowledgeGraph(); kg == nil {
		t.Error("NewKnowledgeGraph returned nil")
	}
}

func TestNewMapLayoutConstructor(t *testing.T) {
	if ml := NewMapLayout(); ml == nil {
		t.Error("NewMapLayout returned nil")
	}
}

func TestNewSocialGraphConstructor(t *testing.T) {
	if sg := NewSocialGraph(1); sg == nil {
		t.Error("NewSocialGraph returned nil")
	}
}

func TestExampleSocialGraphAddEntityAndLink(t *testing.T) {
	g := NewExampleSocialGraph()
	root := Name("Root")
	leaf := Name("Leaf")

	g.AddEntity(root)
	g.AddEntity(leaf)
	if _, ok := g.GraphNodes[root]; !ok {
		t.Error("AddEntity did not register root")
	}

	g.AddLink(root, leaf)
	links := g.Links[root]
	if _, ok := links[leaf]; !ok {
		t.Error("AddLink did not record root -> leaf")
	}
}

func TestSparseMatrixDimensions(t *testing.T) {
	m := NewSparseMatrix(3, 4)
	if m == nil {
		t.Fatal("NewSparseMatrix returned nil")
	}
	if rows, cols := m.Shape(); rows != 3 || cols != 4 {
		t.Errorf("matrix dimensions = %dx%d, want 3x4", rows, cols)
	}
}

func TestChapter09MainSmoke(t *testing.T) {
	main()
}

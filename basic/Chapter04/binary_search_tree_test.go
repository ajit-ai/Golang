package main

import "testing"

func newTestBST(values ...int) *BinarySearchTree {
	tree := &BinarySearchTree{}
	tree.rootNode = nil
	for _, v := range values {
		tree.InsertElement(v, v)
	}
	return tree
}

func bstInOrder(tree *BinarySearchTree) []int {
	var got []int
	tree.InOrderTraverseTree(func(v int) {
		got = append(got, v)
	})
	return got
}

func intsEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestBSTInsertInOrder(t *testing.T) {
	tree := newTestBST(5, 3, 8, 1, 4, 7, 9)
	got := bstInOrder(tree)
	want := []int{1, 3, 4, 5, 7, 8, 9}
	if !intsEqual(got, want) {
		t.Errorf("in-order = %v, want %v", got, want)
	}
}

func TestBSTMinMax(t *testing.T) {
	tree := newTestBST(5, 3, 8, 1, 9)
	if min := tree.MinNode(); min == nil || *min != 1 {
		t.Errorf("MinNode = %v, want 1", min)
	}
	if max := tree.MaxNode(); max == nil || *max != 9 {
		t.Errorf("MaxNode = %v, want 9", max)
	}
}

func TestBSTSearch(t *testing.T) {
	tree := newTestBST(5, 3, 8)
	if !tree.SearchNode(3) || !tree.SearchNode(8) {
		t.Error("existing keys not found")
	}
	if tree.SearchNode(99) {
		t.Error("missing key reported found")
	}
}

func TestBSTRemove(t *testing.T) {
	tree := newTestBST(5, 3, 8, 1, 4)
	tree.RemoveNode(3)
	got := bstInOrder(tree)
	want := []int{1, 4, 5, 8}
	if !intsEqual(got, want) {
		t.Errorf("after remove in-order = %v, want %v", got, want)
	}
}

func TestBinarySearchTreeMainSmoke(t *testing.T) {
	BinarySearchTreeMain()
}

func TestBSTRemoveRootNode(t *testing.T) {
	tree := newTestBST(5, 3, 8)
	tree.RemoveNode(5)
	if tree.SearchNode(5) {
		t.Error("root node still searchable after removal")
	}
	got := bstInOrder(tree)
	want := []int{3, 8}
	if !intsEqual(got, want) {
		t.Errorf("after root removal in-order = %v, want %v", got, want)
	}
}

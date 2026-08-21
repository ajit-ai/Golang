package main

import "testing"

func buildTestTree() *Tree {
	tree := &Tree{nil, 1, nil}
	tree.Insert(3)
	tree.Insert(5)
	tree.LeftNode.Insert(7)
	return tree
}

func TestInsert(t *testing.T) {
	tree := buildTestTree()
	if tree.Value != 1 {
		t.Errorf("root value = %d, want 1", tree.Value)
	}
	if tree.LeftNode == nil || tree.LeftNode.Value != 3 {
		t.Errorf("left child = %v, want value 3", tree.LeftNode)
	}
	if tree.RightNode == nil || tree.RightNode.Value != 5 {
		t.Errorf("right child = %v, want value 5", tree.RightNode)
	}
	if tree.LeftNode.LeftNode == nil || tree.LeftNode.LeftNode.Value != 7 {
		t.Errorf("left-left grandchild = %v, want value 7", tree.LeftNode.LeftNode)
	}
	if tree.LeftNode.RightNode != nil {
		t.Error("left-right grandchild should be nil")
	}
	if tree.RightNode.LeftNode != nil || tree.RightNode.RightNode != nil {
		t.Error("right grandchildren should be nil")
	}
}

func TestInsertOnNilTreeDoesNothing(t *testing.T) {
	var tree *Tree
	tree.Insert(5)
	if tree != nil {
		t.Error("insert on nil receiver should leave the pointer nil")
	}
}

func TestInorderValues(t *testing.T) {
	tests := []struct {
		name string
		tree *Tree
		want []int
	}{
		{"nil tree", nil, nil},
		{"single node", &Tree{nil, 1, nil}, []int{1}},
		{"demo tree", buildTestTree(), []int{7, 3, 1, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InorderValues(tt.tree)
			if len(got) != len(tt.want) {
				t.Fatalf("InorderValues() = %v, want %v", got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("InorderValues() = %v, want %v", got, tt.want)
					break
				}
			}
		})
	}
}

func TestTreeMain(t *testing.T) {
	TreeMain()
}

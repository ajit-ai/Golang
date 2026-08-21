package main

import "testing"

func avlInOrder(node *AVLTreeNode, acc *[]int) {
	if node == nil {
		return
	}
	avlInOrder(node.LinkedNodes[0], acc)
	*acc = append(*acc, int(node.KeyValue.(integerKey)))
	avlInOrder(node.LinkedNodes[1], acc)
}

func TestAvlTreeInsertKeepsSortedOrder(t *testing.T) {
	var tree *AVLTreeNode
	for _, k := range []int{5, 3, 8, 7, 6, 10, 1} {
		InsertNode(&tree, integerKey(k))
	}
	var got []int
	avlInOrder(tree, &got)
	want := []int{1, 3, 5, 6, 7, 8, 10}
	if len(got) != len(want) {
		t.Fatalf("in-order = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("in-order = %v, want %v", got, want)
		}
	}
}

func TestAvlTreeRemove(t *testing.T) {
	var tree *AVLTreeNode
	for _, k := range []int{5, 3, 8} {
		InsertNode(&tree, integerKey(k))
	}
	RemoveNode(&tree, integerKey(3))

	var got []int
	avlInOrder(tree, &got)
	if len(got) != 2 || got[0] != 5 || got[1] != 8 {
		t.Errorf("after remove in-order = %v, want [5 8]", got)
	}
}

func TestAvlTreeMainSmoke(t *testing.T) {
	AvlTreeMain()
}

func TestCircularListMainSmoke(t *testing.T) {
	CircularListMain()
}

func TestHashMainSmoke(t *testing.T) {
	HashMain()
}

func TestTableMainSmoke(t *testing.T) {
	TableMain()
}

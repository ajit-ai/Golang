// main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// Tree struct
type Tree struct {
	LeftNode  *Tree
	Value     int
	RightNode *Tree
}

// Tree insert method for inserting at m position
func (tree *Tree) Insert(m int) {
	if tree != nil {

		if tree.LeftNode == nil {
			tree.LeftNode = &Tree{nil, m, nil}
		} else {
			if tree.RightNode == nil {
				tree.RightNode = &Tree{nil, m, nil}
			} else {

				if tree.LeftNode != nil {

					tree.LeftNode.Insert(m)
				} else {

					tree.RightNode.Insert(m)
				}

			}

		}

	} else {
		tree = &Tree{nil, m, nil}
	}
}

// InorderValues collects the values of the tree in inorder traversal
func InorderValues(tree *Tree) []int {
	if tree == nil {
		return nil
	}

	left := InorderValues(tree.LeftNode)
	values := append(left, tree.Value)
	return append(values, InorderValues(tree.RightNode)...)
}

// Print method for printing a Tree
func Print(tree *Tree) {
	if tree != nil {

		fmt.Println(" Value", tree.Value)
		fmt.Printf("Tree Node Left")
		Print(tree.LeftNode)
		fmt.Printf("Tree Node Right")
		Print(tree.RightNode)
	} else {
		fmt.Printf("Nil\n")
	}
}

// TreeMain method
func TreeMain() {
	var tree *Tree = &Tree{nil, 1, nil}
	Print(tree)
	tree.Insert(3)
	Print(tree)
	tree.Insert(5)
	Print(tree)
	tree.LeftNode.Insert(7)
	Print(tree)

}

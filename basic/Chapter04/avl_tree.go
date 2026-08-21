// /main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"encoding/json"
	"fmt"
)

// KeyValue type
type KeyValue interface {
	LessThan(KeyValue) bool
	EqualTo(KeyValue) bool
}

// AVLTreeNode class
type AVLTreeNode struct {
	KeyValue     KeyValue
	BalanceValue int
	LinkedNodes  [2]*AVLTreeNode
}

// opposite method
func opposite(nodeValue int) int {
	return 1 - nodeValue
}

// single rotation method
func singleRotation(rootNode *AVLTreeNode, nodeValue int) *AVLTreeNode {

	var saveNode *AVLTreeNode
	saveNode = rootNode.LinkedNodes[opposite(nodeValue)]
	rootNode.LinkedNodes[opposite(nodeValue)] = saveNode.LinkedNodes[nodeValue]
	saveNode.LinkedNodes[nodeValue] = rootNode
	return saveNode
}

// double rotation method
func doubleRotation(rootNode *AVLTreeNode, nodeValue int) *AVLTreeNode {

	var saveNode *AVLTreeNode
	saveNode = rootNode.LinkedNodes[opposite(nodeValue)].LinkedNodes[nodeValue]

	rootNode.LinkedNodes[opposite(nodeValue)].LinkedNodes[nodeValue] = saveNode.LinkedNodes[opposite(nodeValue)]
	saveNode.LinkedNodes[opposite(nodeValue)] = rootNode.LinkedNodes[opposite(nodeValue)]
	rootNode.LinkedNodes[opposite(nodeValue)] = saveNode

	saveNode = rootNode.LinkedNodes[opposite(nodeValue)]
	rootNode.LinkedNodes[opposite(nodeValue)] = saveNode.LinkedNodes[nodeValue]
	saveNode.LinkedNodes[nodeValue] = rootNode
	return saveNode
}

// adjust balance method
func adjustBalance(rootNode *AVLTreeNode, nodeValue int, balanceValue int) {

	var node *AVLTreeNode
	node = rootNode.LinkedNodes[nodeValue]
	var oppNode *AVLTreeNode
	oppNode = node.LinkedNodes[opposite(nodeValue)]
	switch oppNode.BalanceValue {
	case 0:
		rootNode.BalanceValue = 0
		node.BalanceValue = 0
	case balanceValue:
		rootNode.BalanceValue = -balanceValue
		node.BalanceValue = 0
	default:
		rootNode.BalanceValue = 0
		node.BalanceValue = balanceValue
	}
	oppNode.BalanceValue = 0
}

// balanceTree method
func BalanceTree(rootNode *AVLTreeNode, nodeValue int) *AVLTreeNode {
	var node *AVLTreeNode
	node = rootNode.LinkedNodes[nodeValue]
	var balance int
	balance = 2*nodeValue - 1
	if node.BalanceValue == balance {
		rootNode.BalanceValue = 0
		node.BalanceValue = 0
		return singleRotation(rootNode, opposite(nodeValue))
	}
	adjustBalance(rootNode, nodeValue, balance)
	return doubleRotation(rootNode, opposite(nodeValue))
}

// insertRNode method
func insertRNode(rootNode *AVLTreeNode, key KeyValue) (*AVLTreeNode, bool) {
	if rootNode == nil {
		return &AVLTreeNode{KeyValue: key}, false
	}
	var dir int
	dir = 0
	if rootNode.KeyValue.LessThan(key) {
		dir = 1
	}
	var done bool
	rootNode.LinkedNodes[dir], done = insertRNode(rootNode.LinkedNodes[dir], key)
	if done {
		return rootNode, true
	}
	rootNode.BalanceValue = rootNode.BalanceValue + (2*dir - 1)
	switch rootNode.BalanceValue {
	case 0:
		return rootNode, true
	case 1, -1:
		return rootNode, false
	}
	return BalanceTree(rootNode, dir), true
}

// InsertNode method
func InsertNode(AVLTreeNode **AVLTreeNode, key KeyValue) {
	*AVLTreeNode, _ = insertRNode(*AVLTreeNode, key)
}

// RemoveNode method
func RemoveNode(AVLTreeNode **AVLTreeNode, key KeyValue) {
	*AVLTreeNode, _ = removeRNode(*AVLTreeNode, key)
}

// removeBalance method
func removeBalance(rootNode *AVLTreeNode, nodeValue int) (*AVLTreeNode, bool) {
	var node *AVLTreeNode
	node = rootNode.LinkedNodes[opposite(nodeValue)]
	var balance int
	balance = 2*nodeValue - 1
	switch node.BalanceValue {
	case -balance:
		rootNode.BalanceValue = 0
		node.BalanceValue = 0
		return singleRotation(rootNode, nodeValue), false
	case balance:
		adjustBalance(rootNode, opposite(nodeValue), -balance)
		return doubleRotation(rootNode, nodeValue), false
	}
	rootNode.BalanceValue = -balance
	node.BalanceValue = balance
	return singleRotation(rootNode, nodeValue), true
}

// removeRNode method
func removeRNode(rootNode *AVLTreeNode, key KeyValue) (*AVLTreeNode, bool) {
	if rootNode == nil {
		return nil, false
	}
	if rootNode.KeyValue.EqualTo(key) {
		switch {
		case rootNode.LinkedNodes[0] == nil:
			return rootNode.LinkedNodes[1], false
		case rootNode.LinkedNodes[1] == nil:
			return rootNode.LinkedNodes[0], false
		}
		var heirNode *AVLTreeNode
		heirNode = rootNode.LinkedNodes[0]
		for heirNode.LinkedNodes[1] != nil {
			heirNode = heirNode.LinkedNodes[1]
		}
		rootNode.KeyValue = heirNode.KeyValue
		key = heirNode.KeyValue
	}
	var dir int
	dir = 0
	if rootNode.KeyValue.LessThan(key) {
		dir = 1
	}
	var done bool
	rootNode.LinkedNodes[dir], done = removeRNode(rootNode.LinkedNodes[dir], key)
	if done {
		return rootNode, true
	}
	rootNode.BalanceValue = rootNode.BalanceValue + (1 - 2*dir)
	switch rootNode.BalanceValue {
	case 1, -1:
		return rootNode, true
	case 0:
		return rootNode, false
	}
	return removeBalance(rootNode, dir)
}

type integerKey int

func (k integerKey) LessThan(k1 KeyValue) bool { return k < k1.(integerKey) }
func (k integerKey) EqualTo(k1 KeyValue) bool  { return k == k1.(integerKey) }

// main method
func AvlTreeMain() {
	var AVLTreeNode *AVLTreeNode
	fmt.Println("Tree is empty")
	var avlTree []byte
	avlTree, _ = json.MarshalIndent(AVLTreeNode, "", "   ")
	fmt.Println(string(avlTree))

	fmt.Println("\n Add Tree")
	InsertNode(&AVLTreeNode, integerKey(5))
	InsertNode(&AVLTreeNode, integerKey(3))
	InsertNode(&AVLTreeNode, integerKey(8))
	InsertNode(&AVLTreeNode, integerKey(7))
	InsertNode(&AVLTreeNode, integerKey(6))
	InsertNode(&AVLTreeNode, integerKey(10))
	avlTree, _ = json.MarshalIndent(AVLTreeNode, "", "   ")
	fmt.Println(string(avlTree))

	fmt.Println("\n Delete Tree")
	RemoveNode(&AVLTreeNode, integerKey(3))
	RemoveNode(&AVLTreeNode, integerKey(7))
	avlTree, _ = json.MarshalIndent(AVLTreeNode, "", "   ")
	fmt.Println(string(avlTree))
}

// main runs the demo entry points of this package
func main() {
	AvlTreeMain()
	BinarySearchTreeMain()
	CircularListMain()
	HashMain()
	TableMain()
}

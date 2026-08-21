// main package has examples shown
// in Hands-On Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// SinglyNode class
type SinglyNode struct {
	property int
	nextNode *SinglyNode
}

// SinglyLinkedList class
type SinglyLinkedList struct {
	headNode *SinglyNode
}

// AddToHead method of SinglyLinkedList class
func (slist *SinglyLinkedList) AddToHead(property int) {
	var snode = &SinglyNode{}
	snode.property = property
	snode.nextNode = nil

	if slist.headNode != nil {
		snode.nextNode = slist.headNode
	}

	slist.headNode = snode

}

// NodeWithValue method returns SinglyNode given parameter property
func (slist *SinglyLinkedList) NodeWithValue(property int) *SinglyNode {
	var snode *SinglyNode
	var nodeWith *SinglyNode
	for snode = slist.headNode; snode != nil; snode = snode.nextNode {
		if snode.property == property {
			nodeWith = snode
			break
		}
	}
	return nodeWith
}

// AddAfter method  adds a SinglyNode with nodeProperty after SinglyNode with property
func (slist *SinglyLinkedList) AddAfter(nodeProperty int, property int) {
	var snode = &SinglyNode{}
	snode.property = property
	snode.nextNode = nil

	var nodeWith *SinglyNode

	nodeWith = slist.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		snode.nextNode = nodeWith.nextNode
		nodeWith.nextNode = snode
	}

}

// LastNode method returns the last SinglyNode
func (slist *SinglyLinkedList) LastNode() *SinglyNode {
	var snode *SinglyNode
	var lastNode *SinglyNode
	for snode = slist.headNode; snode != nil; snode = snode.nextNode {
		if snode.nextNode == nil {
			lastNode = snode
		}
	}
	return lastNode
}

// AddToEnd method adds the SinglyNode with property to the end
func (slist *SinglyLinkedList) AddToEnd(property int) {
	var snode = &SinglyNode{}
	snode.property = property
	snode.nextNode = nil

	var lastNode *SinglyNode

	lastNode = slist.LastNode()

	if lastNode != nil {
		lastNode.nextNode = snode
	}

}

// IterateList method iterates over SinglyLinkedList
func (slist *SinglyLinkedList) IterateList() {

	var snode *SinglyNode
	for snode = slist.headNode; snode != nil; snode = snode.nextNode {
		fmt.Println(snode.property)

	}
}

// SinglyLinkedListMain method
func SinglyLinkedListMain() {

	var slist SinglyLinkedList

	slist.AddToHead(1)
	slist.AddToHead(3)
	slist.AddToEnd(5)
	slist.AddAfter(1, 7)

	slist.IterateList()

}

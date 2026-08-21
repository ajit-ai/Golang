// main package has examples shown
// in Hands-On Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// DoublyNode class
type DoublyNode struct {
	property     int
	nextNode     *DoublyNode
	previousNode *DoublyNode
}

// DoublyLinkedList class
type DoublyLinkedList struct {
	headNode *DoublyNode
}

// AddToHead method of DoublyLinkedList
func (DoublyLinkedList *DoublyLinkedList) AddToHead(property int) {
	var dnode = &DoublyNode{}
	dnode.property = property
	dnode.nextNode = nil
	if DoublyLinkedList.headNode != nil {
		dnode.nextNode = DoublyLinkedList.headNode
		DoublyLinkedList.headNode.previousNode = dnode
	}

	DoublyLinkedList.headNode = dnode

}

// NodeWithValue method of DoublyLinkedList
func (DoublyLinkedList *DoublyLinkedList) NodeWithValue(property int) *DoublyNode {
	var dnode *DoublyNode
	var nodeWith *DoublyNode
	for dnode = DoublyLinkedList.headNode; dnode != nil; dnode = dnode.nextNode {
		if dnode.property == property {
			nodeWith = dnode
			break
		}
	}
	return nodeWith
}

// AddAfter method of DoublyLinkedList
func (DoublyLinkedList *DoublyLinkedList) AddAfter(nodeProperty int, property int) {
	var dnode = &DoublyNode{}
	dnode.property = property
	dnode.nextNode = nil

	var nodeWith *DoublyNode

	nodeWith = DoublyLinkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		dnode.nextNode = nodeWith.nextNode
		dnode.previousNode = nodeWith
		nodeWith.nextNode = dnode
	}

}

// LastNode method of DoublyLinkedList
func (DoublyLinkedList *DoublyLinkedList) LastNode() *DoublyNode {
	var dnode *DoublyNode
	var lastNode *DoublyNode
	for dnode = DoublyLinkedList.headNode; dnode != nil; dnode = dnode.nextNode {
		if dnode.nextNode == nil {
			lastNode = dnode
		}
	}
	return lastNode
}

// AddToEnd method of DoublyLinkedList
func (DoublyLinkedList *DoublyLinkedList) AddToEnd(property int) {
	var dnode = &DoublyNode{}
	dnode.property = property
	dnode.nextNode = nil

	var lastNode *DoublyNode

	lastNode = DoublyLinkedList.LastNode()

	if lastNode != nil {

		lastNode.nextNode = dnode
		dnode.previousNode = lastNode
	}
}

// IterateList method of DoublyLinkedList
func (DoublyLinkedList *DoublyLinkedList) IterateList() {

	var dnode *DoublyNode
	for dnode = DoublyLinkedList.headNode; dnode != nil; dnode = dnode.nextNode {

		fmt.Println(dnode.property)
	}
}

// NodeBetweenValues method of DoublyLinkedList
func (DoublyLinkedList *DoublyLinkedList) NodeBetweenValues(firstProperty int, secondProperty int) *DoublyNode {
	var dnode *DoublyNode
	var nodeWith *DoublyNode
	for dnode = DoublyLinkedList.headNode; dnode != nil; dnode = dnode.nextNode {
		if dnode.previousNode != nil && dnode.nextNode != nil {
			if dnode.previousNode.property == firstProperty && dnode.nextNode.property == secondProperty {
				nodeWith = dnode
				break
			}
		}
	}
	return nodeWith
}

// DoublyLinkedListMain method
func DoublyLinkedListMain() {

	var dlist DoublyLinkedList

	dlist.AddToHead(1)
	dlist.AddToHead(3)
	dlist.AddToEnd(5)
	dlist.AddAfter(1, 7)
	fmt.Println(dlist.headNode.property)

	var found *DoublyNode
	found = dlist.NodeBetweenValues(1, 5)
	fmt.Println(found.property)

}

// main runs the demo entry points of this package
// (SyncQueueMain is excluded: it runs an endless concurrency demo)
func main() {
	DoublyLinkedListMain()
	QueueMain()
	SetMain()
	SinglyLinkedListMain()
	StackMain()
	TuplesMain()
}

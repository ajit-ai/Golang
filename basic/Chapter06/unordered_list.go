// main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import (
	"fmt"
)

// UnorderedNode class
type UnorderedNode struct {
	property int
	nextNode *UnorderedNode
}

// UnOrderedList class
type UnOrderedList struct {
	headNode *UnorderedNode
}

// AddToHead method of UnOrderedList class
func (UnOrderedList *UnOrderedList) AddToHead(property int) {
	var UnorderedNode = &UnorderedNode{}
	UnorderedNode.property = property
	UnorderedNode.nextNode = nil

	if UnOrderedList.headNode != nil {
		UnorderedNode.nextNode = UnOrderedList.headNode
	}

	UnOrderedList.headNode = UnorderedNode

}

// IterateList method iterates over UnOrderedList
func (UnOrderedList *UnOrderedList) IterateList() {

	var UnorderedNode *UnorderedNode
	for UnorderedNode = UnOrderedList.headNode; UnorderedNode != nil; UnorderedNode = UnorderedNode.nextNode {
		fmt.Println(UnorderedNode.property)

	}
}

// main method
func UnorderedListMain() {

	var unOrderedList UnOrderedList

	unOrderedList = UnOrderedList{}

	unOrderedList.AddToHead(1)
	unOrderedList.AddToHead(3)
	unOrderedList.AddToHead(5)
	unOrderedList.AddToHead(7)
	unOrderedList.IterateList()

}

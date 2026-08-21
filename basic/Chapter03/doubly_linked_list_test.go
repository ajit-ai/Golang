package main

import "testing"

func doublyValues(l *DoublyLinkedList) []int {
	var values []int
	for dnode := l.headNode; dnode != nil; dnode = dnode.nextNode {
		values = append(values, dnode.property)
	}
	return values
}

func TestDoublyLinkedListInsertions(t *testing.T) {
	var l DoublyLinkedList
	l.AddToHead(1)
	l.AddToHead(3)
	l.AddToEnd(5)
	l.AddAfter(1, 7)

	got := doublyValues(&l)
	want := []int{3, 1, 7, 5}
	if len(got) != len(want) {
		t.Fatalf("list = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("list = %v, want %v", got, want)
		}
	}
}

func TestDoublyLinkedListPreviousLinks(t *testing.T) {
	var l DoublyLinkedList
	l.AddToHead(1)
	l.AddToEnd(5)
	l.AddToEnd(9)

	last := l.LastNode()
	if last == nil || last.property != 9 {
		t.Fatalf("LastNode = %v, want 9", last)
	}
	if last.previousNode == nil || last.previousNode.property != 5 {
		t.Errorf("previous link broken: got %v, want 5", last.previousNode)
	}
	if l.headNode.previousNode != nil {
		t.Error("head node previous should be nil")
	}
}

func TestDoublyLinkedListNodeBetweenValues(t *testing.T) {
	var l DoublyLinkedList
	l.AddToHead(1)
	l.AddToEnd(5)
	l.AddToEnd(9)

	// list layout: 1 -> 5 -> 9; the node between values 1 and 9 is 5
	if n := l.NodeBetweenValues(1, 9); n == nil || n.property != 5 {
		t.Errorf("NodeBetweenValues(1, 9) = %v, want node with property 5", n)
	}
	if n := l.NodeBetweenValues(1, 5); n != nil {
		t.Errorf("NodeBetweenValues(1, 5) = %v, want nil", n)
	}
}

func TestDoublyLinkedListMainSmoke(t *testing.T) {
	DoublyLinkedListMain()
}

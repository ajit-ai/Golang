package main

import "testing"

func singlyValues(l *SinglyLinkedList) []int {
	var values []int
	for snode := l.headNode; snode != nil; snode = snode.nextNode {
		values = append(values, snode.property)
	}
	return values
}

func TestSinglyLinkedListInsertions(t *testing.T) {
	var l SinglyLinkedList
	l.AddToHead(1)
	l.AddToHead(3)
	l.AddToEnd(5)
	l.AddAfter(1, 7)

	got := singlyValues(&l)
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

func TestSinglyLinkedListNodeWithValue(t *testing.T) {
	var l SinglyLinkedList
	l.AddToHead(4)
	l.AddToHead(8)
	if n := l.NodeWithValue(4); n == nil || n.property != 4 {
		t.Errorf("NodeWithValue(4) = %v, want node 4", n)
	}
	if n := l.NodeWithValue(99); n != nil {
		t.Errorf("NodeWithValue(99) = %v, want nil", n)
	}
}

func TestSinglyLinkedListLastNodeEmpty(t *testing.T) {
	var l SinglyLinkedList
	if l.LastNode() != nil {
		t.Error("LastNode of empty list should be nil")
	}
}

func TestSinglyLinkedListMainSmoke(t *testing.T) {
	SinglyLinkedListMain()
}

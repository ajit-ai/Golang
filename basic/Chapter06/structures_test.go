package main

import (
	"sort"
	"testing"
)

func TestCircularQueueFIFO(t *testing.T) {
	q := NewQueue(3)
	if !q.IsUnUsed() {
		t.Error("fresh queue should be unused")
	}
	q.Add(1)
	q.Add(2)
	q.Add(3)
	if q.IsUnUsed() {
		t.Error("queue with elements should not be unused")
	}
	for _, want := range []int{1, 2, 3} {
		got := q.MoveOneStep()
		if got != want {
			t.Errorf("MoveOneStep = %v, want %v", got, want)
		}
	}
	if got := q.MoveOneStep(); got != nil {
		t.Errorf("MoveOneStep on drained queue = %v, want nil", got)
	}
}

func TestCircularQueueWrapsAround(t *testing.T) {
	q := NewQueue(2)
	q.Add(1)
	q.Add(2)
	if got := q.MoveOneStep(); got != 1 {
		t.Errorf("first MoveOneStep = %v, want 1", got)
	}
	q.Add(3)
	if got := q.MoveOneStep(); got != 2 {
		t.Errorf("second MoveOneStep = %v, want 2", got)
	}
	if got := q.MoveOneStep(); got != 3 {
		t.Errorf("third MoveOneStep = %v, want 3", got)
	}
}

func TestCircularQueueAddWhenFullPanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Add on full queue should panic")
		}
	}()
	q := NewQueue(1)
	q.Add(1)
	q.Add(2)
}

func TestReverseLinkedList(t *testing.T) {
	head := CreateLinkedList()
	reversed := ReverseLinkedList(head)

	count := 0
	curr := reversed
	for curr != nil {
		count++
		if curr.nextNode == nil && curr.property != 'a' {
			t.Errorf("last node of reversed list = %c, want 'a'", curr.property)
		}
		curr = curr.nextNode
	}
	if count != 26 {
		t.Errorf("reversed list length = %d, want 26", count)
	}
	if reversed.property != 'z' {
		t.Errorf("head of reversed list = %c, want 'z'", reversed.property)
	}
}

func TestSortByFactor(t *testing.T) {
	things := []Thing{
		{name: "a", mass: 3.0},
		{name: "b", mass: 1.0},
		{name: "c", mass: 2.0},
	}
	ByFactor(func(t1, t2 *Thing) bool { return t1.mass < t2.mass }).Sort(things)
	if things[0].name != "b" || things[1].name != "c" || things[2].name != "a" {
		t.Errorf("sorted by mass = %v", things)
	}
}

func TestSortMultiKeysOrderedBy(t *testing.T) {
	commits := []Commit{
		{username: "bob", numlines: 3},
		{username: "alice", numlines: 1},
		{username: "alice", numlines: 5},
	}
	OrderedBy(
		func(p1, p2 *Commit) bool { return p1.username < p2.username },
		func(p1, p2 *Commit) bool { return p1.numlines < p2.numlines },
	).Sort(commits)

	want := []Commit{
		{username: "alice", numlines: 1},
		{username: "alice", numlines: 5},
		{username: "bob", numlines: 3},
	}
	for i := range want {
		if commits[i] != want[i] {
			t.Fatalf("multi-key sort result[%d] = %+v, want %+v", i, commits[i], want[i])
		}
	}
}

func TestSortByAge(t *testing.T) {
	employees := SortByAge{{Name: "old", Age: 50}, {Name: "young", Age: 20}}
	sort.Sort(employees)
	if employees[0].Name != "young" {
		t.Errorf("SortByAge first = %s, want young", employees[0].Name)
	}
}

func TestUnOrderedListAddToHead(t *testing.T) {
	var list UnOrderedList
	list.AddToHead(1)
	list.AddToHead(3)
	if list.headNode == nil || list.headNode.property != 3 {
		t.Error("AddToHead should place newest element at head")
	}
	if list.headNode.nextNode == nil || list.headNode.nextNode.property != 1 {
		t.Error("second node should hold first added value")
	}
}

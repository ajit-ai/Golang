package main

import (
	"container/heap"
	"reflect"
	"testing"
)

func TestIntegerHeapLen(t *testing.T) {
	iheap := IntegerHeap{3, 1, 2}
	if got := iheap.Len(); got != 3 {
		t.Errorf("Len() = %d, want 3", got)
	}
}

func TestIntegerHeapLess(t *testing.T) {
	iheap := IntegerHeap{1, 2}
	if !iheap.Less(0, 1) {
		t.Error("Less(0, 1) = false, want true for min-heap ordering")
	}
	if iheap.Less(1, 0) {
		t.Error("Less(1, 0) = true, want false")
	}
}

func TestIntegerHeapSwap(t *testing.T) {
	iheap := IntegerHeap{1, 2}
	iheap.Swap(0, 1)
	if iheap[0] != 2 || iheap[1] != 1 {
		t.Errorf("after Swap(0, 1) = %v, want [2 1]", iheap)
	}
}

func TestIntegerHeapPushPop(t *testing.T) {
	intHeap := &IntegerHeap{1, 4, 5}

	heap.Init(intHeap)
	heap.Push(intHeap, 2)

	if got := (*intHeap)[0]; got != 1 {
		t.Errorf("minimum = %d, want 1", got)
	}

	var got []int
	for intHeap.Len() > 0 {
		got = append(got, heap.Pop(intHeap).(int))
	}
	want := []int{1, 2, 4, 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("pop order = %v, want %v", got, want)
	}
}

func TestHeapMain(t *testing.T) {
	HeapMain()
}

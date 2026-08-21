package main

import "testing"

func TestStackPushPopLIFO(t *testing.T) {
	var stack Stack
	stack.New()

	first := &Element{}
	second := &Element{}
	stack.Push(first)
	stack.Push(second)

	if stack.Pop() != second {
		t.Error("Pop should return last pushed element (LIFO)")
	}
	if stack.Pop() != first {
		t.Error("Pop should return first pushed element afterwards")
	}
	if stack.Pop() != nil {
		t.Error("Pop on empty stack should return nil")
	}
}

func TestStackMainSmoke(t *testing.T) {
	StackMain()
}

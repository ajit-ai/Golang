// Package gostack provides a small generic LIFO stack.
//
// It is the repository's first importable library package: unlike every
// other directory here (which are `package main` demo programs), this
// package is meant to be imported, e.g.
//
//	import "github.com/ajit-ai/Golang/pkg/gostack"
package gostack

// Stack is a generic last-in-first-out collection.
// The zero value is an empty stack ready to use.
type Stack[T any] struct {
	elements []T
}

// New returns a stack with capacity pre-allocated for n elements
func New[T any](n int) *Stack[T] {
	return &Stack[T]{elements: make([]T, 0, n)}
}

// Push adds value to the top of the stack
func (s *Stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

// Pop removes and returns the top element. ok is false when empty.
func (s *Stack[T]) Pop() (value T, ok bool) {
	if len(s.elements) == 0 {
		return value, false
	}
	value = s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return value, true
}

// Peek returns the top element without removing it. ok is false when empty.
func (s *Stack[T]) Peek() (value T, ok bool) {
	if len(s.elements) == 0 {
		return value, false
	}
	return s.elements[len(s.elements)-1], true
}

// Len returns the number of elements on the stack
func (s *Stack[T]) Len() int {
	return len(s.elements)
}

// Drain pops everything and returns the elements in pop order (top first)
func (s *Stack[T]) Drain() []T {
	out := make([]T, 0, len(s.elements))
	for {
		v, ok := s.Pop()
		if !ok {
			return out
		}
		out = append(out, v)
	}
}

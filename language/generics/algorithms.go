// main package demonstrates Go generics (type parameters, constraints)
package main

import (
	"cmp"
	"fmt"
)

// Number constrains type parameters to numeric types
type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

// GenericStack is a LIFO stack that works for any element type
type GenericStack[T any] struct {
	elements []T
}

// Push adds an element on top of the stack
func (s *GenericStack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

// Pop removes and returns the top element; ok is false when empty
func (s *GenericStack[T]) Pop() (value T, ok bool) {
	if len(s.elements) == 0 {
		return value, false
	}
	value = s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return value, true
}

// Len returns the number of elements in the stack
func (s *GenericStack[T]) Len() int {
	return len(s.elements)
}

// MapSlice applies f to every element and returns a new slice
func MapSlice[T, R any](input []T, f func(T) R) []R {
	output := make([]R, 0, len(input))
	for _, v := range input {
		output = append(output, f(v))
	}
	return output
}

// Filter keeps only the elements for which predicate returns true
func Filter[T any](input []T, predicate func(T) bool) []T {
	var output []T
	for _, v := range input {
		if predicate(v) {
			output = append(output, v)
		}
	}
	return output
}

// Min returns the smallest of two ordered values
func Min[T cmp.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max returns the largest of two ordered values
func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Sum adds all numbers in the slice
func Sum[T Number](values []T) T {
	var total T
	for _, v := range values {
		total += v
	}
	return total
}

// Keys returns the keys of a map in unspecified order
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// GenericStackMain demonstrates the generic stack with different types
func GenericStackMain() {
	ints := &GenericStack[int]{}
	ints.Push(1)
	ints.Push(2)
	if v, ok := ints.Pop(); ok {
		fmt.Println("popped int:", v)
	}

	strs := &GenericStack[string]{}
	strs.Push("go")
	strs.Push("generics")
	fmt.Println("stack length:", strs.Len())
}

// GenericAlgorithmsMain demonstrates generic helper functions
func GenericAlgorithmsMain() {
	fmt.Println("map x2:", MapSlice([]int{1, 2, 3}, func(v int) int { return v * 2 }))
	fmt.Println("filter even:", Filter([]int{1, 2, 3, 4}, func(v int) bool { return v%2 == 0 }))
	fmt.Println("min:", Min(7, 3), "max:", Max(7, 3))
	fmt.Println("sum floats:", Sum([]float64{1.5, 2.5}))
	fmt.Println("keys:", Keys(map[string]int{"a": 1}))
}

// main runs the demo entry points of this package
func main() {
	GenericStackMain()
	GenericAlgorithmsMain()
}

// Command gostack-demo consumes the github.com/ajit-ai/Golang/pkg/gostack
// library package, demonstrating module imports and exported identifiers.
package main

import (
	"fmt"

	"github.com/ajit-ai/Golang/pkg/gostack"
)

func main() {
	s := gostack.New[int](3)
	for _, v := range []int{1, 2, 3} {
		s.Push(v)
	}

	top, _ := s.Peek()
	fmt.Println("top:", top)
	fmt.Println("pop order:", s.Drain())
	fmt.Println("length now:", s.Len())
}

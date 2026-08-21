package main

import (
	"fmt"
)

func OffsetArray() [10]int {
	var m [10]int

	for k := 0; k < 10; k++ {
		m[k] = k + 200
	}

	return m
}

// ComplexityMain method
func ComplexityMain() {
	m := OffsetArray()

	for k := 0; k < 10; k++ {
		fmt.Printf("Element[%d] = %d\n", k, m[k])
	}
}

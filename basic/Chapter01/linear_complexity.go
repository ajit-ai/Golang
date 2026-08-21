// main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

func ScaledArray() [10]int {
	var m [10]int

	for k := 0; k < 10; k++ {
		m[k] = k * 200
	}

	return m
}

// LinearComplexityMain method
func LinearComplexityMain() {
	m := ScaledArray()

	for k := 0; k < 10; k++ {
		fmt.Printf("Element[%d] = %d\n", k, m[k])
	}
}

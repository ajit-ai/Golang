// main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

func TableRow(k int) []int {
	row := make([]int, 10)

	for l := 1; l <= 10; l++ {
		row[l-1] = l * k
	}

	return row
}

// QuadraticComplexityMain method
func QuadraticComplexityMain() {

	for k := 1; k <= 10; k++ {
		fmt.Println(" Multiplication Table", k)
		for _, x := range TableRow(k) {
			fmt.Println(x)
		}

	}
}

// /main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// main method
func NullMatrixMain() {

	var matrix = [3][3]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0}}

	fmt.Println(matrix)
}

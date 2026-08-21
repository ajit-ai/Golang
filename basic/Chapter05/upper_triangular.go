// /main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// main method
func UpperTriangularMain() {

	var matrix = [3][3]int{
		{1, 2, 3},
		{0, 1, 4},
		{0, 0, 1}}

	fmt.Println(matrix)
}

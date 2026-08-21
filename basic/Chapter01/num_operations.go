// main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// Add method given a and b integers
func Add(a int, b int) int {
	return a + b
}

// NumOperationsMain method
func NumOperationsMain() {

	//fmt.Println("My favorite number is", rand.Intn(10))

	//fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	//fmt.Println(math.Pi)

	fmt.Println(Add(42, 13))

}

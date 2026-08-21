// main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// Fibonacci method given k integer
func Fibonacci(k int) int {

	if k <= 1 {
		return 1
	}
	return Fibonacci(k-1) + Fibonacci(k-2)

}

// DivideMain method
func DivideMain() {

	var m int = 5

	for m = 0; m < 8; m++ {

		var fib = Fibonacci(m)
		fmt.Println(fib)
	}

}

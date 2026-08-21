// /main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// increment method
func addOne(num int) int {

	num++
	fmt.Println("added to num", num, "Address of num", &num)
	return num
}

// StackMemoryAllocationMain method
func StackMemoryAllocationMain() {

	var number int
	number = 17

	fmt.Println("value of number", number, "Address of number", &number)

	addOne(number)

	fmt.Println("value of number after adding One", number, "Address of", &number)
}

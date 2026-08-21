// /main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// increment method
func addOnePointer(num *int) {

	*num++
	fmt.Println("added to num", num, "Address of num", &num, "Value Points To", *num)
}

// StackMemoryPointerMain method
func StackMemoryPointerMain() {

	var number int
	number = 17

	fmt.Println("value of number", number, "Address of number", &number)

	addOnePointer(&number)

	fmt.Println("value of number after adding One", number, "Address of", &number)
}

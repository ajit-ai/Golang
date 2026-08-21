// main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// hash method computes a polynomial rolling hash of the string
func stringHash(str string) int {

	var constant int

	constant = 42

	var hashCode int
	for i := 0; i < len(str); i++ {
		hashCode = hashCode*constant + int(str[i])
	}

	return hashCode

}

// main method
func HashStringMain() {

	var str string

	str = "checkforhash"

	var hashCode int

	hashCode = stringHash(str)

	fmt.Println(hashCode)

}

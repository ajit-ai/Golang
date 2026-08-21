// main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import (
	"fmt"
)

// main method
func TwoDArrayMain() {
	var TwoDArray [8][8]int

	TwoDArray[3][6] = 18

	TwoDArray[7][4] = 3

	fmt.Println(TwoDArray)

}

// main runs the demo entry points of this package
// (CrmAppMain, DatabaseOperationsMain and WebFormsMain are excluded:
// they need a live MySQL server or block serving HTTP on :8000)
func main() {
	AppendSliceMain()
	BasicSliceMain()
	ArraysMain()
	MapsMain()
	SlicesMain()
	TwodArrayMain()
	TwodSlicesMain()
	TwoDArrayMain()
}

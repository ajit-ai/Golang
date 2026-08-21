package main

import "fmt"

func ValueTypesMain() {
	// Variable declaration with initial value
	var a = "initial"
	fmt.Println(a)

	// Multiple variable declaration
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Type inference
	var d = true
	fmt.Println(d)

	// Zero value
	var e int
	fmt.Println(e)

	// Short variable declaration
	f := "short"
	fmt.Println(f)
}

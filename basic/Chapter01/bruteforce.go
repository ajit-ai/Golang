package main

import (
	"fmt"
)

// FindElement method given array and k element
func FindElement(arr [10]int, k int) bool {
	var i int
	for i = 0; i < 10; i++ {

		if arr[i] == k {
			return true
		}
	}
	return false
}

// BruteforceMain method
func BruteforceMain() {

	var arr = [10]int{1, 4, 7, 8, 3, 9, 2, 4, 1, 8}

	var check bool = FindElement(arr, 10)

	fmt.Println(check)

	var check2 bool = FindElement(arr, 9)

	fmt.Println(check2)

}

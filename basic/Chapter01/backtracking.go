package main

import (
	"fmt"
)

// FindElementsWithSum  of k from arr of size
func FindElementsWithSum(arr [10]int, combinations [19]int, size int, k int, addValue int, l int, m int) int {

	var num int = 0

	if addValue > k {
		return 0
	}

	if addValue == k {
		num = num + 1
		var p int = 0
		for p = 0; p < m; p++ {

			fmt.Printf("%d,", arr[combinations[p]])
		}
		fmt.Println(" ")
	}

	if m >= len(combinations) {
		return num
	}

	var i int
	for i = l; i < size; i++ {

		//fmt.Println(" m", m)
		combinations[m] = l

		num = num + FindElementsWithSum(arr, combinations, size, k, addValue+arr[i], l, m+1)
		l = l + 1
	}
	return num
}

// BacktrackingMain method
func BacktrackingMain() {

	var arr = [10]int{1, 4, 7, 8, 3, 9, 2, 4, 1, 8}

	var addedSum int = 18

	var combinations [19]int

	FindElementsWithSum(arr, combinations, 10, addedSum, 0, 0, 0)

	//fmt.Println(check)

	//var check2 bool = findElement(arr,9)

	//fmt.Println(check2)

}

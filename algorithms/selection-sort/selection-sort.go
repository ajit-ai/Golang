package main

import (
	"fmt"
)

// 1. Selection sort for []int (ascending)
func SelectionSortInt(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

// 2. Selection sort for []float64 (ascending)
func SelectionSortFloat(arr []float64) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

// 3. Selection sort for []string (ascending)
func SelectionSortString(arr []string) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

// 4. Selection sort for []int (descending)
func SelectionSortIntDesc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
	}
}

// 5. Selection sort for slice of struct by field
type Person struct {
	Name string
	Age  int
}

func SelectionSortPersonByAge(arr []Person) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j].Age < arr[minIdx].Age {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

// --- Example usage in main ---

func main() {
	ints := []int{5, 3, 8, 1}
	SelectionSortInt(ints)
	fmt.Println("Sorted ints:", ints)

	floats := []float64{2.2, 3.3, 1.1}
	SelectionSortFloat(floats)
	fmt.Println("Sorted floats:", floats)

	strs := []string{"z", "a", "m"}
	SelectionSortString(strs)
	fmt.Println("Sorted strings:", strs)

	intsDesc := []int{4, 2, 7, 1}
	SelectionSortIntDesc(intsDesc)
	fmt.Println("Sorted ints desc:", intsDesc)

	people := []Person{
		{"Ajit", 40},
		{"Dinesh", 25},
		{"Ramu", 35},
	}
	SelectionSortPersonByAge(people)
	fmt.Println("Sorted people by age:", people)
}

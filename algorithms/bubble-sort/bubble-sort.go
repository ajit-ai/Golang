package main

import (
	"fmt"
)

// 1. Bubble sort for []int (ascending)
func BubbleSortInt(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 2. Bubble sort for []float64 (ascending)
func BubbleSortFloat(arr []float64) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func BubbleSortPersonByAge(arr []Person) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j].Age > arr[j+1].Age {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 3. Bubble sort for []string (ascending)
func BubbleSortString(arr []string) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 4. Bubble sort for []int (descending)
func BubbleSortIntDesc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 5. Bubble sort for slice of struct by field
type Person struct {
	Name string
	Age  int
}

// --- Testing code ---

// --- Example usage in main ---

func main() {
	ints := []int{5, 3, 8, 1}
	BubbleSortInt(ints)
	fmt.Println("Sorted ints:", ints)

	floats := []float64{2.2, 3.3, 1.1}
	BubbleSortFloat(floats)
	fmt.Println("Sorted floats:", floats)

	strs := []string{"z", "a", "m"}
	BubbleSortString(strs)
	fmt.Println("Sorted strings:", strs)

	intsDesc := []int{4, 2, 7, 1}
	BubbleSortIntDesc(intsDesc)
	fmt.Println("Sorted ints desc:", intsDesc)

	people := []Person{
		{"Ajit", 40},
		{"dinesh", 25},
		{"Ramu", 35},
	}
	BubbleSortPersonByAge(people)
	fmt.Println("Sorted people by age:", people)
}

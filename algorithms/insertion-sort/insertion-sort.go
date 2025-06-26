package main

import (
	"fmt"
)

// 1. Insertion sort for []int (ascending)
func InsertionSortInt(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// 2. Insertion sort for []int (descending)
func InsertionSortIntDesc(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] < key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// 3. Insertion sort for []float64 (ascending)
func InsertionSortFloat(arr []float64) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// 4. Insertion sort for []string (ascending)
func InsertionSortString(arr []string) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// 5. Insertion sort for []string (descending)
func InsertionSortStringDesc(arr []string) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] < key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// 6. Insertion sort for slice of struct by int field
type Person struct {
	Name string
	Age  int
}

func InsertionSortPersonByAge(arr []Person) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j].Age > key.Age {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// 7. Insertion sort for slice of struct by string field
func InsertionSortPersonByName(arr []Person) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j].Name > key.Name {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// 8. Insertion sort for []rune (ascending)
func InsertionSortRune(arr []rune) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// 9. Insertion sort for []byte (ascending)
func InsertionSortByte(arr []byte) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// 10. Insertion sort for []float64 (descending)
func InsertionSortFloatDesc(arr []float64) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] < key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// --- Example usage in main ---

func main() {
	ints := []int{5, 3, 8, 1}
	InsertionSortInt(ints)
	fmt.Println("Sorted ints:", ints)

	intsDesc := []int{4, 2, 7, 1}
	InsertionSortIntDesc(intsDesc)
	fmt.Println("Sorted ints desc:", intsDesc)

	floats := []float64{2.2, 3.3, 1.1}
	InsertionSortFloat(floats)
	fmt.Println("Sorted floats:", floats)

	floatsDesc := []float64{2.2, 3.3, 1.1}
	InsertionSortFloatDesc(floatsDesc)
	fmt.Println("Sorted floats desc:", floatsDesc)

	strs := []string{"z", "a", "m"}
	InsertionSortString(strs)
	fmt.Println("Sorted strings:", strs)

	strsDesc := []string{"z", "a", "m"}
	InsertionSortStringDesc(strsDesc)
	fmt.Println("Sorted strings desc:", strsDesc)

	people := []Person{
		{"Dave", 40},
		{"Carol", 25},
		{"Eve", 35},
	}
	InsertionSortPersonByAge(people)
	fmt.Println("Sorted people by age:", people)

	people2 := []Person{
		{"Dave", 40},
		{"Carol", 25},
		{"Eve", 35},
	}
	InsertionSortPersonByName(people2)
	fmt.Println("Sorted people by name:", people2)

	runes := []rune{'z', 'a', 'm'}
	InsertionSortRune(runes)
	fmt.Println("Sorted runes:", string(runes))

	bytes := []byte{'z', 'a', 'm'}
	InsertionSortByte(bytes)
	fmt.Println("Sorted bytes:", string(bytes))
}

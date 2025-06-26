package main

import (
	"fmt"
	"strings"
)

// 1. Merge sort for []int (ascending)
func MergeSortInt(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSortInt(arr[:mid])
	right := MergeSortInt(arr[mid:])
	return mergeInt(left, right)
}
func mergeInt(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	return append(result, left[i:]...)
}

// 2. Merge sort for []int (descending)
func MergeSortIntDesc(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSortIntDesc(arr[:mid])
	right := MergeSortIntDesc(arr[mid:])
	return mergeIntDesc(left, right)
}
func mergeIntDesc(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] > right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	return append(result, left[i:]...)
}

// 3. Merge sort for []float64 (ascending)
func MergeSortFloat(arr []float64) []float64 {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSortFloat(arr[:mid])
	right := MergeSortFloat(arr[mid:])
	return mergeFloat(left, right)
}
func mergeFloat(left, right []float64) []float64 {
	result := make([]float64, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	return append(result, left[i:]...)
}

// 4. Merge sort for []string (ascending)
func MergeSortString(arr []string) []string {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSortString(arr[:mid])
	right := MergeSortString(arr[mid:])
	return mergeString(left, right)
}
func mergeString(left, right []string) []string {
	result := make([]string, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	return append(result, left[i:]...)
}

// 5. Merge sort for []string (case-insensitive)
func MergeSortStringCI(arr []string) []string {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSortStringCI(arr[:mid])
	right := MergeSortStringCI(arr[mid:])
	return mergeStringCI(left, right)
}
func mergeStringCI(left, right []string) []string {
	result := make([]string, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if strings.ToLower(left[i]) < strings.ToLower(right[j]) {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	return append(result, left[i:]...)
}

// 6. Merge sort for []rune (ascending)
func MergeSortRune(arr []rune) []rune {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSortRune(arr[:mid])
	right := MergeSortRune(arr[mid:])
	return mergeRune(left, right)
}
func mergeRune(left, right []rune) []rune {
	result := make([]rune, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// 7. Merge sort for []byte (ascending)
func MergeSortByte(arr []byte) []byte {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSortByte(arr[:mid])
	right := MergeSortByte(arr[mid:])
	return mergeByte(left, right)
}
func mergeByte(left, right []byte) []byte {
	result := make([]byte, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	//return append(result, left[i:]...)

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// 8. Merge sort for slice of struct by int field
type Person struct {
	Name string
	Age  int
}

func MergeSortPersonByAge(arr []Person) []Person {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSortPersonByAge(arr[:mid])
	right := MergeSortPersonByAge(arr[mid:])
	return mergePersonByAge(left, right)
}
func mergePersonByAge(left, right []Person) []Person {
	result := make([]Person, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i].Age < right[j].Age {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	//return append(result, left[i:]...)
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// 9. Merge sort for slice of struct by string field
func MergeSortPersonByName(arr []Person) []Person {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSortPersonByName(arr[:mid])
	right := MergeSortPersonByName(arr[mid:])
	return mergePersonByName(left, right)
}
func mergePersonByName(left, right []Person) []Person {
	result := make([]Person, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i].Name < right[j].Name {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	//return append(result, left[i:]...)
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// 10. Merge sort for custom comparator (descending float64)
func MergeSortFloatDesc(arr []float64) []float64 {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSortFloatDesc(arr[:mid])
	right := MergeSortFloatDesc(arr[mid:])
	return mergeFloatDesc(left, right)
}
func mergeFloatDesc(left, right []float64) []float64 {
	result := make([]float64, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] > right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	//return append(result, left[i:]...)
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// --- Example usage in main ---

func main() {
	ints := []int{5, 3, 8, 1}
	fmt.Println("Sorted ints:", MergeSortInt(ints))

	intsDesc := []int{4, 2, 7, 1}
	fmt.Println("Sorted ints desc:", MergeSortIntDesc(intsDesc))

	floats := []float64{2.2, 3.3, 1.1}
	fmt.Println("Sorted floats:", MergeSortFloat(floats))

	floatsDesc := []float64{2.2, 3.3, 1.1}
	fmt.Println("Sorted floats desc:", MergeSortFloatDesc(floatsDesc))

	strs := []string{"z", "a", "m"}
	fmt.Println("Sorted strings:", MergeSortString(strs))

	strsCI := []string{"Zebra", "apple", "Monkey"}
	fmt.Println("Sorted strings (case-insensitive):", MergeSortStringCI(strsCI))

	people := []Person{
		{"Dave", 40},
		{"Carol", 25},
		{"Eve", 35},
	}
	fmt.Println("Sorted people by age:", MergeSortPersonByAge(people))

	people2 := []Person{
		{"Dave", 40},
		{"Carol", 25},
		{"Eve", 35},
	}
	fmt.Println("Sorted people by name:", MergeSortPersonByName(people2))

	runes := []rune{'z', 'a', 'm'}
	fmt.Println("Sorted runes:", string(MergeSortRune(runes)))

	bytes := []byte{'z', 'a', 'm'}
	fmt.Println("Sorted bytes:", string(MergeSortByte(bytes)))
}

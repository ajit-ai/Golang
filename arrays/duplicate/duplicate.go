package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	fmt.Println("Hello, Ajit")

	// Example of finding duplicates in an array
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3}
	duplicates := findDuplicates(arr)
	fmt.Println("1. Using a map=>", duplicates)

	// Example of finding duplicates using a set
	duplicatesSet := findDuplicatestwo(arr)
	fmt.Println("2. Using a set=>", duplicatesSet)

	// Example of finding duplicates using a hash table
	duplicatesHash := findDuplicatesthree(arr)
	fmt.Println("3. Using a hash table=>", duplicatesHash)

	// Example of finding duplicates using a slice of booleans
	duplicatesBool := findDuplicatesfour(arr)
	fmt.Println("4. Using a slice of booleans=>", duplicatesBool)

	// Example of finding duplicates using a sorting algorithm
	duplicatesSort := findDuplicatessix(arr)
	fmt.Println("5. Using a sorting algorithm=>", duplicatesSort)

	// Example of finding duplicates using a binary search
	duplicatesBinary := findDuplicatessix(arr)
	fmt.Println("6. Using a binary search=>", duplicatesBinary)

	// Example of finding duplicates using a recursive function
	duplicatesRecursive := findDuplicatesseven(arr)
	fmt.Println("7. Using a recursive function=>", duplicatesRecursive)

	// Example of finding duplicates using a closure
	duplicatesClosure := findDuplicateseight(arr)
	fmt.Println("8. Using a closure=>", duplicatesClosure)

	// Example of finding duplicates using a goroutine
	duplicatesGoroutine := findDuplicatesnine(arr)
	fmt.Println("9. Using a goroutine=>", duplicatesGoroutine)

	// Example of finding duplicates using a WaitGroup
	duplicatesWaitGroup := findDuplicatesten(arr)
	fmt.Println("10. Using a WaitGroup=>", duplicatesWaitGroup)

	// Example of finding duplicates using a mutex
	duplicatesMutex := findDuplicateseleven(arr)
	fmt.Println("11. Using a mutex=>", duplicatesMutex)

	// Example of finding duplicates using a channel
	duplicatesChannel := findDuplicatesnine(arr)
	fmt.Println("12. Using a channel=>", duplicatesChannel)

}

// 1. Using a map
func findDuplicates(arr []int) []int {
	m := make(map[int]bool)
	var duplicates []int
	for _, num := range arr {
		if m[num] {
			duplicates = append(duplicates, num)
		}
		m[num] = true
	}
	return duplicates
}

//2. Using a set

func findDuplicatestwo(arr []int) []int {
	s := make(map[int]struct{})
	var duplicates []int
	for _, num := range arr {
		if _, ok := s[num]; ok {
			duplicates = append(duplicates, num)
		}
		s[num] = struct{}{}
	}
	return duplicates
}

//3. Using a hash table

func findDuplicatesthree(arr []int) []int {
	h := make(map[int]int)
	var duplicates []int
	for _, num := range arr {
		h[num]++
		if h[num] > 1 {
			duplicates = append(duplicates, num)
		}
	}
	return duplicates
}

//4. Using a slice of booleans

func findDuplicatesfour(arr []int) []int {
	b := make([]bool, len(arr))
	var duplicates []int
	for _, num := range arr {
		if b[num] {
			duplicates = append(duplicates, num)
		}
		b[num] = true
	}
	return duplicates
}

//5. Using a sorting algorithm

func findDuplicatesfive(arr []int) []int {
	sort.Ints(arr)
	var duplicates []int
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			duplicates = append(duplicates, arr[i])
		}
	}
	return duplicates
}

//6. Using a binary search

func findDuplicatessix(arr []int) []int {
	sort.Ints(arr)
	var duplicates []int
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			duplicates = append(duplicates, arr[i])
		}
	}
	return duplicates
}

//7. Using a recursive function

func findDuplicatesseven(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	if contains(arr[1:], arr[0]) {
		return append(findDuplicates(arr[1:]), arr[0])
	}
	return findDuplicates(arr[1:])
}

func contains(arr []int, num int) bool {
	for _, n := range arr {
		if n == num {
			return true
		}
	}
	return false
}

//8. Using a closure

func findDuplicateseight(arr []int) []int {
	seen := make(map[int]bool)
	var duplicates []int
	f := func(num int) {
		if seen[num] {
			duplicates = append(duplicates, num)
		}
		seen[num] = true
	}
	for _, num := range arr {
		f(num)
	}
	return duplicates
}

//9. Using a goroutine

func findDuplicatesnine(arr []int) []int {
	ch := make(chan int)
	var duplicates []int
	go func() {
		seen := make(map[int]bool)
		for _, num := range arr {
			if seen[num] {
				ch <- num
			}
			seen[num] = true
		}
		close(ch)
	}()
	for num := range ch {
		duplicates = append(duplicates, num)
	}
	return duplicates
}

//10. Using a WaitGroup

func findDuplicatesten(arr []int) []int {
	var wg sync.WaitGroup
	ch := make(chan int)
	var duplicates []int
	wg.Add(1)
	go func() {
		defer wg.Done()
		seen := make(map[int]bool)
		for _, num := range arr {
			if seen[num] {
				ch <- num
			}
			seen[num] = true
		}
		close(ch)
	}()
	go func() {
		for num := range ch {
			duplicates = append(duplicates, num)
		}
		wg.Done()
	}()
	wg.Wait()
	return duplicates
}

//11. Using a mutex

func findDuplicateseleven(arr []int) []int {
	mu := &sync.Mutex{}
	var duplicates []int
	seen := make(map[int]bool)
	for _, num := range arr {
		mu.Lock()
		if seen[num] {
			duplicates = append(duplicates, num)
		}
		seen[num] = true
	}
	return duplicates
}

package main

import (
	"fmt"
	"strings"
)

func ForUseMain() {
	// 1. Basic for loop
	for i := 0; i < 3; i++ {
		fmt.Println("Basic:", i)
	}

	// 2. While-like loop
	i := 0
	for i < 3 {
		fmt.Println("While-like:", i)
		i++
	}

	// 3. Infinite loop with break
	count := 0
	for {
		if count == 2 {
			break
		}
		fmt.Println("Infinite with break:", count)
		count++
	}

	// 4. Loop with continue
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Continue:", i)
	}

	// 5. Loop over slice
	s := []string{"a", "b", "c"}
	for idx, val := range s {
		fmt.Println("Slice:", idx, val)
	}

	// 6. Loop over array
	arr := [3]int{1, 2, 3}
	for i, v := range arr {
		fmt.Println("Array:", i, v)
	}

	// 7. Loop over map
	m := map[string]int{"x": 1, "y": 2}
	for k, v := range m {
		fmt.Println("Map:", k, v)
	}

	// 8. Loop over string (runes)
	str := "go!"
	for i, c := range str {
		fmt.Printf("String: %d %c\n", i, c)
	}

	// 9. Nested loops
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println("Nested:", i, j)
		}
	}

	// 10. Loop with label and break
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break outer
			}
			fmt.Println("Label break:", i, j)
		}
	}

	// 11. Loop with label and continue
outer2:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				continue outer2
			}
			fmt.Println("Label continue:", i, j)
		}
	}

	// 12. Loop with function call in condition
	for f := func() int { return 2 }; f() > 1; f = func() int { return 0 } {
		fmt.Println("Func in condition")
		break
	}

	// 13. Loop with decrement
	for i := 3; i > 0; i-- {
		fmt.Println("Decrement:", i)
	}

	// 14. Loop with multiple variables
	for i, j := 0, 10; i < 3 && j > 7; i, j = i+1, j-1 {
		fmt.Println("Multiple vars:", i, j)
	}

	// 15. Loop with no body
	for i := 0; i < 3; i++ {
	}
	fmt.Println("No body loop done")

	// 16. Loop to sum slice
	sum := 0
	for _, v := range []int{1, 2, 3} {
		sum += v
	}
	fmt.Println("Sum:", sum)

	// 17. Loop to find value in slice
	found := false
	for _, v := range []int{1, 2, 3} {
		if v == 2 {
			found = true
			break
		}
	}
	fmt.Println("Found 2:", found)

	// 18. Loop to reverse slice
	a := []int{1, 2, 3}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println("Reversed:", a)

	// 19. Loop to count runes in string
	countRunes := 0
	for range "hello" {
		countRunes++
	}
	fmt.Println("Rune count:", countRunes)

	// 20. Loop to build string
	var builder strings.Builder
	for i := 0; i < 5; i++ {
		builder.WriteString("x")
	}
	fmt.Println("Built string:", builder.String())

	// 21. Loop to filter slice
	nums := []int{1, 2, 3, 4, 5}
	var odds []int
	for _, n := range nums {
		if n%2 == 1 {
			odds = append(odds, n)
		}
	}
	fmt.Println("Odds:", odds)

	// 22. Loop to copy slice
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	for i, v := range src {
		dst[i] = v
	}
	fmt.Println("Copied:", dst)

	// 23. Loop to find min in slice
	min := nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	fmt.Println("Min:", min)

	// 24. Loop to find max in slice
	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	fmt.Println("Max:", max)

	// 25. Loop to count map keys
	mapCount := 0
	for range m {
		mapCount++
	}
	fmt.Println("Map key count:", mapCount)

	// 26. Loop to print diagonal of 2D array
	matrix := [2][2]int{{1, 2}, {3, 4}}
	for i := 0; i < 2; i++ {
		fmt.Println("Diagonal:", matrix[i][i])
	}

	// 27. Loop to flatten 2D array
	var flat []int
	for _, row := range matrix {
		for _, v := range row {
			flat = append(flat, v)
		}
	}
	fmt.Println("Flattened:", flat)

	// 28. Loop to check palindrome
	word := "level"
	isPalindrome := true
	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-1-i] {
			isPalindrome = false
			break
		}
	}
	fmt.Println("Is palindrome:", isPalindrome)

	// 29. Loop to count vowels
	vowels := "aeiou"
	vowelCount := 0
	for _, c := range word {
		if strings.ContainsRune(vowels, c) {
			vowelCount++
		}
	}
	fmt.Println("Vowel count:", vowelCount)

	// 30. Loop to print Fibonacci numbers
	f0, f1 := 0, 1
	for i := 0; i < 5; i++ {
		fmt.Print(f0, " ")
		f0, f1 = f1, f0+f1
	}
	fmt.Println()

	// 31. Loop to print multiplication table
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d*%d=%d ", i, j, i*j)
		}
		fmt.Println()
	}

	// 32. Loop to remove duplicates from slice
	unique := []int{}
	seen := map[int]bool{}
	for _, n := range []int{1, 2, 2, 3} {
		if !seen[n] {
			unique = append(unique, n)
			seen[n] = true
		}
	}
	fmt.Println("Unique:", unique)

	// 33. Loop to count words in string
	text := "go is fun"
	wordCount := 1
	for _, c := range text {
		if c == ' ' {
			wordCount++
		}
	}
	fmt.Println("Word count:", wordCount)

	// 34. Loop to print ASCII values
	for c := 'A'; c <= 'C'; c++ {
		fmt.Printf("%c:%d ", c, c)
	}
	fmt.Println()

	// 35. Loop to sum digits of number
	num := 1234
	sumDigits := 0
	for num > 0 {
		sumDigits += num % 10
		num /= 10
	}
	fmt.Println("Sum digits:", sumDigits)

	// 36. Loop to print powers of 2
	for n := 1; n <= 16; n *= 2 {
		fmt.Println("Power of 2:", n)
	}

	// 37. Loop to print even numbers
	for n := 2; n <= 10; n += 2 {
		fmt.Println("Even:", n)
	}

	// 38. Loop to print odd numbers
	for n := 1; n <= 9; n += 2 {
		fmt.Println("Odd:", n)
	}

	// 39. Loop to print reverse numbers
	for n := 5; n > 0; n-- {
		fmt.Println("Reverse:", n)
	}

	// 40. Loop to print squares
	for n := 1; n <= 5; n++ {
		fmt.Println("Square:", n*n)
	}

	// 41. Loop to print cube
	for n := 1; n <= 3; n++ {
		fmt.Println("Cube:", n*n*n)
	}

	// 42. Loop to print factorial
	fact := 1
	for n := 1; n <= 5; n++ {
		fact *= n
	}
	fmt.Println("Factorial 5:", fact)

	// 43. Loop to print divisors
	num = 12
	for d := 1; d <= num; d++ {
		if num%d == 0 {
			fmt.Println("Divisor:", d)
		}
	}

	// 44. Loop to print prime numbers up to 10
	for n := 2; n <= 10; n++ {
		isPrime := true
		for d := 2; d*d <= n; d++ {
			if n%d == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println("Prime:", n)
		}
	}

	// 45. Loop to print triangle pattern
	for i := 1; i <= 3; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// 46. Loop to print reverse triangle
	for i := 3; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// 47. Loop to print table of n
	n := 3
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", n, i, n*i)
	}

	// 48. Loop to print sum of series
	sumSeries := 0
	for i := 1; i <= 5; i++ {
		sumSeries += i
	}
	fmt.Println("Sum series:", sumSeries)

	// 49. Loop to print alternating series
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			fmt.Print(-i, " ")
		} else {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()

	// 50. Loop to print index and value of slice
	for i, v := range []string{"go", "is", "awesome"} {
		fmt.Printf("Index %d: %s\n", i, v)
	}
}

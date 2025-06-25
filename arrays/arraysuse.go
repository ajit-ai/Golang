package main

import (
	"fmt"
)

func main() {
	// 1. Declare an array of 5 ints
	var a [5]int
	fmt.Println("1:", a)

	// 2. Set and get array value
	a[2] = 7
	fmt.Println("2:", a[2])

	// 3. Array literal initialization
	b := [3]int{1, 2, 3}
	fmt.Println("3:", b)

	// 4. Array with inferred length
	c := [...]int{4, 5, 6}
	fmt.Println("4:", c)

	// 5. Array with specific index initialization
	d := [...]int{2: 10, 4: 20}
	fmt.Println("5:", d)

	// 6. Array of strings
	e := [2]string{"go", "lang"}
	fmt.Println("6:", e)

	// 7. Array of booleans
	f := [3]bool{true, false, true}
	fmt.Println("7:", f)

	// 8. Array of floats
	g := [2]float64{3.14, 2.71}
	fmt.Println("8:", g)

	// 9. Array of arrays (2D array)
	var h [2][3]int
	fmt.Println("9:", h)

	// 10. Set values in 2D array
	for i := range h {
		for j := range h[i] {
			h[i][j] = i + j
		}
	}
	fmt.Println("10:", h)

	// 11. Array length
	fmt.Println("11:", len(a))

	// 12. Iterate with for loop
	for i := 0; i < len(b); i++ {
		fmt.Print("12:", b[i], " ")
	}
	fmt.Println()

	// 13. Iterate with range
	for i, v := range c {
		fmt.Printf("13: idx %d val %d\n", i, v)
	}

	// 14. Copy array
	i := b
	fmt.Println("14:", i)

	// 15. Compare arrays
	fmt.Println("15:", b == i)

	// 16. Array of structs
	type Point struct{ x, y int }
	j := [2]Point{{1, 2}, {3, 4}}
	fmt.Println("16:", j)

	// 17. Array of pointers
	k := [2]*int{&a[0], &a[1]}
	fmt.Println("17:", *k[0], *k[1])

	// 18. Array as function argument
	printArray(b)

	// 19. Array as function return
	l := makeArray()
	fmt.Println("19:", l)

	// 20. Zero value array
	var m [4]int
	fmt.Println("20:", m)

	// 21. Array of bytes (string to array)
	n := [5]byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println("21:", n)

	// 22. Array of runes
	o := [3]rune{'你', '好', '！'}
	fmt.Println("22:", o)

	// 23. Array of interfaces
	var p [2]interface{}
	p[0] = 42
	p[1] = "go"
	fmt.Println("23:", p)

	// 24. Array of bools, all true
	q := [3]bool{true, true, true}
	fmt.Println("24:", q)

	// 25. Array with only one non-zero value
	r := [5]int{2: 99}
	fmt.Println("25:", r)

	// 26. Array with descending values
	s := [5]int{5, 4, 3, 2, 1}
	fmt.Println("26:", s)

	// 27. Array with computed values
	var t [5]int
	for i := range t {
		t[i] = i * i
	}
	fmt.Println("27:", t)

	// 28. Array of arrays (3D)
	var u [2][2][2]int
	u[1][1][1] = 8
	fmt.Println("28:", u)

	// 29. Array of slices (not common, but possible)
	var v [2][]int
	v[0] = []int{1, 2}
	v[1] = []int{3, 4}
	fmt.Println("29:", v)

	// 30. Array of maps (not common, but possible)
	var w [2]map[string]int
	w[0] = map[string]int{"a": 1}
	w[1] = map[string]int{"b": 2}
	fmt.Println("30:", w)

	// 31. Array of channels
	var x [2]chan int
	x[0] = make(chan int, 1)
	x[1] = make(chan int, 1)
	x[0] <- 10
	x[1] <- 20
	fmt.Println("31:", <-x[0], <-x[1])

	// 32. Array with custom type
	type MyInt int
	y := [2]MyInt{1, 2}
	fmt.Println("32:", y)

	// 33. Array with constants
	const (
		A = iota
		B
		C
	)
	z := [3]int{A, B, C}
	fmt.Println("33:", z)

	// 34. Array with anonymous struct
	aa := [2]struct{ name string }{{"foo"}, {"bar"}}
	fmt.Println("34:", aa)

	// 35. Array with nil pointers
	var bb [2]*int
	fmt.Println("35:", bb)

	// 36. Array with initialized pointers
	cc := [2]*int{new(int), new(int)}
	*cc[0] = 11
	*cc[1] = 22
	fmt.Println("36:", *cc[0], *cc[1])

	// 37. Array with bool toggle
	dd := [2]bool{}
	dd[0] = !dd[1]
	fmt.Println("37:", dd)

	// 38. Array with float math
	ee := [3]float64{1.1, 2.2, 3.3}
	sum := 0.0
	for _, v := range ee {
		sum += v
	}
	fmt.Println("38: sum", sum)

	// 39. Array with reversed values
	ff := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(ff)/2; i++ {
		ff[i], ff[len(ff)-1-i] = ff[len(ff)-1-i], ff[i]
	}
	fmt.Println("39: reversed", ff)

	// 40. Array with min/max
	gg := [5]int{3, 1, 4, 1, 5}
	min, max := gg[0], gg[0]
	for _, v := range gg {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	fmt.Println("40: min", min, "max", max)

	// 41. Array with sum of even numbers
	hh := [5]int{1, 2, 3, 4, 5}
	evenSum := 0
	for _, v := range hh {
		if v%2 == 0 {
			evenSum += v
		}
	}
	fmt.Println("41: even sum", evenSum)

	// 42. Array with count of positive numbers
	ii := [5]int{-1, 2, -3, 4, 5}
	count := 0
	for _, v := range ii {
		if v > 0 {
			count++
		}
	}
	fmt.Println("42: positive count", count)

	// 43. Array with all zeros
	var jj [5]int
	fmt.Println("43:", jj)

	// 44. Array with all ones
	kk := [5]int{}
	for i := range kk {
		kk[i] = 1
	}
	fmt.Println("44:", kk)

	// 45. Array with alternating values
	ll := [6]int{}
	for i := range ll {
		if i%2 == 0 {
			ll[i] = 1
		} else {
			ll[i] = -1
		}
	}
	fmt.Println("45:", ll)

	// 46. Array with Fibonacci numbers
	mm := [7]int{0, 1}
	for i := 2; i < len(mm); i++ {
		mm[i] = mm[i-1] + mm[i-2]
	}
	fmt.Println("46: Fibonacci", mm)

	// 47. Array with unique values
	nn := [5]int{1, 2, 3, 4, 5}
	unique := true
	for i := 0; i < len(nn); i++ {
		for j := i + 1; j < len(nn); j++ {
			if nn[i] == nn[j] {
				unique = false
			}
		}
	}
	fmt.Println("47: unique?", unique)

	// 48. Array with duplicate check
	oo := [5]int{1, 2, 2, 4, 5}
	hasDup := false
	for i := 0; i < len(oo); i++ {
		for j := i + 1; j < len(oo); j++ {
			if oo[i] == oo[j] {
				hasDup = true
			}
		}
	}
	fmt.Println("48: has duplicate?", hasDup)

	// 49. Array with sum of all elements
	pp := [5]int{1, 2, 3, 4, 5}
	total := 0
	for _, v := range pp {
		total += v
	}
	fmt.Println("49: total sum", total)

	// 50. Array with product of all elements
	qq := [5]int{1, 2, 3, 4, 5}
	prod := 1
	for _, v := range qq {
		prod *= v
	}
	fmt.Println("50: total product", prod)
}

// Example function taking array as argument
func printArray(arr [3]int) {
	fmt.Println("18: printArray", arr)
}

// Example function returning array
func makeArray() [3]int {
	return [3]int{7, 8, 9}
}

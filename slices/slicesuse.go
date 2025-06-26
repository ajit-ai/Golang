package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	// 1. Declare an empty slice
	var a []int
	fmt.Println("1:", a)

	// 2. Check if slice is nil
	fmt.Println("2:", a == nil)

	// 3. Create slice with make
	b := make([]int, 5)
	fmt.Println("3:", b)

	// 4. Set values
	b[0] = 10
	b[1] = 20
	fmt.Println("4:", b)

	// 5. Get length and capacity
	fmt.Println("5: len", len(b), "cap", cap(b))

	// 6. Append single value
	b = append(b, 30)
	fmt.Println("6:", b)

	// 7. Append multiple values
	b = append(b, 40, 50)
	fmt.Println("7:", b)

	// 8. Copy slice
	c := make([]int, len(b))
	copy(c, b)
	fmt.Println("8:", c)

	// 9. Slice literal
	d := []string{"go", "lang"}
	fmt.Println("9:", d)

	// 10. Slice a slice
	e := b[1:4]
	fmt.Println("10:", e)

	// 11. Slice from start
	fmt.Println("11:", b[:3])

	// 12. Slice to end
	fmt.Println("12:", b[2:])

	// 13. Iterate with for
	for i := 0; i < len(b); i++ {
		fmt.Print("13:", b[i], " ")
	}
	fmt.Println()

	// 14. Iterate with range
	for i, v := range d {
		fmt.Printf("14: idx %d val %s\n", i, v)
	}

	// 15. Compare slices (Go 1.18+)
	f := []int{1, 2, 3}
	g := []int{1, 2, 3}
	fmt.Println("15:", slices.Equal(f, g))

	// 16. Remove element (by index)
	h := []int{1, 2, 3, 4}
	h = append(h[:2], h[3:]...)
	fmt.Println("16:", h)

	// 17. Insert element (at index)
	i := []int{1, 2, 4, 5}
	idx := 2
	val := 3
	i = append(i[:idx], append([]int{val}, i[idx:]...)...)
	fmt.Println("17:", i)

	// 18. Reverse slice
	j := []int{1, 2, 3, 4}
	for l, r := 0, len(j)-1; l < r; l, r = l+1, r-1 {
		j[l], j[r] = j[r], j[l]
	}
	fmt.Println("18:", j)

	// 19. Find value
	k := []string{"a", "b", "c"}
	found := false
	for _, v := range k {
		if v == "b" {
			found = true
			break
		}
	}
	fmt.Println("19: found b?", found)

	// 20. Filter slice
	l := []int{1, 2, 3, 4, 5}
	var filtered []int
	for _, v := range l {
		if v%2 == 0 {
			filtered = append(filtered, v)
		}
	}
	fmt.Println("20:", filtered)

	// 21. Map slice (double values)
	m := []int{1, 2, 3}
	for i := range m {
		m[i] *= 2
	}
	fmt.Println("21:", m)

	// 22. Sum slice
	n := []int{1, 2, 3, 4}
	sum := 0
	for _, v := range n {
		sum += v
	}
	fmt.Println("22: sum", sum)

	// 23. Min/max in slice
	o := []int{5, 2, 8, 1}
	min, max := o[0], o[0]
	for _, v := range o {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	fmt.Println("23: min", min, "max", max)

	// 24. Remove duplicates
	p := []int{1, 2, 2, 3, 4, 4}
	unique := []int{}
	seen := map[int]bool{}
	for _, v := range p {
		if !seen[v] {
			unique = append(unique, v)
			seen[v] = true
		}
	}
	fmt.Println("24:", unique)

	// 25. Join slice of strings
	q := []string{"go", "is", "fun"}
	fmt.Println("25:", strings.Join(q, " "))

	// 26. Split string to slice
	r := strings.Split("a,b,c", ",")
	fmt.Println("26:", r)

	// 27. Slice of slices (2D)
	s := make([][]int, 2)
	for i := range s {
		s[i] = make([]int, 3)
	}
	fmt.Println("27:", s)

	// 28. Fill slice with value
	t := make([]int, 5)
	for i := range t {
		t[i] = 7
	}
	fmt.Println("28:", t)

	// 29. Grow slice with append in loop
	u := []int{}
	for i := 0; i < 5; i++ {
		u = append(u, i)
	}
	fmt.Println("29:", u)

	// 30. Shrink slice
	v := []int{1, 2, 3, 4, 5}
	v = v[:3]
	fmt.Println("30:", v)

	// 31. Clear slice
	w := []int{1, 2, 3}
	w = w[:0]
	fmt.Println("31:", w)

	// 32. Copy part of slice
	x := []int{1, 2, 3, 4, 5}
	y := make([]int, 2)
	copy(y, x[1:3])
	fmt.Println("32:", y)

	// 33. Slice capacity doubling
	z := make([]int, 0, 2)
	for i := 0; i < 5; i++ {
		z = append(z, i)
		fmt.Println("33: len", len(z), "cap", cap(z))
	}

	// 34. Slice with custom type
	type Point struct{ x, y int }
	aa := []Point{{1, 2}, {3, 4}}
	fmt.Println("34:", aa)

	// 35. Slice of pointers
	bb := []*int{new(int), new(int)}
	*bb[0] = 10
	*bb[1] = 20
	fmt.Println("35:", *bb[0], *bb[1])

	// 36. Slice as function argument
	printSlice([]int{1, 2, 3})

	// 37. Slice as function return
	cc := makeSlice()
	fmt.Println("37:", cc)

	// 38. Slice of interfaces
	var dd []interface{}
	dd = append(dd, 1, "go", 3.14)
	fmt.Println("38:", dd)

	// 39. Slice of bools
	ee := []bool{true, false, true}
	fmt.Println("39:", ee)

	// 40. Slice of runes
	ff := []rune("hello")
	fmt.Println("40:", ff)

	// 41. Slice of bytes
	gg := []byte("golang")
	fmt.Println("41:", gg)

	// 42. Slice with only one non-zero value
	hh := make([]int, 5)
	hh[2] = 99
	fmt.Println("42:", hh)

	// 43. Slice with descending values
	ii := []int{5, 4, 3, 2, 1}
	fmt.Println("43:", ii)

	// 44. Slice with computed values
	jj := make([]int, 5)
	for i := range jj {
		jj[i] = i * i
	}
	fmt.Println("44:", jj)

	// 45. Slice with alternating values
	kk := make([]int, 6)
	for i := range kk {
		if i%2 == 0 {
			kk[i] = 1
		} else {
			kk[i] = -1
		}
	}
	fmt.Println("45:", kk)

	// 46. Slice with Fibonacci numbers
	ll := make([]int, 7)
	ll[0], ll[1] = 0, 1
	for i := 2; i < len(ll); i++ {
		ll[i] = ll[i-1] + ll[i-2]
	}
	fmt.Println("46:", ll)

	// 47. Slice with all zeros
	mm := make([]int, 5)
	fmt.Println("47:", mm)

	// 48. Slice with all ones
	nn := make([]int, 5)
	for i := range nn {
		nn[i] = 1
	}
	fmt.Println("48:", nn)

	// 49. Slice with sum of even numbers
	oo := []int{1, 2, 3, 4, 5}
	evenSum := 0
	for _, v := range oo {
		if v%2 == 0 {
			evenSum += v
		}
	}
	fmt.Println("49: even sum", evenSum)

	// 50. Slice with count of positive numbers
	pp := []int{-1, 2, -3, 4, 5}
	count := 0
	for _, v := range pp {
		if v > 0 {
			count++
		}
	}
	fmt.Println("50: positive count", count)

	// 51. Slice with min/max index
	qq := []int{3, 1, 4, 1, 5}
	minIdx, maxIdx := 0, 0
	for i, v := range qq {
		if v < qq[minIdx] {
			minIdx = i
		}
		if v > qq[maxIdx] {
			maxIdx = i
		}
	}
	fmt.Println("51: minIdx", minIdx, "maxIdx", maxIdx)

	// 52. Slice with product of all elements
	rr := []int{1, 2, 3, 4, 5}
	prod := 1
	for _, v := range rr {
		prod *= v
	}
	fmt.Println("52: product", prod)

	// 53. Slice with palindrome check
	ss := []int{1, 2, 3, 2, 1}
	isPalindrome := true
	for i := 0; i < len(ss)/2; i++ {
		if ss[i] != ss[len(ss)-1-i] {
			isPalindrome = false
			break
		}
	}
	fmt.Println("53: palindrome?", isPalindrome)

	// 54. Slice with sum of digits
	tt := []int{1, 2, 3, 4}
	digitSum := 0
	for _, v := range tt {
		for v > 0 {
			digitSum += v % 10
			v /= 10
		}
	}
	fmt.Println("54: digit sum", digitSum)

	// 55. Slice with all true
	uu := make([]bool, 5)
	for i := range uu {
		uu[i] = true
	}
	fmt.Println("55:", uu)

	// 56. Slice with all false
	vv := make([]bool, 5)
	fmt.Println("56:", vv)

	// 57. Slice with alternating bools
	ww := make([]bool, 6)
	for i := range ww {
		ww[i] = i%2 == 0
	}
	fmt.Println("57:", ww)

	// 58. Slice with string concatenation
	xx := []string{"a", "b", "c"}
	concat := ""
	for _, v := range xx {
		concat += v
	}
	fmt.Println("58:", concat)

	// 59. Slice with uppercase strings
	yy := []string{"go", "lang"}
	for i := range yy {
		yy[i] = strings.ToUpper(yy[i])
	}
	fmt.Println("59:", yy)

	// 60. Slice with lowercase strings
	zz := []string{"Go", "LANG"}
	for i := range zz {
		zz[i] = strings.ToLower(zz[i])
	}
	fmt.Println("60:", zz)

	// 61. Slice with string length
	aaa := []string{"go", "lang"}
	lengths := make([]int, len(aaa))
	for i, v := range aaa {
		lengths[i] = len(v)
	}
	fmt.Println("61:", lengths)

	// 62. Slice with string contains
	bbb := []string{"go", "lang", "fun"}
	containsGo := false
	for _, v := range bbb {
		if strings.Contains(v, "go") {
			containsGo = true
		}
	}
	fmt.Println("62: contains 'go'?", containsGo)

	// 63. Slice with prefix check
	ccc := []string{"go", "lang", "fun"}
	hasPrefix := false
	for _, v := range ccc {
		if strings.HasPrefix(v, "g") {
			hasPrefix = true
		}
	}
	fmt.Println("63: has prefix 'g'?", hasPrefix)

	// 64. Slice with suffix check
	ddd := []string{"go", "lang", "fun"}
	hasSuffix := false
	for _, v := range ddd {
		if strings.HasSuffix(v, "n") {
			hasSuffix = true
		}
	}
	fmt.Println("64: has suffix 'n'?", hasSuffix)

	// 65. Slice with count of vowels
	eee := []string{"go", "lang"}
	vowelCount := 0
	for _, v := range eee {
		for _, c := range v {
			if strings.ContainsRune("aeiou", c) {
				vowelCount++
			}
		}
	}
	fmt.Println("65: vowel count", vowelCount)

	// 66. Slice with word count
	fff := []string{"go is fun", "hello world"}
	wordCount := 0
	for _, v := range fff {
		wordCount += len(strings.Fields(v))
	}
	fmt.Println("66: word count", wordCount)

	// 67. Slice with ASCII values
	ggg := []byte("go")
	for _, v := range ggg {
		fmt.Print("67:", v, " ")
	}
	fmt.Println()

	// 68. Slice with rune values
	hhh := []rune("go")
	for _, v := range hhh {
		fmt.Print("68:", v, " ")
	}
	fmt.Println()

	// 69. Slice with index of value
	iii := []int{1, 2, 3, 4}
	idx = -1
	for i, v := range iii {
		if v == 3 {
			idx = i
			break
		}
	}
	fmt.Println("69: index of 3", idx)

	// 70. Slice with remove by value
	jjj := []int{1, 2, 3, 4}
	val = 2
	for i, v := range jjj {
		if v == val {
			jjj = append(jjj[:i], jjj[i+1:]...)
			break
		}
	}
	fmt.Println("70: removed 2", jjj)

	// 71. Slice with insert at beginning
	kkk := []int{2, 3, 4}
	kkk = append([]int{1}, kkk...)
	fmt.Println("71:", kkk)

	// 72. Slice with insert at end
	lll := []int{1, 2, 3}
	lll = append(lll, 4)
	fmt.Println("72:", lll)

	// 73. Slice with insert at middle
	mmm := []int{1, 2, 4, 5}
	mmm = append(mmm[:2], append([]int{3}, mmm[2:]...)...)
	fmt.Println("73:", mmm)

	// 74. Slice with remove from beginning
	nnn := []int{1, 2, 3, 4}
	nnn = nnn[1:]
	fmt.Println("74:", nnn)

	// 75. Slice with remove from end
	ooo := []int{1, 2, 3, 4}
	ooo = ooo[:len(ooo)-1]
	fmt.Println("75:", ooo)

	// 76. Slice with remove from middle
	ppp := []int{1, 2, 3, 4}
	ppp = append(ppp[:2], ppp[3:]...)
	fmt.Println("76:", ppp)

	// 77. Slice with clear (reset to empty)
	qqq := []int{1, 2, 3}
	qqq = qqq[:0]
	fmt.Println("77:", qqq)

	// 78. Slice with nil assignment
	var rrr []int
	rrr = nil
	fmt.Println("78:", rrr)

	// 79. Slice with capacity reservation
	sss := make([]int, 0, 10)
	fmt.Println("79: cap", cap(sss))

	// 80. Slice with append in loop (dynamic growth)
	ttt := []int{}
	for i := 0; i < 10; i++ {
		ttt = append(ttt, i)
	}
	fmt.Println("80:", ttt)

	// 81. Slice with sum of squares
	uuu := []int{1, 2, 3}
	sqSum := 0
	for _, v := range uuu {
		sqSum += v * v
	}
	fmt.Println("81: sum of squares", sqSum)

	// 82. Slice with all positive check
	vvv := []int{1, 2, 3}
	allPos := true
	for _, v := range vvv {
		if v <= 0 {
			allPos = false
			break
		}
	}
	fmt.Println("82: all positive?", allPos)

	// 83. Slice with any negative check
	www := []int{1, -2, 3}
	anyNeg := false
	for _, v := range www {
		if v < 0 {
			anyNeg = true
			break
		}
	}
	fmt.Println("83: any negative?", anyNeg)

	// 84. Slice with count of zeros
	xxx := []int{0, 1, 0, 2, 0}
	zeroCount := 0
	for _, v := range xxx {
		if v == 0 {
			zeroCount++
		}
	}
	fmt.Println("84: zero count", zeroCount)

	// 85. Slice with running total
	yyy := []int{1, 2, 3, 4}
	running := make([]int, len(yyy))
	total := 0
	for i, v := range yyy {
		total += v
		running[i] = total
	}
	fmt.Println("85: running total", running)

	// 86. Slice with even/odd separation
	zzz := []int{1, 2, 3, 4, 5}
	evens, odds := []int{}, []int{}
	for _, v := range zzz {
		if v%2 == 0 {
			evens = append(evens, v)
		} else {
			odds = append(odds, v)
		}
	}
	fmt.Println("86: evens", evens, "odds", odds)

	// 87. Slice with flattening 2D slice
	aaaa := [][]int{{1, 2}, {3, 4}}
	flat := []int{}
	for _, row := range aaaa {
		flat = append(flat, row...)
	}
	fmt.Println("87: flattened", flat)

	// 88. Slice with chunking
	bbbb := []int{1, 2, 3, 4, 5, 6}
	chunkSize := 2
	chunks := [][]int{}
	for i := 0; i < len(bbbb); i += chunkSize {
		end := i + chunkSize
		if end > len(bbbb) {
			end = len(bbbb)
		}
		chunks = append(chunks, bbbb[i:end])
	}
	fmt.Println("88: chunks", chunks)

	// 89. Slice with zip (pairing)
	cccc := []int{1, 2, 3}
	dddd := []string{"a", "b", "c"}
	zipped := []struct {
		n int
		s string
	}{}
	for i := 0; i < len(cccc) && i < len(dddd); i++ {
		zipped = append(zipped, struct {
			n int
			s string
		}{cccc[i], dddd[i]})
	}
	fmt.Println("89: zipped", zipped)

	// 90. Slice with unzipping
	nums := []int{}
	strs := []string{}
	for _, pair := range zipped {
		nums = append(nums, pair.n)
		strs = append(strs, pair.s)
	}
	fmt.Println("90: unzipped", nums, strs)

	// 91. Slice with rotate left
	eeee := []int{1, 2, 3, 4}
	eeee = append(eeee[1:], eeee[0])
	fmt.Println("91: rotate left", eeee)

	// 92. Slice with rotate right
	ffff := []int{1, 2, 3, 4}
	ffff = append([]int{ffff[len(ffff)-1]}, ffff[:len(ffff)-1]...)
	fmt.Println("92: rotate right", ffff)

	// 93. Slice with duplicate each element
	gggg := []int{1, 2, 3}
	dup := []int{}
	for _, v := range gggg {
		dup = append(dup, v, v)
	}
	fmt.Println("93: duplicated", dup)

	// 94. Slice with interleave
	hhhh := []int{1, 3, 5}
	iiii := []int{2, 4, 6}
	interleaved := []int{}
	for i := 0; i < len(hhhh) && i < len(iiii); i++ {
		interleaved = append(interleaved, hhhh[i], iiii[i])
	}
	fmt.Println("94: interleaved", interleaved)

	// 95. Slice with windowed sum
	jjjj := []int{1, 2, 3, 4, 5}
	window := 3
	windowSums := []int{}
	for i := 0; i <= len(jjjj)-window; i++ {
		sum := 0
		for j := 0; j < window; j++ {
			sum += jjjj[i+j]
		}
		windowSums = append(windowSums, sum)
	}
	fmt.Println("95: window sums", windowSums)

	// 96. Slice with index/value swap
	kkkk := []int{0, 1, 2, 3}
	for i := range kkkk {
		kkkk[i] = i
	}
	fmt.Println("96:", kkkk)

	// 97. Slice with all indices
	llll := make([]int, 5)
	for i := range llll {
		llll[i] = i
	}
	fmt.Println("97:", llll)

	// 98. Slice with all values squared
	mmmm := []int{1, 2, 3, 4}
	for i := range mmmm {
		mmmm[i] *= mmmm[i]
	}
	fmt.Println("98:", mmmm)

	// 99. Slice with all values cubed
	nnnn := []int{1, 2, 3, 4}
	for i := range nnnn {
		nnnn[i] = nnnn[i] * nnnn[i] * nnnn[i]
	}
	fmt.Println("99:", nnnn)

	// 100. Slice with all values negated
	oooo := []int{1, -2, 3, -4}
	for i := range oooo {
		oooo[i] = -oooo[i]
	}
	fmt.Println("100:", oooo)
}

func printSlice(s []int) {
	fmt.Println("36:", s)
}

func makeSlice() []int {
	return []int{7, 8, 9}
}

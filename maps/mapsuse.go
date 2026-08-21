package main

import (
	"fmt"
	"maps"
	"strings"
)

func MapsUseMain() {
	// 1. Declare a map
	var a map[string]int
	fmt.Println("1:", a)

	// 2. Make a map
	b := make(map[string]int)
	fmt.Println("2:", b)

	// 3. Map literal
	c := map[string]int{"a": 1, "b": 2}
	fmt.Println("3:", c)

	// 4. Set value
	c["c"] = 3
	fmt.Println("4:", c)

	// 5. Get value
	fmt.Println("5:", c["a"])

	// 6. Get non-existent key (zero value)
	fmt.Println("6:", c["z"])

	// 7. Check key existence
	_, ok := c["b"]
	fmt.Println("7:", ok)

	// 8. Delete key
	delete(c, "b")
	fmt.Println("8:", c)

	// 9. Clear map (Go 1.18+)
	clear(c)
	fmt.Println("9:", c)

	// 10. Map with int keys
	d := map[int]string{1: "one", 2: "two"}
	fmt.Println("10:", d)

	// 11. Map with struct keys
	type Point struct{ x, y int }
	e := map[Point]string{{1, 2}: "A", {3, 4}: "B"}
	fmt.Println("11:", e)

	// 12. Map with slice values
	f := map[string][]int{"nums": {1, 2, 3}}
	fmt.Println("12:", f)

	// 13. Map with map values
	g := map[string]map[string]int{"outer": {"inner": 1}}
	fmt.Println("13:", g)

	// 14. Iterate map
	h := map[string]int{"a": 1, "b": 2}
	for k, v := range h {
		fmt.Printf("14: %s=%d\n", k, v)
	}

	// 15. Map length
	fmt.Println("15:", len(h))

	// 16. Map of bools (set)
	set := map[string]bool{"a": true, "b": true}
	fmt.Println("16:", set)

	// 17. Map of interfaces
	i := map[string]interface{}{"num": 1, "str": "go"}
	fmt.Println("17:", i)

	// 18. Map as function argument
	printMap(h)

	// 19. Map as function return
	j := makeMap()
	fmt.Println("19:", j)

	// 20. Map with default value logic
	k := map[string]int{}
	if _, ok := k["foo"]; !ok {
		k["foo"] = 42
	}
	fmt.Println("20:", k)

	// 21. Map with increment
	l := map[string]int{}
	l["count"]++
	fmt.Println("21:", l)

	// 22. Map with decrement
	l["count"]--
	fmt.Println("22:", l)

	// 23. Map with composite literal
	m := map[int][]string{
		1: {"a", "b"},
		2: {"c"},
	}
	fmt.Println("23:", m)

	// 24. Map with nested struct value
	type User struct{ name string }
	n := map[int]User{1: {"Alice"}, 2: {"Bob"}}
	fmt.Println("24:", n)

	// 25. Map with pointer values
	o := map[string]*int{}
	val := 10
	o["ptr"] = &val
	fmt.Println("25:", *o["ptr"])

	// 26. Map with pointer keys
	p := map[*int]string{}
	ptr := &val
	p[ptr] = "value"
	fmt.Println("26:", p[ptr])

	// 27. Map with rune keys
	q := map[rune]string{'a': "A", 'b': "B"}
	fmt.Println("27:", q)

	// 28. Map with byte keys
	r := map[byte]string{'x': "X", 'y': "Y"}
	fmt.Println("28:", r)

	// 29. Map with float keys
	s := map[float64]string{3.14: "pi", 2.71: "e"}
	fmt.Println("29:", s)

	// 30. Map with bool keys
	t := map[bool]string{true: "yes", false: "no"}
	fmt.Println("30:", t)

	// 31. Map with array keys
	u := map[[2]int]string{{1, 2}: "pair"}
	fmt.Println("31:", u)

	// 32. Map with zero value check
	v := map[string]int{}
	fmt.Println("32:", v["missing"])

	// 33. Map with key overwrite
	w := map[string]int{"a": 1}
	w["a"] = 2
	fmt.Println("33:", w)

	// 34. Map with key removal
	delete(w, "a")
	fmt.Println("34:", w)

	// 35. Map with all keys
	x := map[string]int{"a": 1, "b": 2}
	for k := range x {
		fmt.Println("35: key", k)
	}

	// 36. Map with all values
	for _, v := range x {
		fmt.Println("36: value", v)
	}

	// 37. Map with sum of values
	sum := 0
	for _, v := range x {
		sum += v
	}
	fmt.Println("37: sum", sum)

	// 38. Map with max value
	max := 0
	for _, v := range x {
		if v > max {
			max = v
		}
	}
	fmt.Println("38: max", max)

	// 39. Map with min value
	min := 1<<31 - 1
	for _, v := range x {
		if v < min {
			min = v
		}
	}
	fmt.Println("39: min", min)

	// 40. Map with key existence check
	if _, ok := x["a"]; ok {
		fmt.Println("40: key a exists")
	}

	// 41. Map with value existence check
	valExists := false
	for _, v := range x {
		if v == 2 {
			valExists = true
		}
	}
	fmt.Println("41: value 2 exists?", valExists)

	// 42. Map with count of keys
	fmt.Println("42: key count", len(x))

	// 43. Map with count of values > 1
	count := 0
	for _, v := range x {
		if v > 1 {
			count++
		}
	}
	fmt.Println("43: values > 1 count", count)

	// 44. Map with string keys and int values
	y := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("44:", y)

	// 45. Map with string keys and string values
	z := map[string]string{"foo": "bar"}
	fmt.Println("45:", z)

	// 46. Map with int keys and string values
	aa := map[int]string{1: "one"}
	fmt.Println("46:", aa)

	// 47. Map with int keys and bool values
	bb := map[int]bool{1: true, 2: false}
	fmt.Println("47:", bb)

	// 48. Map with string keys and bool values
	cc := map[string]bool{"a": true}
	fmt.Println("48:", cc)

	// 49. Map with string keys and float values
	dd := map[string]float64{"pi": 3.14}
	fmt.Println("49:", dd)

	// 50. Map with string keys and slice values
	ee := map[string][]int{"nums": {1, 2, 3}}
	fmt.Println("50:", ee)

	// 51. Map with string keys and map values
	ff := map[string]map[string]int{"outer": {"inner": 1}}
	fmt.Println("51:", ff)

	// 52. Map with string keys and struct values
	type Item struct{ id int }
	gg := map[string]Item{"item1": {1}}
	fmt.Println("52:", gg)

	// 53. Map with string keys and pointer to struct values
	hh := map[string]*Item{"item2": &Item{2}}
	fmt.Println("53:", hh)

	// 54. Map with string keys and interface values
	ii := map[string]interface{}{"num": 1, "str": "go"}
	fmt.Println("54:", ii)

	// 55. Map with string keys and function values
	jj := map[string]func(int) int{"double": func(x int) int { return x * 2 }}
	fmt.Println("55:", jj["double"](5))

	// 56. Map with string keys and channel values
	kk := map[string]chan int{"ch": make(chan int, 1)}
	kk["ch"] <- 42
	fmt.Println("56:", <-kk["ch"])

	// 57. Map with string keys and array values
	ll := map[string][2]int{"pair": {1, 2}}
	fmt.Println("57:", ll)

	// 58. Map with string keys and rune values
	mm := map[string]rune{"a": 'A'}
	fmt.Println("58:", mm)

	// 59. Map with string keys and byte values
	nn := map[string]byte{"b": 'B'}
	fmt.Println("59:", nn)

	// 60. Map with string keys and bool toggle
	oo := map[string]bool{"flag": false}
	oo["flag"] = !oo["flag"]
	fmt.Println("60:", oo)

	// 61. Map with string keys and increment
	pp := map[string]int{"count": 0}
	pp["count"]++
	fmt.Println("61:", pp)

	// 62. Map with string keys and decrement
	pp["count"]--
	fmt.Println("62:", pp)

	// 63. Map with string keys and append to slice value
	qq := map[string][]int{"nums": {1, 2}}
	qq["nums"] = append(qq["nums"], 3)
	fmt.Println("63:", qq)

	// 64. Map with string keys and remove from slice value
	qq["nums"] = qq["nums"][:2]
	fmt.Println("64:", qq)

	// 65. Map with string keys and string concatenation
	rr := map[string]string{"msg": "hello"}
	rr["msg"] += " world"
	fmt.Println("65:", rr)

	// 66. Map with string keys and string split
	ss := map[string][]string{"words": strings.Split("go is fun", " ")}
	fmt.Println("66:", ss)

	// 67. Map with string keys and join slice value
	ss["words"] = append(ss["words"], "!")
	fmt.Println("67:", strings.Join(ss["words"], " "))

	// 68. Map with string keys and count of slice value
	fmt.Println("68:", len(ss["words"]))

	// 69. Map with string keys and check for empty slice
	tt := map[string][]int{"nums": {}}
	fmt.Println("69:", len(tt["nums"]) == 0)

	// 70. Map with string keys and check for nil slice
	var uu map[string][]int
	fmt.Println("70:", uu == nil)

	// 71. Map with string keys and check for nil map
	var vv map[string]int
	fmt.Println("71:", vv == nil)

	// 72. Map with string keys and initialize if nil
	if vv == nil {
		vv = make(map[string]int)
	}
	fmt.Println("72:", vv)

	// 73. Map with string keys and default value logic
	ww := map[string]int{}
	if _, ok := ww["missing"]; !ok {
		ww["missing"] = 100
	}
	fmt.Println("73:", ww)

	// 74. Map with string keys and delete if exists
	if _, ok := ww["missing"]; ok {
		delete(ww, "missing")
	}
	fmt.Println("74:", ww)

	// 75. Map with string keys and check for all keys
	xx := map[string]int{"a": 1, "b": 2, "c": 3}
	allKeys := []string{"a", "b", "c"}
	allExist := true
	for _, k := range allKeys {
		if _, ok := xx[k]; !ok {
			allExist = false
		}
	}
	fmt.Println("75: all keys exist?", allExist)

	// 76. Map with string keys and check for any key
	anyExist := false
	for _, k := range []string{"x", "y", "a"} {
		if _, ok := xx[k]; ok {
			anyExist = true
		}
	}
	fmt.Println("76: any key exist?", anyExist)

	// 77. Map with string keys and count of keys with value > 1
	count = 0
	for _, v := range xx {
		if v > 1 {
			count++
		}
	}
	fmt.Println("77: count > 1", count)

	// 78. Map with string keys and sum of all values
	sum = 0
	for _, v := range xx {
		sum += v
	}
	fmt.Println("78: sum", sum)

	// 79. Map with string keys and find max value
	max = 0
	for _, v := range xx {
		if v > max {
			max = v
		}
	}
	fmt.Println("79: max", max)

	// 80. Map with string keys and find min value
	min = 1<<31 - 1
	for _, v := range xx {
		if v < min {
			min = v
		}
	}
	fmt.Println("80: min", min)

	// 81. Map with string keys and find key for max value
	maxKey := ""
	max = 0
	for k, v := range xx {
		if v > max {
			max = v
			maxKey = k
		}
	}
	fmt.Println("81: max key", maxKey)

	// 82. Map with string keys and find key for min value
	minKey := ""
	min = 1<<31 - 1
	for k, v := range xx {
		if v < min {
			min = v
			minKey = k
		}
	}
	fmt.Println("82: min key", minKey)

	// 83. Map with string keys and swap keys/values
	yy := map[string]int{"a": 1, "b": 2}
	swapped := map[int]string{}
	for k, v := range yy {
		swapped[v] = k
	}
	fmt.Println("83:", swapped)

	// 84. Map with string keys and reverse lookup
	val = 2
	key := ""
	for k, v := range yy {
		if v == val {
			key = k
			break
		}
	}
	fmt.Println("84: key for value 2", key)

	// 85. Map with string keys and merge two maps
	zz := map[string]int{"c": 3}
	for k, v := range zz {
		yy[k] = v
	}
	fmt.Println("85: merged", yy)

	// 86. Map with string keys and difference
	diff := map[string]int{}
	for k, v := range yy {
		if _, ok := zz[k]; !ok {
			diff[k] = v
		}
	}
	fmt.Println("86: diff", diff)

	// 87. Map with string keys and intersection
	inter := map[string]int{}
	for k, v := range yy {
		if _, ok := zz[k]; ok {
			inter[k] = v
		}
	}
	fmt.Println("87: intersection", inter)

	// 88. Map with string keys and union
	union := map[string]int{}
	for k, v := range yy {
		union[k] = v
	}
	for k, v := range zz {
		union[k] = v
	}
	fmt.Println("88: union", union)

	// 89. Map with string keys and keys as slice
	keys := []string{}
	for k := range yy {
		keys = append(keys, k)
	}
	fmt.Println("89: keys", keys)

	// 90. Map with string keys and values as slice
	values := []int{}
	for _, v := range yy {
		values = append(values, v)
	}
	fmt.Println("90: values", values)

	// 91. Map with string keys and sorted keys
	keys = []string{}
	for k := range yy {
		keys = append(keys, k)
	}
	// sort.Strings(keys) // Uncomment if you want sorted keys
	fmt.Println("91: keys (unsorted)", keys)

	// 92. Map with string keys and sorted values
	values = []int{}
	for _, v := range yy {
		values = append(values, v)
	}
	// sort.Ints(values) // Uncomment if you want sorted values
	fmt.Println("92: values (unsorted)", values)

	// 93. Map with string keys and check for duplicate values
	seenVals := map[int]bool{}
	hasDup := false
	for _, v := range yy {
		if seenVals[v] {
			hasDup = true
			break
		}
		seenVals[v] = true
	}
	fmt.Println("93: has duplicate values?", hasDup)

	// 94. Map with string keys and count of unique values
	seenVals = map[int]bool{}
	for _, v := range yy {
		seenVals[v] = true
	}
	fmt.Println("94: unique value count", len(seenVals))

	// 95. Map with string keys and clear all values
	for k := range yy {
		yy[k] = 0
	}
	fmt.Println("95: cleared values", yy)

	// 96. Map with string keys and remove all keys
	for k := range yy {
		delete(yy, k)
	}
	fmt.Println("96: removed all keys", yy)

	// 97. Map with string keys and check if empty
	fmt.Println("97: is empty?", len(yy) == 0)

	// 98. Map with string keys and copy to new map
	copyMap := map[string]int{}
	for k, v := range xx {
		copyMap[k] = v
	}
	fmt.Println("98: copy", copyMap)

	// 99. Map with string keys and compare two maps (Go 1.18+)
	xx2 := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("99: maps equal?", maps.Equal(xx, xx2))

	// 100. Map with string keys and check for subset
	subset := true
	for k, v := range xx {
		if v2, ok := xx2[k]; !ok || v2 != v {
			subset = false
		}
	}
	fmt.Println("100: is subset?", subset)
}

func printMap(m map[string]int) {
	fmt.Println("18:", m)
}

func makeMap() map[string]int {
	return map[string]int{"x": 10, "y": 20}
}

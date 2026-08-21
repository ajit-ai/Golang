package main

import (
	"fmt"
	"math"
)

const Pi = 3.1415

type Person struct {
	name string
	age  int
}

func UseVariableMain() {
	// 1. Basic declaration
	var a int
	a = 10
	fmt.Println(a)

	// 2. Declaration with initialization
	var b int = 20
	fmt.Println(b)

	// 3. Type inference
	var c = 30
	fmt.Println(c)

	// 4. Short variable declaration
	d := 40
	fmt.Println(d)

	// 5. Multiple variables
	var e, f int = 50, 60
	fmt.Println(e, f)

	// 6. Multiple short declarations
	g, h := 70, 80
	fmt.Println(g, h)

	// 7. Different types
	var i, j = "hello", 3.14
	fmt.Println(i, j)

	// 8. Zero value
	var k float64
	fmt.Println(k)

	// 9. Boolean
	var l bool = true
	fmt.Println(l)

	// 10. String
	var m string = "golang"
	fmt.Println(m)

	// 11. Array
	var n [3]int = [3]int{1, 2, 3}
	fmt.Println(n)

	// 12. Slice
	o := []string{"a", "b", "c"}
	fmt.Println(o)

	// 13. Map
	p := map[string]int{"one": 1, "two": 2}
	fmt.Println(p)

	// 14. Struct
	q := Person{name: "Alice", age: 25}
	fmt.Println(q)

	// 15. Pointer
	var r *int = &a
	fmt.Println(*r)

	// 16. Constant
	const s = 100
	fmt.Println(s)

	// 17. iota
	const (
		t = iota
		u
		v
	)
	fmt.Println(t, u, v)

	// 18. Function return
	w, _ := add(1, 2)
	fmt.Println(w)

	// 19. Anonymous variable
	_, x := add(3, 4)
	fmt.Println(x)

	// 20. Global variable
	fmt.Println(globalVar)

	// 21. Shadowing
	y := 200
	{
		y := 300
		fmt.Println(y)
	}
	fmt.Println(y)

	// 22. Rune
	var z rune = 'A'
	fmt.Println(z)

	// 23. Byte
	var aa byte = 'B'
	fmt.Println(aa)

	// 24. Complex
	var bb complex64 = 1 + 2i
	fmt.Println(bb)

	// 25. Float32
	var cc float32 = 3.14
	fmt.Println(cc)

	// 26. Float64
	var dd float64 = math.Sqrt(2)
	fmt.Println(dd)

	// 27. Unsigned int
	var ee uint = 123
	fmt.Println(ee)

	// 28. Int8
	var ff int8 = -128
	fmt.Println(ff)

	// 29. Uint8
	var gg uint8 = 255
	fmt.Println(gg)

	// 30. Int16
	var hh int16 = -32768
	fmt.Println(hh)

	// 31. Uint16
	var ii uint16 = 65535
	fmt.Println(ii)

	// 32. Int32
	var jj int32 = -2147483648
	fmt.Println(jj)

	// 33. Uint32
	var kk uint32 = 4294967295
	fmt.Println(kk)

	// 34. Int64
	var ll int64 = -9223372036854775808
	fmt.Println(ll)

	// 35. Uint64
	var mm uint64 = 18446744073709551615
	fmt.Println(mm)

	// 36. Default value for bool
	var nn bool
	fmt.Println(nn)

	// 37. Default value for string
	var oo string
	fmt.Println(oo)

	// 38. Default value for pointer
	var pp *int
	fmt.Println(pp)

	// 39. Default value for slice
	var qq []int
	fmt.Println(qq)

	// 40. Default value for map
	var rr map[string]int
	fmt.Println(rr)

	// 41. Default value for struct
	var ss Person
	fmt.Println(ss)

	// 42. Default value for array
	var tt [2]int
	fmt.Println(tt)

	// 43. Variable in for loop
	for uu := 0; uu < 1; uu++ {
		fmt.Println(uu)
	}

	// 44. Variable in if statement
	if vv := 10; vv > 5 {
		fmt.Println(vv)
	}

	// 45. Variable in switch statement
	switch ww := "go"; ww {
	case "go":
		fmt.Println(ww)
	}

	// 46. Variable from function multiple return
	xx, yy := multiReturn()
	fmt.Println(xx, yy)

	// 47. Variable from type assertion
	var zz interface{} = "hello"
	aaa := zz.(string)
	fmt.Println(aaa)

	// 48. Variable from type conversion
	var floatVal float64 = 3.14
	var bbb int = int(floatVal)
	fmt.Println(bbb)

	// 49. Variable from new()
	ccc := new(int)
	*ccc = 42
	fmt.Println(*ccc)

	// 50. Variable from make()
	ddd := make([]int, 3)
	fmt.Println(ddd)
}

var globalVar = "I am global"

func add(a, b int) (int, int) {
	return a + b, a - b
}

func multiReturn() (int, string) {
	return 1, "two"
}

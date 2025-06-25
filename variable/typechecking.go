package main

import (
	"fmt"
	"math"
)

func main() {
	// 1. int
	var a int = 10
	fmt.Println(a)

	// 2. float64
	var b float64 = 3.14
	fmt.Println(b)

	// 3. string
	var c string = "hello"
	fmt.Println(c)

	// 4. bool
	var d bool = true
	fmt.Println(d)

	// 5. int32
	var e int32 = -100
	fmt.Println(e)

	// 6. uint
	var f uint = 42
	fmt.Println(f)

	// 7. byte (alias for uint8)
	var g byte = 255
	fmt.Println(g)

	// 8. rune (alias for int32)
	var h rune = 'A'
	fmt.Println(h)

	// 9. array
	var i [3]int = [3]int{1, 2, 3}
	fmt.Println(i)

	// 10. slice
	var j []string = []string{"go", "lang"}
	fmt.Println(j)

	// 11. map
	var k map[string]int = map[string]int{"one": 1, "two": 2}
	fmt.Println(k)

	// 12. struct
	type Person struct {
		name string
		age  int
	}
	var l Person = Person{name: "Alice", age: 30}
	fmt.Println(l)

	// 13. pointer
	var m *int = &a
	fmt.Println(*m)

	// 14. complex64
	var n complex64 = 1 + 2i
	fmt.Println(n)

	// 15. constant
	const o float64 = math.Pi
	fmt.Println(o)

	// 16. function variable
	var p func(int) int = func(x int) int { return x * x }
	fmt.Println(p(5))

	// 17. interface
	var q interface{} = "interface value"
	fmt.Println(q)

	// 18. channel
	var r chan int = make(chan int, 1)
	r <- 7
	fmt.Println(<-r)

	// 19. zero value for bool
	var s bool
	fmt.Println(s)

	// 20. zero value for string
	var t string
	fmt.Println(t)
}

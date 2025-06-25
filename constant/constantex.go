package main

import (
	"fmt"
	"math"
	"time"
)

const Pi = 3.14159 // 1. Global float constant

const (
	// 2-5. Multiple constants in a block
	Zero  = 0
	One   = 1
	Two   = 2
	Three = 3
)

const Greeting string = "Hello, Go!" // 6. String constant

const MaxUint8 = 255 // 7. Unsigned integer constant

const MinInt8 int8 = -128 // 8. Signed integer constant

const (
	// 9-12. Boolean constants
	TrueConst  = true
	FalseConst = false
	IsGoFun    = true
	IsHard     = false
)

const (
	// 13-16. Constants with iota
	A = iota
	B
	C
	D
)

const (
	// 17-20. Bit shifting with iota
	Bit0 = 1 << iota
	Bit1
	Bit2
	Bit3
)

const (
	// 21-24. Typed constants
	TypedInt    int     = 42
	TypedFloat  float64 = 2.718
	TypedString string  = "typed"
	TypedBool   bool    = false
)

const (
	// 25-28. Untyped constants
	UntypedInt   = 100
	UntypedFloat = 1.23
	UntypedStr   = "untyped"
	UntypedBool  = true
)

const (
	// 29-32. Constants in expressions
	Sum      = 1 + 2
	Product  = 2 * 3
	Division = 10 / 2
	Modulus  = 10 % 3
)

const (
	// 33-36. Constants with math
	Sqrt2   = 1.41421356237
	Square4 = 4 * 4
	CircleA = Pi * 2 * 2
	CircleC = 2 * Pi * 2
)

const (
	// 37-40. Constants for time durations
	Second = 1
	Minute = 60 * Second
	Hour   = 60 * Minute
	Day    = 24 * Hour
)

const (
	// 41-44. Constants for file permissions
	Read    = 1 << 2
	Write   = 1 << 1
	Execute = 1 << 0
	All     = Read | Write | Execute
)

const (
	// 45-48. Constants for HTTP status codes
	StatusOK       = 200
	StatusNotFound = 404
	StatusError    = 500
	StatusCreated  = 201
)

const (
	// 49. Constant for array size
	ArraySize = 5
)

const (
	// 50. Constant for buffer size
	BufferSize = 1024
)

func main() {
	fmt.Println("Pi:", Pi)
	fmt.Println("Zero, One, Two, Three:", Zero, One, Two, Three)
	fmt.Println("Greeting:", Greeting)
	fmt.Println("MaxUint8:", MaxUint8)
	fmt.Println("MinInt8:", MinInt8)
	fmt.Println("TrueConst, FalseConst:", TrueConst, FalseConst)
	fmt.Println("iota A-D:", A, B, C, D)
	fmt.Println("Bit shifting:", Bit0, Bit1, Bit2, Bit3)
	fmt.Println("Typed constants:", TypedInt, TypedFloat, TypedString, TypedBool)
	fmt.Println("Untyped constants:", UntypedInt, UntypedFloat, UntypedStr, UntypedBool)
	fmt.Println("Expressions:", Sum, Product, Division, Modulus)
	fmt.Println("Math:", Sqrt2, Square4, CircleA, CircleC)
	fmt.Println("Time durations:", Second, Minute, Hour, Day)
	fmt.Println("File permissions:", Read, Write, Execute, All)
	fmt.Println("HTTP status codes:", StatusOK, StatusNotFound, StatusError, StatusCreated)
	fmt.Println("ArraySize:", ArraySize)
	fmt.Println("BufferSize:", BufferSize)

	// Using constants in code
	arr := [ArraySize]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", arr)

	buffer := make([]byte, BufferSize)
	fmt.Println("Buffer length:", len(buffer))

	// Using time duration constants
	fmt.Println("One hour in seconds:", Hour)

	// Using math constants
	fmt.Println("Area of circle with r=2:", CircleA)

	// Using constants in switch
	switch StatusOK {
	case 200:
		fmt.Println("Status is OK")
	}

	// Using constants in bitwise operations
	perm := Read | Write
	fmt.Println("Permissions:", perm)

	// Using constants in function calls
	fmt.Println("Sin(Pi):", math.Sin(Pi))

	// Using constants in type conversion
	var x float64 = float64(UntypedInt)
	fmt.Println("Converted:", x)

	// Using constants in string concatenation
	fmt.Println(Greeting + " Welcome!")

	// Using constants in boolean expressions
	if IsGoFun && !IsHard {
		fmt.Println("Go is fun and not hard!")
	}

	// Using constants in array/slice length
	slice := make([]int, ArraySize)
	fmt.Println("Slice length:", len(slice))

	// Using constants in for loop
	for i := 0; i < ArraySize; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// Using constants in map initialization
	m := make(map[string]int, BufferSize)
	fmt.Println("Map capacity (hint):", BufferSize)

	// Using constants in time.Sleep
	time.Sleep(time.Millisecond * 10)
	fmt.Println("Slept for 10ms (not a constant, but uses time package)")
}

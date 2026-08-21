package main

import (
	"fmt"
	"strings"
	"time"
)

// 1. Simple add function
func add(a, b int) int {
	return a + b
}

// 2. Subtract function
func subtract(a, b int) int {
	return a - b
}

// 3. Multiply function
func multiply(a, b int) int {
	return a * b
}

// 4. Divide function
func divide(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

// 5. Function with no parameters
func hello() {
	fmt.Println("Hello, Go!")
}

// 6. Function with no return value
func printSum(a, b int) {
	fmt.Println("Sum:", a+b)
}

// 7. Function with multiple return values
func swap(a, b string) (string, string) {
	return b, a
}

// 8. Named return values
func minMax(a, b int) (min, max int) {
	if a < b {
		min, max = a, b
	} else {
		min, max = b, a
	}
	return
}

// 9. Variadic function
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// 10. Function as value
func apply(f func(int, int) int, a, b int) int {
	return f(a, b)
}

// 11. Anonymous function
var double = func(x int) int {
	return x * 2
}

// 12. Closure
func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// 13. Recursive function
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// 14. Function with pointer parameter
func increment(a *int) {
	*a++
}

// 15. Function returning pointer
func newInt(val int) *int {
	return &val
}

// 16. Function with slice parameter
func printSlice(s []int) {
	fmt.Println("Slice:", s)
}

// 17. Function with map parameter
func printMap(m map[string]int) {
	fmt.Println("Map:", m)
}

// 18. Function with struct parameter
type Point struct{ X, Y int }

func printPoint(p Point) {
	fmt.Println("Point:", p)
}

// 19. Function returning struct
func newPoint(x, y int) Point {
	return Point{x, y}
}

// 20. Function with interface parameter
func describe(i interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", i, i)
}

// 21. Function with error return
func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// 22. Function with defer
func withDefer() {
	defer fmt.Println("Deferred!")
	fmt.Println("Function body")
}

// 23. Function with panic/recover
func safeCall(f func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()
	f()
}

// 24. Method on struct
type Counter struct{ n int }

func (c *Counter) Inc() { c.n++ }

// 25. Method with value receiver
func (c Counter) Value() int { return c.n }

// 26. Method with pointer receiver
func (c *Counter) Add(x int) { c.n += x }

// 27. Method chaining
func (c *Counter) ChainAdd(x int) *Counter {
	c.n += x
	return c
}

// 28. Function with bool parameter
func printBool(b bool) {
	fmt.Println("Bool:", b)
}

// 29. Function with float parameter
func printFloat(f float64) {
	fmt.Println("Float:", f)
}

// 30. Function with rune parameter
func printRune(r rune) {
	fmt.Printf("Rune: %c\n", r)
}

// 31. Function with byte parameter
func printByte(b byte) {
	fmt.Printf("Byte: %c\n", b)
}

// 32. Function with array parameter
func printArray(a [3]int) {
	fmt.Println("Array:", a)
}

// 33. Function with channel parameter
func sendInt(ch chan<- int, v int) {
	ch <- v
}

// 34. Function with channel return
func makeChan() chan int {
	return make(chan int)
}

// 35. Function with select (Added timeout to prevent deadlock)
func selectExample(ch1, ch2 <-chan int) int {
	select {
	case v := <-ch1:
		return v
	case v := <-ch2:
		return v
	case <-time.After(time.Second): // Added timeout
		return 0
	}
}

// 36. Function with go routine
func runAsync(f func()) {
	go f()
}

// 37. Function with string parameter
func shout(s string) {
	fmt.Println(strings.ToUpper(s))
}

// 38. Function with string return
func greet(name string) string {
	return "Hello, " + name
}

// 39. Function with multiple types
func printAll(a int, b string, c float64) {
	fmt.Println(a, b, c)
}

// 40. Function with default argument (simulate)
func greetDefault(name ...string) string {
	if len(name) == 0 {
		return "Hello, Guest"
	}
	return "Hello, " + name[0]
}

// 41. Function with callback
func doTwice(f func()) {
	f()
	f()
}

// 42. Function with function return
func makeAdder(x int) func(int) int {
	return func(y int) int { return x + y }
}

// 43. Function with type assertion
func assertInt(i interface{}) int {
	if v, ok := i.(int); ok {
		return v
	}
	return 0
}

// 44. Function with type switch
func typeSwitch(i interface{}) string {
	switch v := i.(type) {
	case int:
		return fmt.Sprintf("int: %d", v)
	case string:
		return "string: " + v
	default:
		return "unknown"
	}
}

// 45. Function with variadic string
func join(sep string, vals ...string) string {
	return strings.Join(vals, sep)
}

// 46. Function with variadic interface
func printAny(vals ...interface{}) {
	for _, v := range vals {
		fmt.Println(v)
	}
}

// 47. Function with named type
type Age int

func printAge(a Age) {
	fmt.Println("Age:", a)
}

// 48. Function with const parameter (simulate)
const Pi = 3.14

func areaCircle(r float64) float64 {
	return Pi * r * r
}

// 49. Function with global variable
var global int

func setGlobal(v int) {
	global = v
}

// 50. Function with local variable
func localVar() int {
	x := 42
	return x
}

// 51. Function with shadowed variable
func shadow(x int) int {
	x = x + 1
	return x
}

// 52. Function with for loop
func sumToN(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}
	return total
}

// 53. Function with if statement
func isEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

// 54. Function with switch statement
func dayName(n int) string {
	switch n {
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	default:
		return "Other"
	}
}

// 55. Function with label/goto
func skipToEnd() {
	goto End
	// fmt.Println("Skipped")
End:
	fmt.Println("End")
}

// 56. Function with defer and named return
func withDeferReturn() (result int) {
	defer func() { result++ }()
	result = 1
	return
}

// 57. Function with recover
func safeDivideRecover(a, b int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			result = 0
		}
	}()
	return a / b
}

// 58. Function with interface return
func returnInterface(a int) interface{} {
	return a
}

// 59. Function with empty interface parameter
func printInterface(i interface{}) {
	fmt.Println(i)
}

// 60. Function with slice return
func makeSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = i
	}
	return s
}

// 61. Function with map return
func makeMap(n int) map[int]int {
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		m[i] = i * i
	}
	return m
}

// 62. Function with struct return
func makeStruct(x, y int) Point {
	return Point{x, y}
}

// 63. Function with pointer return
func makePointer(x int) *int {
	return &x
}

// 64. Function with bool return
func isPositive(n int) bool {
	return n > 0
}

// 65. Function with float return
func half(n int) float64 {
	return float64(n) / 2
}

// 66. Function with rune return (Added empty string check)
func firstRune(s string) rune {
	if s == "" {
		return 0 // Return 0 for empty string
	}
	return []rune(s)[0]
}

// 67. Function with byte return (Added empty string check)
func firstByte(s string) byte {
	if s == "" {
		return 0 // Return 0 for empty string
	}
	return s[0]
}

// 68. Function with error return
func alwaysError() error {
	return fmt.Errorf("always error")
}

// 69. Function with nil return
func returnNil() error {
	return nil
}

// 70. Function with closure capturing variable
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// 71. Function with recursion (fibonacci)
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// 72. Function with multiple closures
func makeFuncs() (func(), func()) {
	a := 0
	return func() { a++ }, func() { fmt.Println(a) }
}

// 73. Function with defer in loop
func deferLoop() {
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
}

// 74. Function with select and channels
func selectChan(ch1, ch2 <-chan int) int {
	select {
	case v := <-ch1:
		return v
	case v := <-ch2:
		return v
	}
}

// 75. Function with buffered channel
func bufferedChan() chan int {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	return ch
}

// 76. Function with unbuffered channel
func unbufferedChan() chan int {
	ch := make(chan int)
	go func() { ch <- 1 }()
	return ch
}

// 77. Function with close channel
func closeChan(ch chan int) {
	close(ch)
}

// 78. Function with range over channel
func rangeChan(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

// 79. Function with send-only channel
func sendOnly(ch chan<- int, v int) {
	ch <- v
}

// 80. Function with receive-only channel
func recvOnly(ch <-chan int) int {
	return <-ch
}

// 81. Function with interface slice
func printInterfaces(vals []interface{}) {
	for _, v := range vals {
		fmt.Println(v)
	}
}

// 82. Function with type conversion
func toString(n int) string {
	return fmt.Sprintf("%d", n)
}

// 83. Function with string to int conversion (Added error handling)
func toInt(s string) (int, error) {
	var n int
	_, err := fmt.Sscanf(s, "%d", &n)
	if err != nil {
		return 0, fmt.Errorf("invalid integer: %v", err)
	}
	return n, nil
}

// 84. Function with string split
func splitWords(s string) []string {
	return strings.Fields(s)
}

// 85. Function with string join
func joinWords(words []string) string {
	return strings.Join(words, " ")
}

// 86. Function with string replace
func replaceGo(s string) string {
	return strings.ReplaceAll(s, "go", "GO")
}

// 87. Function with string contains
func containsGo(s string) bool {
	return strings.Contains(s, "go")
}

// 88. Function with string prefix
func hasPrefixGo(s string) bool {
	return strings.HasPrefix(s, "go")
}

// 89. Function with string suffix
func hasSuffixGo(s string) bool {
	return strings.HasSuffix(s, "go")
}

// 90. Function with string repeat
func repeatGo(n int) string {
	return strings.Repeat("go", n)
}

// 91. Function with string trim
func trimSpace(s string) string {
	return strings.TrimSpace(s)
}

// 92. Function with string to upper
func toUpper(s string) string {
	return strings.ToUpper(s)
}

// 93. Function with string to lower
func toLower(s string) string {
	return strings.ToLower(s)
}

// 94. Function with string length
func strLen(s string) int {
	return len(s)
}

// 95. Function with rune count
func runeCount(s string) int {
	return len([]rune(s))
}

// 96. Function with byte slice
func toBytes(s string) []byte {
	return []byte(s)
}

// 97. Function with runes slice
func toRunes(s string) []rune {
	return []rune(s)
}

// 98. Function with string builder
func buildString(parts ...string) string {
	var b strings.Builder
	for _, p := range parts {
		b.WriteString(p)
	}
	return b.String()
}

// 99. Function with defer and recover
func safeRun(f func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in safeRun")
		}
	}()
	f()
}

// 100. Function with main logic (Uncommented in main)
func mainLogic() {
	fmt.Println("This is the main logic function.")
}

func FunctionsUseMain() {
	fmt.Println("1+2 =", add(1, 2))
	fmt.Println("3-1 =", subtract(3, 1))
	fmt.Println("2*3 =", multiply(2, 3))
	fmt.Println("6/2 =", divide(6, 2))
	hello()
	printSum(4, 5)
	a, b := swap("foo", "bar")
	fmt.Println("swap:", a, b)
	min, max := minMax(3, 7)
	fmt.Println("min:", min, "max:", max)
	fmt.Println("sum:", sum(1, 2, 3, 4))
	fmt.Println("apply multiply:", apply(multiply, 2, 3))
	fmt.Println("double 5:", double(5))
	triple := makeMultiplier(3)
	fmt.Println("triple 4:", triple(4))
	fmt.Println("factorial 5:", factorial(5))
	x := 10
	increment(&x)
	fmt.Println("incremented:", x)
	fmt.Println("newInt:", *newInt(7))
	printSlice([]int{1, 2, 3})
	printMap(map[string]int{"a": 1})
	printPoint(Point{1, 2})
	fmt.Println("newPoint:", newPoint(3, 4))
	describe(42)
	res, err := safeDivide(10, 0)
	fmt.Println("safeDivide:", res, err)
	withDefer()
	safeCall(func() { panic("fail!") })
	c := Counter{}
	c.Inc()
	fmt.Println("Counter value:", c.Value())
	c.Add(5)
	fmt.Println("Counter value after Add:", c.Value())
	c.ChainAdd(10)
	fmt.Println("Counter value after ChainAdd:", c.Value())
	printBool(true)
	printFloat(3.14)
	printRune('G')
	printByte('B')
	printArray([3]int{1, 2, 3})
	ch := make(chan int, 1)
	sendInt(ch, 99)
	fmt.Println("chan val:", <-ch)
	ch2 := makeChan()
	go func() { ch2 <- 42 }()
	fmt.Println("makeChan val:", <-ch2)
	// Create fresh channels for selectExample to avoid deadlock
	ch3, ch4 := make(chan int), make(chan int)
	go func() { ch3 <- 100 }()
	fmt.Println("selectExample:", selectExample(ch3, ch4))
	runAsync(func() { fmt.Println("async run") })
	time.Sleep(time.Millisecond * 100) // Allow async to print
	shout("hello")
	fmt.Println(greet("Go"))
	printAll(1, "two", 3.0)
	fmt.Println(greetDefault())
	doTwice(func() { fmt.Print("twice ") })
	fmt.Println()
	adder := makeAdder(10)
	fmt.Println("adder(5):", adder(5))
	fmt.Println("assertInt(7):", assertInt(7))
	fmt.Println("typeSwitch(42):", typeSwitch(42))
	fmt.Println("join:", join("-", "a", "b", "c"))
	printAny(1, "two", 3.0)
	printAge(21)
	fmt.Println("areaCircle(2):", areaCircle(2))
	setGlobal(100)
	fmt.Println("global:", global)
	fmt.Println("localVar:", localVar())
	fmt.Println("shadow(5):", shadow(5)) // Consider renaming to incrementCopy
	fmt.Println("sumToN(5):", sumToN(5))
	fmt.Println("isEven(4):", isEven(4))
	fmt.Println("dayName(2):", dayName(2))
	skipToEnd()
	fmt.Println("withDeferReturn:", withDeferReturn())
	fmt.Println("safeDivideRecover(4,0):", safeDivideRecover(4, 0))
	fmt.Println("returnInterface(5):", returnInterface(5))
	printInterface("interface value")
	fmt.Println("makeSlice(3):", makeSlice(3))
	fmt.Println("makeMap(3):", makeMap(3))
	fmt.Println("makeStruct(1,2):", makeStruct(1, 2))
	fmt.Println("makePointer(7):", *makePointer(7))
	fmt.Println("isPositive(3):", isPositive(3))
	fmt.Println("half(5):", half(5))
	fmt.Printf("firstRune('go'):%c\n", firstRune("go"))
	fmt.Printf("firstByte('go'):%c\n", firstByte("go"))
	fmt.Printf("firstRune(''):%c\n", firstRune("")) // Test empty string
	fmt.Printf("firstByte(''):%c\n", firstByte("")) // Test empty string
	fmt.Println("alwaysError():", alwaysError())
	fmt.Println("returnNil():", returnNil())
	counter := makeCounter()
	fmt.Println("counter:", counter(), counter())
	fmt.Println("fib(6):", fib(6))
	f1, f2 := makeFuncs()
	f1()
	f2()
	deferLoop()
	ch5 := bufferedChan()
	fmt.Println("bufferedChan:", <-ch5, <-ch5)
	ch6 := unbufferedChan()
	fmt.Println("unbufferedChan:", <-ch6)
	ch7 := make(chan int, 2)
	sendOnly(ch7, 123)
	fmt.Println("recvOnly:", recvOnly(ch7))
	// Demonstrate closeChan and rangeChan
	ch8 := make(chan int, 2)
	go func() {
		ch8 <- 1
		ch8 <- 2
		close(ch8) // Close channel for rangeChan
	}()
	rangeChan(ch8)
	printInterfaces([]interface{}{1, "two", 3.0})
	fmt.Println("toString(42):", toString(42))
	n, err := toInt("42")
	fmt.Println("toInt('42'):", n, err)
	n, err = toInt("invalid")
	fmt.Println("toInt('invalid'):", n, err)
	fmt.Println("splitWords('go is fun'):", splitWords("go is fun"))
	fmt.Println("joinWords:", joinWords([]string{"go", "is", "fun"}))
	fmt.Println("replaceGo('go go'):", replaceGo("go go"))
	fmt.Println("containsGo('golang'):", containsGo("golang"))
	fmt.Println("hasPrefixGo('golang'):", hasPrefixGo("golang"))
	fmt.Println("hasSuffixGo('long'):", hasSuffixGo("long"))
	fmt.Println("repeatGo(3):", repeatGo(3))
	fmt.Println("trimSpace('  go  '):", trimSpace("  go  "))
	fmt.Println("toUpper('go'):", toUpper("go"))
	fmt.Println("toLower('GO'):", toLower("GO"))
	fmt.Println("strLen('go'):", strLen("go"))
	fmt.Println("runeCount('你好'):", runeCount("你好"))
	fmt.Println("toBytes('go'):", toBytes("go"))
	fmt.Println("toRunes('go'):", toRunes("go"))
	fmt.Println("buildString('a','b','c'):", buildString("a", "b", "c"))
	safeRun(func() { panic("fail!") })
	mainLogic() // Uncommented to include
}

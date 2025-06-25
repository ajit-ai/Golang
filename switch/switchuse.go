package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// 1. Basic integer switch
	x := 2
	switch x {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other")
	}

	// 2. String switch
	lang := "go"
	switch lang {
	case "python":
		fmt.Println("Python")
	case "go":
		fmt.Println("Go")
	default:
		fmt.Println("Other language")
	}

	// 3. Multiple values in a case
	day := time.Now().Weekday()
	switch day {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	// 4. Switch without an expression
	h := time.Now().Hour()
	switch {
	case h < 12:
		fmt.Println("Morning")
	case h < 18:
		fmt.Println("Afternoon")
	default:
		fmt.Println("Evening")
	}

	// 5. Type switch
	var i interface{} = 3.14
	switch v := i.(type) {
	case int:
		fmt.Println("int", v)
	case float64:
		fmt.Println("float64", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("unknown type")
	}

	// 6. Switch with fallthrough
	num := 1
	switch num {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two or after one")
	default:
		fmt.Println("Other")
	}

	// 7. Switch on boolean
	b := true
	switch b {
	case true:
		fmt.Println("True")
	case false:
		fmt.Println("False")
	}

	// 8. Switch on rune
	ch := 'a'
	switch ch {
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Println("Vowel")
	default:
		fmt.Println("Consonant")
	}

	// 9. Switch on string length
	s := "hello"
	switch l := len(s); {
	case l == 0:
		fmt.Println("Empty")
	case l < 5:
		fmt.Println("Short")
	default:
		fmt.Println("Long")
	}

	// 10. Switch on error value
	var err error
	switch err {
	case nil:
		fmt.Println("No error")
	default:
		fmt.Println("Error occurred")
	}

	// 11. Switch on map key existence
	m := map[string]int{"a": 1}
	switch _, ok := m["b"]; ok {
	case true:
		fmt.Println("Key exists")
	default:
		fmt.Println("Key does not exist")
	}

	// 12. Switch on slice length
	sl := []int{1, 2, 3}
	switch len(sl) {
	case 0:
		fmt.Println("Empty slice")
	case 1, 2:
		fmt.Println("Small slice")
	default:
		fmt.Println("Large slice")
	}

	// 13. Switch on modulo
	n := 7
	switch n % 2 {
	case 0:
		fmt.Println("Even")
	case 1:
		fmt.Println("Odd")
	}

	// 14. Switch on float comparison
	f := 3.5
	switch {
	case f < 0:
		fmt.Println("Negative")
	case f == 0:
		fmt.Println("Zero")
	default:
		fmt.Println("Positive")
	}

	// 15. Switch on struct field
	type Person struct{ age int }
	p := Person{age: 20}
	switch {
	case p.age < 13:
		fmt.Println("Child")
	case p.age < 20:
		fmt.Println("Teenager")
	default:
		fmt.Println("Adult")
	}

	// 16. Switch on time.Month
	month := time.Now().Month()
	switch month {
	case time.December, time.January, time.February:
		fmt.Println("Winter")
	case time.March, time.April, time.May:
		fmt.Println("Spring")
	case time.June, time.July, time.August:
		fmt.Println("Summer")
	default:
		fmt.Println("Autumn")
	}

	// 17. Switch on nil pointer
	var ptr *int
	switch ptr {
	case nil:
		fmt.Println("Pointer is nil")
	default:
		fmt.Println("Pointer is not nil")
	}

	// 18. Switch on string prefix
	str := "golang"
	switch {
	case strings.HasPrefix(str, "go"):
		fmt.Println("Starts with go")
	case strings.HasPrefix(str, "py"):
		fmt.Println("Starts with py")
	default:
		fmt.Println("Other prefix")
	}

	// 19. Switch on string suffix
	switch {
	case strings.HasSuffix(str, "lang"):
		fmt.Println("Ends with lang")
	default:
		fmt.Println("Other suffix")
	}

	// 20. Switch on string contains
	switch {
	case strings.Contains(str, "la"):
		fmt.Println("Contains la")
	default:
		fmt.Println("Does not contain la")
	}

	// 21. Switch on interface type (custom)
	var v interface{} = []int{1, 2, 3}
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case []int:
		fmt.Println("slice of int")
	default:
		fmt.Println("other type")
	}

	// 22. Switch on HTTP status code
	status := 404
	switch status {
	case 200:
		fmt.Println("OK")
	case 404:
		fmt.Println("Not Found")
	case 500:
		fmt.Println("Server Error")
	default:
		fmt.Println("Other status")
	}

	// 23. Switch on day of week
	switch day := time.Now().Weekday(); day {
	case time.Monday:
		fmt.Println("Monday")
	case time.Tuesday:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Other day")
	}

	// 24. Switch on battery level
	battery := 15
	switch {
	case battery < 10:
		fmt.Println("Low")
	case battery < 50:
		fmt.Println("Medium")
	default:
		fmt.Println("High")
	}

	// 25. Switch on grade
	grade := "B"
	switch grade {
	case "A":
		fmt.Println("Excellent")
	case "B":
		fmt.Println("Good")
	case "C":
		fmt.Println("Average")
	default:
		fmt.Println("Poor")
	}

	// 26. Switch on leap year
	year := 2024
	switch {
	case year%400 == 0:
		fmt.Println("Leap year")
	case year%100 == 0:
		fmt.Println("Not leap year")
	case year%4 == 0:
		fmt.Println("Leap year")
	default:
		fmt.Println("Not leap year")
	}

	// 27. Switch on triangle type
	a, b2, c := 3, 4, 5
	switch {
	case a == b2 && b2 == c:
		fmt.Println("Equilateral")
	case a == b2 || b2 == c || a == c:
		fmt.Println("Isosceles")
	default:
		fmt.Println("Scalene")
	}

	// 28. Switch on password length
	pass := "golang"
	switch l := len(pass); {
	case l < 6:
		fmt.Println("Weak")
	case l < 10:
		fmt.Println("Medium")
	default:
		fmt.Println("Strong")
	}

	// 29. Switch on BMI
	bmi := 22.5
	switch {
	case bmi < 18.5:
		fmt.Println("Underweight")
	case bmi < 25:
		fmt.Println("Normal")
	case bmi < 30:
		fmt.Println("Overweight")
	default:
		fmt.Println("Obese")
	}

	// 30. Switch on currency
	currency := "USD"
	switch currency {
	case "USD":
		fmt.Println("$")
	case "EUR":
		fmt.Println("€")
	case "JPY":
		fmt.Println("¥")
	default:
		fmt.Println("Other currency")
	}

	// 31. Switch on voting eligibility
	age := 17
	switch {
	case age < 16:
		fmt.Println("Too young")
	case age < 18:
		fmt.Println("Almost eligible")
	default:
		fmt.Println("Eligible")
	}

	// 32. Switch on water state
	temp := 100
	switch {
	case temp < 0:
		fmt.Println("Solid")
	case temp < 100:
		fmt.Println("Liquid")
	default:
		fmt.Println("Gas")
	}

	// 33. Switch on quadrant
	x2, y2 := -1, 1
	switch {
	case x2 > 0 && y2 > 0:
		fmt.Println("Quadrant I")
	case x2 < 0 && y2 > 0:
		fmt.Println("Quadrant II")
	case x2 < 0 && y2 < 0:
		fmt.Println("Quadrant III")
	case x2 > 0 && y2 < 0:
		fmt.Println("Quadrant IV")
	default:
		fmt.Println("On axis")
	}

	// 34. Switch on character type
	ch2 := 'A'
	switch {
	case ch2 >= 'A' && ch2 <= 'Z':
		fmt.Println("Uppercase")
	case ch2 >= 'a' && ch2 <= 'z':
		fmt.Println("Lowercase")
	case ch2 >= '0' && ch2 <= '9':
		fmt.Println("Digit")
	default:
		fmt.Println("Other character")
	}

	// 35. Switch on shopping discount
	price := 120
	switch {
	case price > 200:
		fmt.Println("20% discount")
	case price > 100:
		fmt.Println("10% discount")
	default:
		fmt.Println("No discount")
	}

	// 36. Switch on exam result
	marks := 75
	switch {
	case marks >= 90:
		fmt.Println("Excellent")
	case marks >= 75:
		fmt.Println("Good")
	case marks >= 50:
		fmt.Println("Pass")
	default:
		fmt.Println("Fail")
	}

	// 37. Switch on internet speed
	speed := 50
	switch {
	case speed < 10:
		fmt.Println("Slow")
	case speed < 50:
		fmt.Println("Average")
	default:
		fmt.Println("Fast")
	}

	// 38. Switch on rectangle or square
	w, h := 5, 5
	switch {
	case w == h:
		fmt.Println("Square")
	default:
		fmt.Println("Rectangle")
	}

	// 39. Switch on driving eligibility
	age2 := 15
	switch {
	case age2 >= 18:
		fmt.Println("Can drive")
	case age2 >= 16:
		fmt.Println("Can drive with supervision")
	default:
		fmt.Println("Cannot drive")
	}

	// 40. Switch on string case
	s2 := "HELLO"
	switch {
	case s2 == strings.ToUpper(s2):
		fmt.Println("Uppercase")
	case s2 == strings.ToLower(s2):
		fmt.Println("Lowercase")
	default:
		fmt.Println("Mixed case")
	}

	// 41. Switch on number range
	n2 := 25
	switch {
	case n2 < 10:
		fmt.Println("Less than 10")
	case n2 < 20:
		fmt.Println("10-19")
	case n2 < 30:
		fmt.Println("20-29")
	default:
		fmt.Println("30 or more")
	}

	// 42. Switch on multiple of 2, 3, or 5
	n3 := 15
	switch {
	case n3%2 == 0:
		fmt.Println("Multiple of 2")
	case n3%3 == 0:
		fmt.Println("Multiple of 3")
	case n3%5 == 0:
		fmt.Println("Multiple of 5")
	default:
		fmt.Println("Not a multiple")
	}

	// 43. Switch on string equality ignoring case
	s3 := "Go"
	switch {
	case strings.EqualFold(s3, "go"):
		fmt.Println("Equal (ignore case)")
	default:
		fmt.Println("Not equal")
	}

	// 44. Switch on nil interface
	var iface interface{}
	switch iface {
	case nil:
		fmt.Println("Interface is nil")
	default:
		fmt.Println("Interface is not nil")
	}

	// 45. Switch on slice length
	sl2 := []int{1, 2, 3}
	switch l := len(sl2); {
	case l == 0:
		fmt.Println("Empty slice")
	case l < 5:
		fmt.Println("Small slice")
	default:
		fmt.Println("Large slice")
	}

	// 46. Switch on map key existence (again)
	m2 := map[string]int{"a": 1}
	switch _, ok := m2["b"]; ok {
	case true:
		fmt.Println("Key b exists")
	default:
		fmt.Println("Key b does not exist")
	}

	// 47. Switch on interface type (again)
	var v2 interface{} = 123
	switch v2.(type) {
	case string:
		fmt.Println("String")
	case int:
		fmt.Println("Int")
	default:
		fmt.Println("Other type")
	}

	// 48. Switch on empty string
	s4 := ""
	switch {
	case s4 == "":
		fmt.Println("Empty string")
	case len(s4) < 5:
		fmt.Println("Short string")
	default:
		fmt.Println("Long string")
	}

	// 49. Switch on pointer nil
	var p2 *int
	switch {
	case p2 == nil:
		fmt.Println("Pointer is nil")
	default:
		fmt.Println("Pointer is not nil")
	}

	// 50. Switch on function result
	switch res := strings.Contains("golang", "go"); res {
	case true:
		fmt.Println("Contains 'go'")
	default:
		fmt.Println("Does not contain 'go'")
	}
}

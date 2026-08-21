package main

import (
	"fmt"
	"strings"
)

func ElseIfUseMain() {
	// 1. Basic number comparison
	x := 10
	if x < 0 {
		fmt.Println("Negative")
	} else if x == 0 {
		fmt.Println("Zero")
	} else {
		fmt.Println("Positive")
	}

	// 2. String comparison
	s := "go"
	if s == "python" {
		fmt.Println("Python")
	} else if s == "java" {
		fmt.Println("Java")
	} else if s == "go" {
		fmt.Println("Go")
	} else {
		fmt.Println("Other")
	}

	// 3. Multiple else if
	y := 15
	if y < 10 {
		fmt.Println("Less than 10")
	} else if y < 20 {
		fmt.Println("Between 10 and 19")
	} else if y < 30 {
		fmt.Println("Between 20 and 29")
	} else {
		fmt.Println("30 or more")
	}

	// 4. Even or odd
	n := 7
	if n%2 == 0 {
		fmt.Println("Even")
	} else if n%2 == 1 {
		fmt.Println("Odd")
	}

	// 5. Grade system
	score := 85
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else if score >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}

	// 6. Age group
	age := 25
	if age < 13 {
		fmt.Println("Child")
	} else if age < 20 {
		fmt.Println("Teenager")
	} else if age < 65 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Senior")
	}

	// 7. String length
	word := "golang"
	if len(word) == 0 {
		fmt.Println("Empty")
	} else if len(word) < 5 {
		fmt.Println("Short")
	} else if len(word) < 10 {
		fmt.Println("Medium")
	} else {
		fmt.Println("Long")
	}

	// 8. Temperature check
	temp := 30
	if temp < 0 {
		fmt.Println("Freezing")
	} else if temp < 20 {
		fmt.Println("Cold")
	} else if temp < 30 {
		fmt.Println("Warm")
	} else {
		fmt.Println("Hot")
	}

	// 9. Nested else if
	a, b := 5, 10
	if a > b {
		fmt.Println("a > b")
	} else if a < b {
		fmt.Println("a < b")
	} else {
		fmt.Println("a == b")
	}

	// 10. Check divisibility
	num := 12
	if num%3 == 0 && num%4 == 0 {
		fmt.Println("Divisible by 3 and 4")
	} else if num%3 == 0 {
		fmt.Println("Divisible by 3")
	} else if num%4 == 0 {
		fmt.Println("Divisible by 4")
	} else {
		fmt.Println("Not divisible by 3 or 4")
	}

	// 11. Traffic light
	light := "yellow"
	if light == "red" {
		fmt.Println("Stop")
	} else if light == "yellow" {
		fmt.Println("Ready")
	} else if light == "green" {
		fmt.Println("Go")
	} else {
		fmt.Println("Unknown")
	}

	// 12. File extension
	filename := "main.go"
	if strings.HasSuffix(filename, ".py") {
		fmt.Println("Python file")
	} else if strings.HasSuffix(filename, ".go") {
		fmt.Println("Go file")
	} else if strings.HasSuffix(filename, ".java") {
		fmt.Println("Java file")
	} else {
		fmt.Println("Other file")
	}

	// 13. Month days
	month := 2
	if month == 2 {
		fmt.Println("28 or 29 days")
	} else if month == 4 || month == 6 || month == 9 || month == 11 {
		fmt.Println("30 days")
	} else {
		fmt.Println("31 days")
	}

	// 14. Triangle type
	a1, b1, c1 := 3, 4, 5
	if a1 == b1 && b1 == c1 {
		fmt.Println("Equilateral")
	} else if a1 == b1 || b1 == c1 || a1 == c1 {
		fmt.Println("Isosceles")
	} else {
		fmt.Println("Scalene")
	}

	// 15. Leap year
	year := 2024
	if year%400 == 0 {
		fmt.Println("Leap year")
	} else if year%100 == 0 {
		fmt.Println("Not leap year")
	} else if year%4 == 0 {
		fmt.Println("Leap year")
	} else {
		fmt.Println("Not leap year")
	}

	// 16. Password strength
	pass := "golang123"
	if len(pass) < 6 {
		fmt.Println("Weak")
	} else if len(pass) < 10 {
		fmt.Println("Medium")
	} else {
		fmt.Println("Strong")
	}

	// 17. BMI category
	bmi := 22.5
	if bmi < 18.5 {
		fmt.Println("Underweight")
	} else if bmi < 25 {
		fmt.Println("Normal")
	} else if bmi < 30 {
		fmt.Println("Overweight")
	} else {
		fmt.Println("Obese")
	}

	// 18. Day of week
	day := 3
	if day == 1 {
		fmt.Println("Monday")
	} else if day == 2 {
		fmt.Println("Tuesday")
	} else if day == 3 {
		fmt.Println("Wednesday")
	} else if day == 4 {
		fmt.Println("Thursday")
	} else if day == 5 {
		fmt.Println("Friday")
	} else if day == 6 {
		fmt.Println("Saturday")
	} else if day == 7 {
		fmt.Println("Sunday")
	} else {
		fmt.Println("Invalid day")
	}

	// 19. HTTP status code
	status := 404
	if status == 200 {
		fmt.Println("OK")
	} else if status == 404 {
		fmt.Println("Not Found")
	} else if status == 500 {
		fmt.Println("Server Error")
	} else {
		fmt.Println("Other status")
	}

	// 20. Currency symbol
	currency := "USD"
	if currency == "USD" {
		fmt.Println("$")
	} else if currency == "EUR" {
		fmt.Println("€")
	} else if currency == "JPY" {
		fmt.Println("¥")
	} else {
		fmt.Println("Other currency")
	}

	// 21. Age for voting
	age2 := 17
	if age2 < 16 {
		fmt.Println("Too young")
	} else if age2 < 18 {
		fmt.Println("Almost eligible")
	} else {
		fmt.Println("Eligible")
	}

	// 22. Water state
	temp2 := 100
	if temp2 < 0 {
		fmt.Println("Solid")
	} else if temp2 < 100 {
		fmt.Println("Liquid")
	} else {
		fmt.Println("Gas")
	}

	// 23. Quadrant of point
	x2, y2 := -1, 1
	if x2 > 0 && y2 > 0 {
		fmt.Println("Quadrant I")
	} else if x2 < 0 && y2 > 0 {
		fmt.Println("Quadrant II")
	} else if x2 < 0 && y2 < 0 {
		fmt.Println("Quadrant III")
	} else if x2 > 0 && y2 < 0 {
		fmt.Println("Quadrant IV")
	} else {
		fmt.Println("On axis")
	}

	// 24. Character type
	ch := 'A'
	if ch >= 'A' && ch <= 'Z' {
		fmt.Println("Uppercase")
	} else if ch >= 'a' && ch <= 'z' {
		fmt.Println("Lowercase")
	} else if ch >= '0' && ch <= '9' {
		fmt.Println("Digit")
	} else {
		fmt.Println("Other character")
	}

	// 25. Shopping discount
	price := 120
	if price > 200 {
		fmt.Println("20% discount")
	} else if price > 100 {
		fmt.Println("10% discount")
	} else {
		fmt.Println("No discount")
	}

	// 26. Exam result
	marks := 75
	if marks >= 90 {
		fmt.Println("Excellent")
	} else if marks >= 75 {
		fmt.Println("Good")
	} else if marks >= 50 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}

	// 27. Internet speed
	speed := 50
	if speed < 10 {
		fmt.Println("Slow")
	} else if speed < 50 {
		fmt.Println("Average")
	} else {
		fmt.Println("Fast")
	}

	// 28. Battery level
	battery := 15
	if battery < 10 {
		fmt.Println("Low")
	} else if battery < 50 {
		fmt.Println("Medium")
	} else {
		fmt.Println("High")
	}

	// 29. Exam grade with +/-
	grade := 88
	if grade >= 90 {
		fmt.Println("A")
	} else if grade >= 80 {
		if grade >= 85 {
			fmt.Println("B+")
		} else {
			fmt.Println("B")
		}
	} else {
		fmt.Println("Below B")
	}

	// 30. Triangle validity
	a2, b2, c2 := 3, 4, 8
	if a2+b2 > c2 && a2+c2 > b2 && b2+c2 > a2 {
		fmt.Println("Valid triangle")
	} else {
		fmt.Println("Invalid triangle")
	}

	// 31. Season by month
	month2 := 7
	if month2 >= 3 && month2 <= 5 {
		fmt.Println("Spring")
	} else if month2 >= 6 && month2 <= 8 {
		fmt.Println("Summer")
	} else if month2 >= 9 && month2 <= 11 {
		fmt.Println("Autumn")
	} else {
		fmt.Println("Winter")
	}

	// 32. Password check
	pass2 := "123456"
	if pass2 == "" {
		fmt.Println("Empty")
	} else if len(pass2) < 6 {
		fmt.Println("Too short")
	} else {
		fmt.Println("OK")
	}

	// 33. File size
	size := 2048
	if size < 1024 {
		fmt.Println("Small")
	} else if size < 1048576 {
		fmt.Println("Medium")
	} else {
		fmt.Println("Large")
	}

	// 34. User role
	role := "admin"
	if role == "admin" {
		fmt.Println("Full access")
	} else if role == "editor" {
		fmt.Println("Edit access")
	} else if role == "viewer" {
		fmt.Println("Read only")
	} else {
		fmt.Println("Unknown role")
	}

	// 35. Exam pass/fail
	score2 := 49
	if score2 >= 50 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}

	// 36. Number sign
	num2 := -5
	if num2 > 0 {
		fmt.Println("Positive")
	} else if num2 < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Zero")
	}

	// 37. String prefix
	s2 := "golang"
	if strings.HasPrefix(s2, "go") {
		fmt.Println("Starts with go")
	} else if strings.HasPrefix(s2, "py") {
		fmt.Println("Starts with py")
	} else {
		fmt.Println("Other prefix")
	}

	// 38. String contains
	s3 := "hello world"
	if strings.Contains(s3, "go") {
		fmt.Println("Contains go")
	} else if strings.Contains(s3, "world") {
		fmt.Println("Contains world")
	} else {
		fmt.Println("No match")
	}

	// 39. Rectangle or square
	w, h := 5, 5
	if w == h {
		fmt.Println("Square")
	} else {
		fmt.Println("Rectangle")
	}

	// 40. Voting eligibility
	age3 := 20
	if age3 >= 18 {
		fmt.Println("Can vote")
	} else {
		fmt.Println("Cannot vote")
	}

	// 41. Driving eligibility
	age4 := 15
	if age4 >= 18 {
		fmt.Println("Can drive")
	} else if age4 >= 16 {
		fmt.Println("Can drive with supervision")
	} else {
		fmt.Println("Cannot drive")
	}

	// 42. String case
	s4 := "HELLO"
	if s4 == strings.ToUpper(s4) {
		fmt.Println("Uppercase")
	} else if s4 == strings.ToLower(s4) {
		fmt.Println("Lowercase")
	} else {
		fmt.Println("Mixed case")
	}

	// 43. Number range
	n2 := 25
	if n2 < 10 {
		fmt.Println("Less than 10")
	} else if n2 < 20 {
		fmt.Println("10-19")
	} else if n2 < 30 {
		fmt.Println("20-29")
	} else {
		fmt.Println("30 or more")
	}

	// 44. Multiple of 2, 3, or 5
	n3 := 15
	if n3%2 == 0 {
		fmt.Println("Multiple of 2")
	} else if n3%3 == 0 {
		fmt.Println("Multiple of 3")
	} else if n3%5 == 0 {
		fmt.Println("Multiple of 5")
	} else {
		fmt.Println("Not a multiple")
	}

	// 45. String equality ignoring case
	s5 := "Go"
	if strings.EqualFold(s5, "go") {
		fmt.Println("Equal (ignore case)")
	} else {
		fmt.Println("Not equal")
	}

	// 46. Check for nil pointer
	var ptr *int
	if ptr == nil {
		fmt.Println("Pointer is nil")
	} else {
		fmt.Println("Pointer is not nil")
	}

	// 47. Check slice length
	sl := []int{1, 2, 3}
	if len(sl) == 0 {
		fmt.Println("Empty slice")
	} else if len(sl) < 5 {
		fmt.Println("Small slice")
	} else {
		fmt.Println("Large slice")
	}

	// 48. Check map key existence
	m2 := map[string]int{"a": 1}
	if _, ok := m2["b"]; ok {
		fmt.Println("Key b exists")
	} else if _, ok := m2["a"]; ok {
		fmt.Println("Key a exists")
	} else {
		fmt.Println("No key found")
	}

	// 49. Check interface type
	var v interface{} = 123
	if _, ok := v.(string); ok {
		fmt.Println("String")
	} else if _, ok := v.(int); ok {
		fmt.Println("Int")
	} else {
		fmt.Println("Other type")
	}

	// 50. Check for empty string
	s6 := ""
	if s6 == "" {
		fmt.Println("Empty string")
	} else if len(s6) < 5 {
		fmt.Println("Short string")
	} else {
		fmt.Println("Long string")
	}
}

// main package demonstrates modern error handling:
// sentinel errors, custom error types, wrapping with %w, errors.Is/As
package main

import (
	"errors"
	"fmt"
)

// ErrNotFound is a sentinel error callers can compare with errors.Is
var ErrNotFound = errors.New("item not found")

// ErrForbidden is another sentinel
var ErrForbidden = errors.New("access denied")

// ValidationError carries structured context about what failed
type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed for %q: %s", e.Field, e.Msg)
}

// FindUser returns a wrapped sentinel error when the user is missing
func FindUser(id int) (string, error) {
	users := map[int]string{1: "ajit", 2: "priya"}
	name, ok := users[id]
	if !ok {
		return "", fmt.Errorf("find user %d: %w", id, ErrNotFound)
	}
	return name, nil
}

// ValidateAge returns a typed error callers can inspect with errors.As
func ValidateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Msg: "cannot be negative"}
	}
	if age > 150 {
		return fmt.Errorf("validate age %d: %w", age, &ValidationError{Field: "age", Msg: "unrealistic"})
	}
	return nil
}

// CheckPermission chains wrapping two levels deep
func CheckPermission(role string) error {
	if role != "admin" {
		return fmt.Errorf("check permission for %q: %w", role, ErrForbidden)
	}
	return nil
}

// SafeCall converts a panic into an error using defer/recover
func SafeCall(fn func() int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered from panic: %v", r)
		}
	}()
	return fn(), nil
}

// CustomErrorsMain demonstrates errors.Is and errors.As
func CustomErrorsMain() {
	_, err := FindUser(99)
	fmt.Println("is not found:", errors.Is(err, ErrNotFound))

	if err := ValidateAge(-5); err != nil {
		var ve *ValidationError
		if errors.As(err, &ve) {
			fmt.Println("invalid field:", ve.Field)
		}
	}

	fmt.Println("permission:", CheckPermission("guest"))
}

// RecoverMain demonstrates panic recovery
func RecoverMain() {
	value, err := SafeCall(func() int { return 42 })
	fmt.Println(value, err)

	value, err = SafeCall(func() int { panic("boom") })
	fmt.Println(value, err)
}

// main runs the demo entry points of this package
func main() {
	CustomErrorsMain()
	RecoverMain()
}

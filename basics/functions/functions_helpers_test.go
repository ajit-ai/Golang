package main

import "testing"

func TestPlus(t *testing.T) {
	if got := plus(1, 2); got != 3 {
		t.Errorf("plus(1, 2) = %d, want 3", got)
	}
	if got := plus(-1, 1); got != 0 {
		t.Errorf("plus(-1, 1) = %d, want 0", got)
	}
}

func TestPlusPlus(t *testing.T) {
	if got := plusPlus(1, 2, 3); got != 6 {
		t.Errorf("plusPlus(1, 2, 3) = %d, want 6", got)
	}
}

func TestFunctionsUseHelpers(t *testing.T) {
	tests := []struct {
		name string
		got  int
		want int
	}{
		{"add", add(2, 3), 5},
		{"subtract", subtract(5, 2), 3},
		{"multiply", multiply(3, 4), 12},
		{"divide", divide(10, 2), 5},
		{"sum variadic", sum(1, 2, 3, 4), 10},
		{"sum empty", sum(), 0},
		{"factorial", factorial(5), 120},
		{"factorial zero", factorial(0), 1},
		{"double closure", double(21), 42},
		{"makeMultiplier", makeMultiplier(3)(4), 12},
		{"sumToN", sumToN(4), 10},
		{"apply", apply(add, 2, 8), 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("%s: got %d, want %d", tt.name, tt.got, tt.want)
			}
		})
	}
}

func TestDivideByZeroReturnsZero(t *testing.T) {
	if got := divide(1, 0); got != 0 {
		t.Errorf("divide(1, 0) = %d, want 0", got)
	}
}

func TestMinMax(t *testing.T) {
	min, max := minMax(3, 7)
	if min != 3 || max != 7 {
		t.Errorf("minMax(3, 7) = %d, %d; want 3, 7", min, max)
	}
	min, max = minMax(9, 1)
	if min != 1 || max != 9 {
		t.Errorf("minMax(9, 1) = %d, %d; want 1, 9", min, max)
	}
}

func TestSwap(t *testing.T) {
	a, b := swap("x", "y")
	if a != "y" || b != "x" {
		t.Errorf("swap = %q, %q; want y, x", a, b)
	}
}

func TestSafeDivide(t *testing.T) {
	if got, err := safeDivide(10, 2); err != nil || got != 5 {
		t.Errorf("safeDivide(10, 2) = %d, %v; want 5, nil", got, err)
	}
	if _, err := safeDivide(1, 0); err == nil {
		t.Error("safeDivide by zero should return an error")
	}
}

func TestIsEvenAndDayName(t *testing.T) {
	if !isEven(4) || isEven(5) {
		t.Error("isEven returned wrong results")
	}
	if got := dayName(1); got != "Monday" {
		t.Errorf("dayName(1) = %q, want Monday", got)
	}
}

func TestFunctionsMainSmoke(t *testing.T) {
	FunctionsMain()
}

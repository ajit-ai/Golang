package main

import (
	"reflect"
	"testing"
)

func TestBasicPureFunctions(t *testing.T) {
	tests := []struct {
		name string
		got  interface{}
		want interface{}
	}{
		{"sum", sum(2, 3), 5},
		{"isEven true", isEven(4), true},
		{"isEven false", isEven(7), false},
		{"factorial", factorial(5), 120},
		{"factorial 0", factorial(0), 1},
		{"reverseString", reverseString("hello"), "olleh"},
		{"reverseString empty", reverseString(""), ""},
		{"isPalindrome true", isPalindrome("racecar"), true},
		{"isPalindrome false", isPalindrome("hello"), false},
		{"fibonacci", fibonacci(10), 55},
		{"isPrime true", isPrime(13), true},
		{"isPrime false", isPrime(15), false},
		{"countVowels", countVowels("golang"), 2},
		{"areAnagrams true", areAnagrams("listen", "silent"), true},
		{"areAnagrams false", areAnagrams("go", "og!"), false},
		{"gcd", gcd(48, 18), 6},
		{"isPowerOfTwo true", isPowerOfTwo(64), true},
		{"isPowerOfTwo false", isPowerOfTwo(63), false},
		{"binarySearch found", binarySearch([]int{1, 3, 5, 7, 9}, 7), 3},
		{"binarySearch missing", binarySearch([]int{1, 3, 5}, 4), -1},
		{"atoi", atoi("12345"), 12345},
		{"isDigit true", isDigit('7'), true},
		{"isDigit false", isDigit('x'), false},
		{"isValidParenthesis true", isValidParenthesis("{[()]}"), true},
		{"isValidParenthesis false", isValidParenthesis("(]"), false},
		{"longestCommonPrefix", longestCommonPrefix([]string{"flower", "flow", "flight"}), "fl"},
		{"max", max(3, 9), 9},
		{"intToRoman", intToRoman(1994), "MCMXCIV"},
		{"isNumberPalindrome true", isNumberPalindrome(121), true},
		{"isNumberPalindrome false", isNumberPalindrome(-121), false},
		{"numberToBinary", numberToBinary(5), "101"},
		{"numberToBinary zero", numberToBinary(0), "0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("%s: got %v, want %v", tt.name, tt.got, tt.want)
			}
		})
	}
}

func TestMaxElement(t *testing.T) {
	if got := maxElement([]int{3, 9, 2}); got != 9 {
		t.Errorf("maxElement = %d, want 9", got)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	got := removeDuplicates([]int{1, 2, 2, 3, 1})
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("removeDuplicates = %v, want %v", got, want)
	}
}

func TestMergeSortedArrays(t *testing.T) {
	got := mergeSortedArrays([]int{1, 3, 5}, []int{2, 4, 6})
	want := []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("mergeSortedArrays = %v, want %v", got, want)
	}
}

func TestFindMissingNumber(t *testing.T) {
	if got := findMissingNumber([]int{1, 2, 4, 5}, 5); got != 3 {
		t.Errorf("findMissingNumber = %d, want 3", got)
	}
}

func TestQuickSort(t *testing.T) {
	got := quickSort([]int{5, 2, 9, 1, 5, 6})
	want := []int{1, 2, 5, 5, 6, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("quickSort = %v, want %v", got, want)
	}
}

func TestSieveOfEratosthenes(t *testing.T) {
	got := sieveOfEratosthenes([]int{2, 3, 4, 5, 6, 7, 8, 9, 10})
	want := []int{2, 3, 5, 7}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("sieveOfEratosthenes = %v, want %v", got, want)
	}
}

func TestAreSlicesEqual(t *testing.T) {
	if !areSlicesEqual([]int{1, 2}, []int{1, 2}) {
		t.Error("equal slices reported unequal")
	}
	if areSlicesEqual([]int{1, 2}, []int{2, 1}) {
		t.Error("unequal slices reported equal")
	}
}

func TestBasicMainSmoke(t *testing.T) {
	BasicMain()
}

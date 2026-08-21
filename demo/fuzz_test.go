package main

import (
	"testing"
)

// FuzzIsPalindrome checks a symmetry property: if a string is a palindrome,
// its reverse must also be a palindrome. Run with:
//
//	go test ./demo -run FuzzIsPalindrome -fuzz FuzzIsPalindrome
func FuzzIsPalindrome(f *testing.F) {
	seeds := []string{"", "a", "racecar", "hello", "abba", "AbBa"}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, s string) {
		if isPalindrome(s) && !isPalindrome(reverseString(s)) {
			t.Errorf("palindrome %q reversed to non-palindrome %q", s, reverseString(s))
		}
	})
}

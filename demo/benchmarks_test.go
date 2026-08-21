package main

import (
	"strings"
	"testing"
)

// concatPlus builds a string with repeated += (quadratic: allocates every time)
func concatPlus(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += "x"
	}
	return s
}

// concatBuilder uses strings.Builder (amortised linear)
func concatBuilder(n int) string {
	var b strings.Builder
	for i := 0; i < n; i++ {
		b.WriteByte('x')
	}
	return b.String()
}

var benchSink string

func BenchmarkStringConcatPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchSink = concatPlus(1000)
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchSink = concatBuilder(1000)
	}
}

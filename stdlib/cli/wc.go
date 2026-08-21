// main package demonstrates the flag package with a testable CLI runner
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

// WordCount counts whitespace-separated words, lower-cased
func WordCount(r io.Reader) map[string]int {
	counts := make(map[string]int)
	buf := make([]byte, 4096)
	var word strings.Builder

	flush := func() {
		if word.Len() > 0 {
			counts[strings.ToLower(word.String())]++
			word.Reset()
		}
	}

	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if b == ' ' || b == '\n' || b == '\t' || b == '\r' {
				flush()
			} else {
				word.WriteByte(b)
			}
		}
		if err != nil {
			break
		}
	}
	flush()
	return counts
}

// TopWords returns the n most frequent words in descending order
func TopWords(counts map[string]int, n int) []string {
	type pair struct {
		word  string
		count int
	}
	pairs := make([]pair, 0, len(counts))
	for w, c := range counts {
		pairs = append(pairs, pair{w, c})
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].count != pairs[j].count {
			return pairs[i].count > pairs[j].count
		}
		return pairs[i].word < pairs[j].word
	})
	if n > len(pairs) {
		n = len(pairs)
	}
	out := make([]string, 0, n)
	for _, p := range pairs[:n] {
		out = append(out, fmt.Sprintf("%s: %d", p.word, p.count))
	}
	return out
}

// Run is the testable entry point: it parses args and writes results
func Run(args []string, input io.Reader, stdout io.Writer) error {
	fs := flag.NewFlagSet("wc", flag.ContinueOnError)
	top := fs.Int("top", 5, "show only the N most frequent words")
	if err := fs.Parse(args); err != nil {
		return err
	}

	counts := WordCount(input)
	for _, line := range TopWords(counts, *top) {
		fmt.Fprintln(stdout, line)
	}
	return nil
}

// CLIMain demonstrates the word counter on a sample text
func CLIMain() {
	sample := "Go is fun. Go is fast. Go ships as one binary."
	fmt.Println("word count for sample text:")
	Run([]string{"-top", "3"}, strings.NewReader(sample), os.Stdout)
}

// main runs the demo entry points of this package
func main() {
	CLIMain()
}

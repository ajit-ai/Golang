package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestWordCount(t *testing.T) {
	counts := WordCount(strings.NewReader("Go go GO rust\nrust go"))
	want := map[string]int{"go": 4, "rust": 2}
	for w, c := range want {
		if counts[w] != c {
			t.Errorf("count[%q] = %d, want %d", w, counts[w], c)
		}
	}
	if len(counts) != len(want) {
		t.Errorf("unexpected extra words: %v", counts)
	}
}

func TestWordCountEmpty(t *testing.T) {
	if counts := WordCount(strings.NewReader("")); len(counts) != 0 {
		t.Errorf("empty input produced %v", counts)
	}
}

func TestTopWordsOrdering(t *testing.T) {
	counts := map[string]int{"b": 2, "a": 2, "c": 5}
	got := TopWords(counts, 2)
	if got[0] != "c: 5" {
		t.Errorf("first = %q, want c: 5", got[0])
	}
	if got[1] != "a: 2" {
		t.Errorf("tie should be broken alphabetically, got %q", got[1])
	}
}

func TestTopWordsNLargerThanDistinct(t *testing.T) {
	got := TopWords(map[string]int{"x": 1}, 10)
	if len(got) != 1 {
		t.Errorf("len = %d, want 1", len(got))
	}
}

func TestRunParsesFlagAndWritesOutput(t *testing.T) {
	var out bytes.Buffer
	err := Run([]string{"-top", "1"}, strings.NewReader("go go go rust"), &out)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out.String(), "go: 3") {
		t.Errorf("output = %q", out.String())
	}
}

func TestRunRejectsUnknownFlag(t *testing.T) {
	var out bytes.Buffer
	if err := Run([]string{"-nope"}, strings.NewReader(""), &out); err == nil {
		t.Error("unknown flag should produce an error")
	}
}

func TestCLIMainSmoke(t *testing.T) {
	CLIMain()
}

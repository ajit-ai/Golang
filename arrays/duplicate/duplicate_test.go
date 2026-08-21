package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindDuplicatesVariants(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3}
	want := []int{1, 2, 3}

	variants := map[string]func([]int) []int{
		"map":        findDuplicates,
		"set":        findDuplicatestwo,
		"hash table": findDuplicatesthree,
		"bool slice": findDuplicatesfour,
		"sorting":    findDuplicatesfive,
		"recursive":  findDuplicatesseven,
		"closure":    findDuplicateseight,
		"goroutine":  findDuplicatesnine,
		"waitgroup":  findDuplicatesten,
		"mutex":      findDuplicateseleven,
	}
	for name, fn := range variants {
		t.Run(name, func(t *testing.T) {
			got := fn(append([]int(nil), arr...))
			sorted := append([]int(nil), got...)
			sort.Ints(sorted)
			if !reflect.DeepEqual(sorted, want) {
				t.Errorf("%s: duplicates = %v, want %v", name, sorted, want)
			}
		})
	}
}

func TestFindDuplicatesNoDuplicates(t *testing.T) {
	if got := findDuplicates([]int{1, 2, 3}); len(got) != 0 {
		t.Errorf("expected no duplicates, got %v", got)
	}
}

func TestFindDuplicatesEmpty(t *testing.T) {
	if got := findDuplicatestwo([]int{}); len(got) != 0 {
		t.Errorf("expected empty result, got %v", got)
	}
}

func TestFindDuplicatesBoolSliceLargeValues(t *testing.T) {
	got := findDuplicatesfour([]int{1000000, 1000000})
	if len(got) != 1 || got[0] != 1000000 {
		t.Errorf("bool-slice variant failed for large values: %v", got)
	}
}

func TestDuplicateMainSmoke(t *testing.T) {
	main()
}

package main

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestWriteReadAppendLines(t *testing.T) {
	path := filepath.Join(t.TempDir(), "notes.txt")

	if err := WriteLines(path, []string{"hello", "world"}); err != nil {
		t.Fatal(err)
	}
	lines, err := ReadLines(path)
	if err != nil || !reflect.DeepEqual(lines, []string{"hello", "world"}) {
		t.Fatalf("ReadLines = %v, %v", lines, err)
	}

	if err := AppendLine(path, "third"); err != nil {
		t.Fatal(err)
	}
	lines, _ = ReadLines(path)
	if len(lines) != 3 || lines[2] != "third" {
		t.Errorf("after append = %v", lines)
	}
}

func TestReadLinesMissingFile(t *testing.T) {
	if _, err := ReadLines(filepath.Join(t.TempDir(), "nope.txt")); err == nil {
		t.Error("expected error for missing file")
	}
}

func TestCountWords(t *testing.T) {
	path := filepath.Join(t.TempDir(), "w.txt")
	if err := os.WriteFile(path, []byte("one two\nthree  four\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	n, err := CountWords(path)
	if err != nil || n != 4 {
		t.Errorf("CountWords = %d, %v; want 4, nil", n, err)
	}
	if _, err := CountWords(filepath.Join(t.TempDir(), "nope")); err == nil {
		t.Error("expected error for missing file")
	}
}

func TestCSVRoundTrip(t *testing.T) {
	path := filepath.Join(t.TempDir(), "p.csv")
	records := [][]string{{"name", "age"}, {"ajit", "30"}, {"priya", "28"}}
	if err := WriteCSV(path, records); err != nil {
		t.Fatal(err)
	}
	got, err := ReadCSV(path)
	if err != nil || !reflect.DeepEqual(got, records) {
		t.Errorf("CSV round trip = %v, %v", got, err)
	}
}

func TestFindFilesByExt(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "a.txt"), []byte("x"), 0o644)
	os.WriteFile(filepath.Join(dir, "b.md"), []byte("x"), 0o644)
	sub := filepath.Join(dir, "sub")
	os.Mkdir(sub, 0o755)
	os.WriteFile(filepath.Join(sub, "c.txt"), []byte("x"), 0o644)

	got, err := FindFilesByExt(dir, ".txt")
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 {
		t.Errorf("found %d .txt files, want 2: %v", len(got), got)
	}
}

func TestFilesMainSmoke(t *testing.T) {
	FilesMain()
}

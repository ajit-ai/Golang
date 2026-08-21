// main package demonstrates file I/O: os read/write/append,
// buffered scanning, CSV and directory walking
package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// WriteLines creates (or truncates) path with one line per entry
func WriteLines(path string, lines []string) error {
	return os.WriteFile(path, []byte(strings.Join(lines, "\n")+"\n"), 0o644)
}

// ReadLines returns the non-empty lines of the file at path
func ReadLines(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var lines []string
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines, nil
}

// AppendLine adds one line to the end of the file, creating it if needed
func AppendLine(path, line string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(line + "\n")
	return err
}

// CountWords counts whitespace-separated words using a buffered scanner
func CountWords(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count, scanner.Err()
}

// WriteCSV writes records in CSV format
func WriteCSV(path string, records [][]string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return csv.NewWriter(f).WriteAll(records)
}

// ReadCSV reads all CSV records from path
func ReadCSV(path string) ([][]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return csv.NewReader(f).ReadAll()
}

// FindFilesByExt walks root recursively and returns files with the extension
func FindFilesByExt(root, ext string) ([]string, error) {
	var matches []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && filepath.Ext(path) == ext {
			matches = append(matches, path)
		}
		return nil
	})
	return matches, err
}

// FilesMain demonstrates the helpers on a temporary directory
func FilesMain() {
	dir, err := os.MkdirTemp("", "gofiles")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	txt := filepath.Join(dir, "notes.txt")
	if err := WriteLines(txt, []string{"hello", "world"}); err != nil {
		panic(err)
	}
	if err := AppendLine(txt, "appended"); err != nil {
		panic(err)
	}
	lines, _ := ReadLines(txt)
	fmt.Println("lines:", lines)

	n, _ := CountWords(txt)
	fmt.Println("words:", n)

	csvPath := filepath.Join(dir, "people.csv")
	_ = WriteCSV(csvPath, [][]string{{"name", "age"}, {"ajit", "30"}})
	records, _ := ReadCSV(csvPath)
	fmt.Println("csv:", records)

	matches, _ := FindFilesByExt(dir, ".txt")
	fmt.Println("txt files:", len(matches))
}

// main runs the demo entry points of this package
func main() {
	FilesMain()
}

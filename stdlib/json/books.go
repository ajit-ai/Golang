// main package demonstrates encoding/json: struct tags, omitempty,
// field skipping, nested objects and streaming decode
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// Author is a nested struct
type Author struct {
	Name string `json:"name"`
	Web  string `json:"web,omitempty"` // omitted when empty
}

// Book shows common struct tag options
type Book struct {
	Title    string  `json:"title"`
	Author   Author  `json:"author"`
	Price    float64 `json:"price,omitempty"` // omitted when zero
	Internal string  `json:"-"`               // never serialised
}

// UnmarshalBooks decodes a JSON array of books
func UnmarshalBooks(data string) ([]Book, error) {
	var books []Book
	if err := json.Unmarshal([]byte(data), &books); err != nil {
		return nil, err
	}
	return books, nil
}

// MarshalBooks encodes books back to JSON
func MarshalBooks(books []Book) (string, error) {
	out, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// DecodeStream reads concatenated JSON objects one at a time with a Decoder
func DecodeStream(data string) ([]Book, error) {
	dec := json.NewDecoder(strings.NewReader(data))
	var books []Book
	for dec.More() {
		var b Book
		if err := dec.Decode(&b); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

// JSONMain demonstrates the round trip end to end
func JSONMain() {
	books := []Book{
		{Title: "Learn Go", Author: Author{Name: "Ajit"}, Price: 29.99},
		{Title: "Go in Depth", Author: Author{Name: "Priya", Web: "priya.dev"}},
	}

	out, err := MarshalBooks(books)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)

	back, err := UnmarshalBooks(out)
	if err != nil {
		panic(err)
	}
	for _, b := range back {
		fmt.Printf("%s by %s\n", b.Title, b.Author.Name)
	}

	var buf bytes.Buffer
	json.NewEncoder(&buf).SetIndent("", " ")
	fmt.Println("encoder works:", buf.Len() == 0)
}

// main runs the demo entry points of this package
func main() {
	JSONMain()
}

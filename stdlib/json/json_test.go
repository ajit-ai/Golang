package main

import (
	"strings"
	"testing"
)

var sample = `[{"title":"Learn Go","author":{"name":"Ajit"},"price":29.99},
 {"title":"Go in Depth","author":{"name":"Priya","web":"priya.dev"}}]`

func TestUnmarshalBooks(t *testing.T) {
	books, err := UnmarshalBooks(sample)
	if err != nil {
		t.Fatal(err)
	}
	if len(books) != 2 {
		t.Fatalf("decoded %d books, want 2", len(books))
	}
	if books[0].Title != "Learn Go" || books[0].Author.Name != "Ajit" || books[0].Price != 29.99 {
		t.Errorf("book 0 = %+v", books[0])
	}
	if books[1].Author.Web != "priya.dev" {
		t.Errorf("book 1 author web = %q", books[1].Author.Web)
	}
}

func TestUnmarshalInvalidJSON(t *testing.T) {
	if _, err := UnmarshalBooks("{not json"); err == nil {
		t.Error("expected error for invalid JSON")
	}
}

func TestMarshalOmitsEmptyAndSkippedFields(t *testing.T) {
	books := []Book{{Title: "T", Author: Author{Name: "N"}, Internal: "secret"}}
	out, err := MarshalBooks(books)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(out, "secret") {
		t.Error(`json:"-" field must not be serialised`)
	}
	if strings.Contains(out, "price") {
		t.Error("zero price with omitempty must be omitted")
	}
	if strings.Contains(out, "web") {
		t.Error("empty web with omitempty must be omitted")
	}
}

func TestRoundTrip(t *testing.T) {
	books := []Book{{Title: "X", Author: Author{Name: "Y"}, Price: 9.5}}
	out, err := MarshalBooks(books)
	if err != nil {
		t.Fatal(err)
	}
	back, err := UnmarshalBooks(out)
	if err != nil {
		t.Fatal(err)
	}
	if back[0] != books[0] {
		t.Errorf("round trip mismatch: %+v vs %+v", back[0], books[0])
	}
}

func TestDecodeStream(t *testing.T) {
	stream := `{"title":"A","author":{"name":"x"}}{"title":"B","author":{"name":"y"}}`
	books, err := DecodeStream(stream)
	if err != nil {
		t.Fatal(err)
	}
	if len(books) != 2 || books[1].Title != "B" {
		t.Errorf("stream decoded %+v", books)
	}
}

func TestJSONMainSmoke(t *testing.T) {
	JSONMain()
}

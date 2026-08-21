// main package demonstrates a small JSON REST API using Go 1.22
// ServeMux method-and-path patterns, tested with httptest
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

// Book is the API resource
type Book struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// BookStore is a concurrency-safe in-memory store
type BookStore struct {
	mu     sync.Mutex
	nextID int
	books  map[int]Book
}

// NewBookStore creates an empty store
func NewBookStore() *BookStore {
	return &BookStore{nextID: 1, books: make(map[int]Book)}
}

// Add stores a book and assigns its ID
func (s *BookStore) Add(b Book) Book {
	s.mu.Lock()
	defer s.mu.Unlock()
	b.ID = s.nextID
	s.nextID++
	s.books[b.ID] = b
	return b
}

// Get returns a book by id; ok is false when missing
func (s *BookStore) Get(id int) (Book, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	b, ok := s.books[id]
	return b, ok
}

// All returns every book in insertion order
func (s *BookStore) All() []Book {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]Book, 0, len(s.books))
	for id := 1; id < s.nextID; id++ {
		if b, ok := s.books[id]; ok {
			out = append(out, b)
		}
	}
	return out
}

// Delete removes a book; ok is false when missing
func (s *BookStore) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.books[id]; !ok {
		return false
	}
	delete(s.books, id)
	return true
}

// NewMux wires all routes using Go 1.22 pattern routing
func NewMux(store *BookStore) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /books", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, store.All())
	})

	mux.HandleFunc("POST /books", func(w http.ResponseWriter, r *http.Request) {
		var b Book
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil || b.Title == "" {
			http.Error(w, "invalid body", http.StatusBadRequest)
			return
		}
		writeJSON(w, http.StatusCreated, store.Add(b))
	})

	mux.HandleFunc("GET /books/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "bad id", http.StatusBadRequest)
			return
		}
		b, ok := store.Get(id)
		if !ok {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		writeJSON(w, http.StatusOK, b)
	})

	mux.HandleFunc("DELETE /books/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "bad id", http.StatusBadRequest)
			return
		}
		if !store.Delete(id) {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})

	return mux
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// HTTPAPIMain starts the API on :8080 (blocking; not wired into auto-runners)
func HTTPAPIMain() {
	fmt.Println("listening on :8080")
	http.ListenAndServe(":8080", NewMux(NewBookStore()))
}

// main runs the demo entry points of this package
func main() {
	HTTPAPIMain()
}

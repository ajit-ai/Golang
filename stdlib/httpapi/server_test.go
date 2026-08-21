package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func do(t *testing.T, mux *http.ServeMux, method, path, body string) (*httptest.ResponseRecorder, *BookStore) {
	t.Helper()
	store := NewBookStore()
	server := httptest.NewServer(mux)
	defer server.Close()

	var req *http.Request
	var err error
	if body == "" {
		req, err = http.NewRequest(method, server.URL+path, nil)
	} else {
		req, err = http.NewRequest(method, server.URL+path, strings.NewReader(body))
	}
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	return rec, store
}

func TestCreateAndGetBook(t *testing.T) {
	mux := NewMux(NewBookStore())

	rec, _ := do(t, mux, "POST", "/books", `{"title":"Learn Go"}`)
	if rec.Code != http.StatusCreated {
		t.Fatalf("POST status = %d, want 201", rec.Code)
	}
	var created Book
	json.Unmarshal(rec.Body.Bytes(), &created)
	if created.ID != 1 || created.Title != "Learn Go" {
		t.Errorf("created = %+v", created)
	}

	rec, _ = do(t, mux, "GET", "/books/1", "")
	if rec.Code != http.StatusOK {
		t.Fatalf("GET status = %d, want 200", rec.Code)
	}
	var got Book
	json.Unmarshal(rec.Body.Bytes(), &got)
	if got.Title != "Learn Go" {
		t.Errorf("got = %+v", got)
	}
}

func TestGetMissingBookReturns404(t *testing.T) {
	rec, _ := do(t, NewMux(NewBookStore()), "GET", "/books/42", "")
	if rec.Code != http.StatusNotFound {
		t.Errorf("status = %d, want 404", rec.Code)
	}
}

func TestPostInvalidBodyReturns400(t *testing.T) {
	for _, body := range []string{"not json", `{}`, `{"title":""}`} {
		rec, _ := do(t, NewMux(NewBookStore()), "POST", "/books", body)
		if rec.Code != http.StatusBadRequest {
			t.Errorf("POST %q status = %d, want 400", body, rec.Code)
		}
	}
}

func TestListAndDelete(t *testing.T) {
	store := NewBookStore()
	store.Add(Book{Title: "A"})
	store.Add(Book{Title: "B"})
	mux := NewMux(store)

	rec, _ := do(t, mux, "GET", "/books", "")
	var books []Book
	json.Unmarshal(rec.Body.Bytes(), &books)
	if len(books) != 2 || books[0].Title != "A" {
		t.Errorf("list = %+v", books)
	}

	rec, _ = do(t, mux, "DELETE", "/books/1", "")
	if rec.Code != http.StatusNoContent {
		t.Fatalf("DELETE status = %d, want 204", rec.Code)
	}
	rec, _ = do(t, mux, "DELETE", "/books/1", "")
	if rec.Code != http.StatusNotFound {
		t.Errorf("second DELETE status = %d, want 404", rec.Code)
	}
}

func TestMethodNotAllowed(t *testing.T) {
	rec, _ := do(t, NewMux(NewBookStore()), "PUT", "/books", "")
	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("PUT /books status = %d, want 405", rec.Code)
	}
}

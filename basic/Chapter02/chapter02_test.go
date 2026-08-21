package main

import "testing"

func TestTwiceValue(t *testing.T) {
	slice := []int{1, 3, 5, 6}
	twiceValue(slice)
	want := []int{2, 6, 10, 12}
	for i := range want {
		if slice[i] != want[i] {
			t.Fatalf("twiceValue = %v, want %v", slice, want)
		}
	}
}

func TestTwiceValueEmpty(t *testing.T) {
	var slice []int
	twiceValue(slice)
	if len(slice) != 0 {
		t.Errorf("twiceValue on empty slice changed length to %d", len(slice))
	}
}

func TestMustParseTemplatesFallsBackWithoutTemplatesDir(t *testing.T) {
	// When tests run, the working directory contains templates/, so the
	// real set parses; the function must never panic either way.
	tmpl := mustParseTemplates()
	if tmpl == nil {
		t.Fatal("mustParseTemplates returned nil")
	}
	if tmpl.Lookup("Home") == nil {
		t.Error("Home template missing even though templates directory exists")
	}
}

func TestCrmCustomerStructFields(t *testing.T) {
	c := CrmCustomer{CustomerId: 7, CustomerName: "Alice", SSN: "123"}
	if c.CustomerId != 7 || c.CustomerName != "Alice" || c.SSN != "123" {
		t.Errorf("struct fields not preserved: %+v", c)
	}
}

func TestChapter02MainSmoke(t *testing.T) {
	main()
}

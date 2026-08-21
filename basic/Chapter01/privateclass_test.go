package main

import (
	"encoding/json"
	"testing"
)

func TestSetDetails(t *testing.T) {
	account := &PrivateClassAccount{CustomerName: "John Smith"}
	account.setDetails("4532", "current")
	if got := account.getId(); got != "4532" {
		t.Errorf("getId() = %q, want %q", got, "4532")
	}
	if got := account.getAccountType(); got != "current" {
		t.Errorf("getAccountType() = %q, want %q", got, "current")
	}
}

func TestPrivateDetailsHiddenFromJSON(t *testing.T) {
	account := &PrivateClassAccount{CustomerName: "John Smith"}
	account.setDetails("4532", "current")
	data, err := json.Marshal(account)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}
	if string(data) != `{"CustomerName":"John Smith"}` {
		t.Errorf("marshaled = %s, want private details hidden", data)
	}
}

func TestPrivateClassMain(t *testing.T) {
	PrivateClassMain()
}

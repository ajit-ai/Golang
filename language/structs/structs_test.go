package main

import (
	"reflect"
	"testing"
)

func TestEmbeddingPromotion(t *testing.T) {
	car := Car{
		Engine: Engine{Horsepower: 300, Type: "V6"},
		Wheels: &Wheels{Count: 4},
		Brand:  "GoMobile",
	}
	if car.Horsepower != 300 {
		t.Error("Engine fields should be promoted to Car")
	}
	if got := car.Start(); got != "V6 engine started" {
		t.Errorf("promoted Start() = %q", got)
	}
	if got := car.Describe(); got != "GoMobile with 300 HP on 4 wheels" {
		t.Errorf("Describe() = %q", got)
	}
}

func TestReceiverSemantics(t *testing.T) {
	c := Counter{}
	if c.IncrementValue().n != 1 {
		t.Error("value receiver should return incremented copy")
	}
	if c.n != 0 {
		t.Errorf("value receiver must not mutate caller: n = %d, want 0", c.n)
	}
	c.IncrementPointer()
	c.IncrementPointer()
	if c.n != 2 {
		t.Errorf("pointer receiver should mutate caller: n = %d, want 2", c.n)
	}
}

func TestTagLookup(t *testing.T) {
	got := TagLookup(Config{})
	want := map[string]string{
		"Host":  "required",
		"Port":  "min=1,max=65535",
		"Debug": "optional",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("TagLookup = %v, want %v", got, want)
	}
}

func TestStructsMainSmoke(t *testing.T) {
	EmbeddingMain()
	ReceiversMain()
	TagsMain()
}

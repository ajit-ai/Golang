package main

import "testing"

func TestGreetAllPolymorphism(t *testing.T) {
	speakers := []Speaker{
		Dog{Name: "Rex"},
		Robot{Model: "RX-7"},
		Human{Name: "Ajit"},
	}
	lines := GreetAll(speakers)
	want := []string{
		"Rex says woof",
		"RX-7 beeps",
		"Ajit says hello",
	}
	if len(lines) != len(want) {
		t.Fatalf("GreetAll = %v, want %v", lines, want)
	}
	for i := range want {
		if lines[i] != want[i] {
			t.Fatalf("GreetAll = %v, want %v", lines, want)
		}
	}
}

func TestHumanSatisfiesTriathlete(t *testing.T) {
	var t1 Triathlete = Human{Name: "Priya"}
	got := TrainAthlete(t1)
	want := "Priya runs 10km; Priya swims 1km; Priya says hello"
	if got != want {
		t.Errorf("TrainAthlete = %q, want %q", got, want)
	}
}

func TestStringer(t *testing.T) {
	p := Person{Name: "Ajit", Age: 30}
	if got := p.String(); got != "Ajit (30 years old)" {
		t.Errorf("String() = %q", got)
	}
}

func TestInterfacesMainSmoke(t *testing.T) {
	InterfacesMain()
	EmbeddingMain()
}

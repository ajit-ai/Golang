package main

import "testing"

func TestMakeMap(t *testing.T) {
	got := makeMap()
	if len(got) != 2 || got["x"] != 10 || got["y"] != 20 {
		t.Errorf("makeMap() = %v, want map[x:10 y:20]", got)
	}
}

func TestMapsUseMainSmoke(t *testing.T) {
	MapsUseMain()
}

package main

import "testing"

var _ AdapterIProcess = Adapter{}

func TestAdapterMain(t *testing.T) {
	AdapterMain()
}

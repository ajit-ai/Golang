package main

import "testing"

var _ IContour = &DrawContour{}

func TestResizeByFactor(t *testing.T) {
	contour := &DrawContour{factor: 2}
	contour.resizeByFactor(5)
	if contour.factor != 5 {
		t.Errorf("contour.factor = %d, want 5", contour.factor)
	}
}

func TestBridgeMain(t *testing.T) {
	BridgeMain()
}

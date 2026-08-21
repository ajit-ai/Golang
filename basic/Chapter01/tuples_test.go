package main

import "testing"

func TestPowerSeries(t *testing.T) {
	tests := []struct {
		a      int
		square int
		cube   int
	}{
		{3, 9, 27},
		{0, 0, 0},
		{1, 1, 1},
		{-2, 4, -8},
		{10, 100, 1000},
	}
	for _, tt := range tests {
		square, cube := PowerSeries(tt.a)
		if square != tt.square || cube != tt.cube {
			t.Errorf("PowerSeries(%d) = (%d, %d), want (%d, %d)", tt.a, square, cube, tt.square, tt.cube)
		}
	}
}

func TestPowerSeriesN(t *testing.T) {
	tests := []struct {
		a      int
		square int
		cube   int
	}{
		{4, 16, 64},
		{0, 0, 0},
		{-3, 9, -27},
	}
	for _, tt := range tests {
		square, cube := PowerSeriesN(tt.a)
		if square != tt.square || cube != tt.cube {
			t.Errorf("PowerSeriesN(%d) = (%d, %d), want (%d, %d)", tt.a, square, cube, tt.square, tt.cube)
		}
	}
}

func TestPowerSeriesE(t *testing.T) {
	tests := []struct {
		a      int
		square int
		cube   int
	}{
		{5, 25, 125},
		{0, 0, 0},
		{-4, 16, -64},
	}
	for _, tt := range tests {
		square, cube, err := PowerSeriesE(tt.a)
		if err != nil {
			t.Fatalf("PowerSeriesE(%d) returned error: %v", tt.a, err)
		}
		if square != tt.square || cube != tt.cube {
			t.Errorf("PowerSeriesE(%d) = (%d, %d), want (%d, %d)", tt.a, square, cube, tt.square, tt.cube)
		}
	}
}

func TestTuplesMain(t *testing.T) {
	TuplesMain()
}

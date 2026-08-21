package main

import "testing"

func TestFilledCube(t *testing.T) {
	got := FilledCube()
	for k := 0; k < 10; k++ {
		for l := 0; l < 10; l++ {
			for m := 0; m < 10; m++ {
				if got[k][l][m] != 1 {
					t.Fatalf("got[%d][%d][%d] = %d, want 1", k, l, m, got[k][l][m])
				}
			}
		}
	}
}

func TestCubicComplexityMain(t *testing.T) {
	CubicComplexityMain()
}

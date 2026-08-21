package main

import "fmt"

func TwodArrayMain() {
	var twod [][]int
	twod = make([][]int, 3)
	for i := 0; i < 3; i++ {
		twod[i] = make([]int, 4)
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			twod[i][j] = i + j
		}
	}
	fmt.Println(twod)

}

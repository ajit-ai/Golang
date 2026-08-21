package main

func TwodSliceMain() {

	// Declares a 2D slice of integers
	var dynamicMatrix [][]int

	// Initialize a 2D slice with specific dimensions
	rows := 3
	cols := 4
	dynamicMatrix = make([][]int, rows)
	for i := range dynamicMatrix {
		dynamicMatrix[i] = make([]int, cols)
	}
	// Assign values to the 2D slice
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			dynamicMatrix[i][j] = i + j
		}
	}

	// Print the 2D slice
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			print(dynamicMatrix[i][j], " ")
		}
		println()
	}
}

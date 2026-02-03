package main

import (
	"fmt"
)

/*
You are given a square grid with n
 rows and n
 columns. Each cell contains either 0
 or 1
.

In an operation, you can select a cell of the grid and flip it (from 0→1
 or 1→0
). Find the minimum number of operations you need to obtain a square that remains the same when rotated 0∘
, 90∘
, 180∘
 and 270∘
.

*/

func main() {

	var testCase int
	fmt.Scan(&testCase)
	var n int
	fmt.Scan(&n)
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	operationCount := 0

	row := 0
	col := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] != matrix[row][col] {
				operationCount = operationCount + 1
				matrix[i][j] = matrix[row][col]
			}
			row = row + 1
		}
		row = 0
		col = col + 1
	}

	fmt.Println(operationCount)
}

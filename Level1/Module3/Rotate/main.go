package main

import "fmt"

/*
Problem Statement
You are given a grid with
N rows and
N columns. An integer
A
i,j
​
  is written on the square at the
i-th row from the top and
j-th column from the left. Here, it is guaranteed that
A
i,j
​
  is either
0 or
1.

Shift the integers written on the outer squares clockwise by one square each, and print the resulting grid.

Here, the outer squares are those in at least one of the
1-st row,
N-th row,
1-st column, and
N-th column.

Constraints
2≤N≤100
0≤A
i,j
​
 ≤1(1≤i,j≤N)
All input values are integers.

Input 1:
4
0101
1101
1111
0000

Output:
1010
1101
0111
0001



*/
func main() {

	var n int
	fmt.Scan(&n)

	matrix := make([][]int, n, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	// Rotate Outer Spaces Clockwise

}

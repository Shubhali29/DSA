package main

import (
	"fmt"
)

// Question: Subarray Sum queries in O(1)

func main() {

	arr2d := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// Precompute prefix sums
	prefixSum2d := make([][]int, len(arr2d))
	for i := range prefixSum2d {
		prefixSum2d[i] = make([]int, len(arr2d[i]))
	}

	for i := 0; i < len(arr2d); i++ {
		for j := 0; j < len(arr2d[i]); j++ {
			if i == 0 && j == 0 {
				prefixSum2d[i][j] = arr2d[i][j]
			} else if i == 0 {
				prefixSum2d[i][j] = prefixSum2d[i][j-1] + arr2d[i][j]
			} else if j == 0 {
				prefixSum2d[i][j] = prefixSum2d[i-1][j] + arr2d[i][j]
			} else {
				prefixSum2d[i][j] = prefixSum2d[i-1][j] + prefixSum2d[i][j-1] - prefixSum2d[i-1][j-1] + arr2d[i][j]
			}
		}
	}

	// Queries
	var q int
	fmt.Println("Enter the number of queries:")
	fmt.Scan(&q)
	queries := make([][4]int, q)
	fmt.Println("Enter the queries (x1 y1 x2 y2):")
	for i := 0; i < q; i++ {
		fmt.Scan(&queries[i][0], &queries[i][1], &queries[i][2], &queries[i][3])
	}

	for _, query := range queries {
		x1, y1, x2, y2 := query[0], query[1], query[2], query[3]
		if x1 < 0 || y1 < 0 || x2 >= len(arr2d) || y2 >= len(arr2d[0]) || x1 > x2 || y1 > y2 {
			fmt.Println("Invalid query")
			continue
		}
		sum := prefixSum2d[x2][y2]
		if x1 > 0 {
			sum -= prefixSum2d[x1-1][y2]
		}
		if y1 > 0 {
			sum -= prefixSum2d[x2][y1-1]
		}
		if x1 > 0 && y1 > 0 {
			sum += prefixSum2d[x1-1][y1-1]
		}
		fmt.Println("Sum of the subarray:", sum)
	}

	// TC = O(n*m + q)
	// SC = O(n*m) for prefix sum array
	// where n is the number of rows and m is the number of columns in the 2D array, and q is the number of queries.
	// The prefix sum array is used to store the cumulative sum of the elements in the 2D array.
}

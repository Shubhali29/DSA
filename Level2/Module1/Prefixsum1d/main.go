package main

// Question: Subarray Sum queries in O(1)

import (
	"fmt"
)

func main() {

	var arr []int
	fmt.Println("Enter the size of the array:")
	var n int
	fmt.Scan(&n)
	arr = make([]int, n)
	fmt.Println("Enter the elements of the array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Enter the number of queries:")
	var q int
	fmt.Scan(&q)
	queries := make([][2]int, q)
	fmt.Println("Enter the queries (l r):")
	for i := 0; i < q; i++ {
		fmt.Scan(&queries[i][0], &queries[i][1])
	}

	// Precompute prefix sums
	prefixSum := make([]int, n)
	prefixSum[0] = arr[0]
	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i-1] + arr[i]
	}

	for _, query := range queries {
		l, r := query[0], query[1]
		if l < 0 || r >= n || l > r {
			fmt.Println("Invalid query")
			continue
		}
		if l == 0 {
			fmt.Println(prefixSum[r])
		} else {
			fmt.Println(prefixSum[r] - prefixSum[l-1])
		}
	}
	// TC := O(n + q)
	// SC := O(n) for prefix sum array
	// where n is the size of the array and q is the number of queries
	// The prefix sum array is used to store the cumulative sum of the elements in the array.
	// This allows us to calculate the sum of any subarray in O(1) time by using the formula:
	// sum(l, r) = prefixSum[r] - prefixSum[l-1] if l > 0
	// sum(l, r) = prefixSum[r] if l == 0
	// where l is the starting index and r is the ending index of the subarray.
	// The prefix sum array is built in O(n) time, and each query is answered in O(1) time.
	// Therefore, the overall time complexity is O(n + q).
	// The space complexity is O(n) for the prefix sum array.
	// This approach is efficient for large arrays and multiple queries.
	// It is also easy to implement and understand.
	// The prefix sum array can be used for various applications, such as range queries, histogram calculations, and more.
	// Overall, this approach is a good choice for solving the problem of subarray sum queries in O(1) time.
	// It is a common technique used in competitive programming and algorithm design.
	// The prefix sum array can be easily modified to handle different types of queries, such as range minimum or maximum queries, by using different data structures or algorithms.
	// This makes it a versatile and powerful tool for solving a wide range of problems in computer science and mathematics.
}

package main

import (
	"fmt"
)

// Question: Given an array of n elements, Answer following queries
// Query: L,R
// Result : A[L] + 2A[L+1] + 3A[L+2] .... + (R-L+1)*A[R]

// Hint: In these type of question try to make summation expression
// Summation(i=L to R) A[i]*(i-L+1)
// Summation(i=L to R) (i*A[i]) - summation(i=L to R) A[i]*(L-1)
// Summation(i=L to R) (i*A[i]) - (L-1) Summation(i=L to R) A[i]

func main() {

	array := []int{1, 2, 3, 1, 2}

	// calculate prefix sum so that query can be resolved in O(1) time
	prefixSum := make([]int, len(array))
	prefixSumI := make([]int, len(array))

	prefixSum[0] = array[0]
	prefixSumI[0] = 0
	for i := 1; i < len(array); i++ {

		prefixSum[i] = prefixSum[i-1] + array[i]
		prefixSumI[i] = prefixSumI[i-1] + (i * array[i])
	}

	// queries
	var queryNum int
	fmt.Println("Enter number of queries")
	fmt.Scan(&queryNum)

	query := make([][2]int, queryNum)

	for i := 0; i < queryNum; i++ {
		fmt.Println("Enter number of query input")
		fmt.Scan(&query[i][0], &query[i][1])
	}

	// calculate final result of query
	for i := 0; i < queryNum; i++ {
		l, r := query[i][0], query[i][1]

		if l < 0 || r < 0 || r >= len(array) || l > r || l >= len(array) {

			fmt.Println("Invalid Query")
			continue
		}
		var sum int
		if l == 0 {
			sum = prefixSumI[r] - (l-1)*prefixSum[r]
		} else {
			sum = (prefixSumI[r] - prefixSumI[l-1]) - (l-1)*(prefixSum[r]-prefixSum[l-1])

		}
		fmt.Println("Sum is : ", sum)
	}

}

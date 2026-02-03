package main

import (
	"fmt"
)

/*
We have a sequence of length
N consisting of positive integers:
A=(A1,…,AN). Any two adjacent terms have different values.

Let us insert some numbers into this sequence by the following procedure.

If every pair of adjacent terms in
A has an absolute difference of
1, terminate the procedure.
Let
Ai,Ai+1 be the pair of adjacent terms nearest to the beginning of
A whose absolute difference is not 1.
If Ai < Ai+1, insert Ai+1,Ai+2,…,Ai+1−1 between Ai and Ai+1.
If Ai>Ai+1, insert Ai−1,Ai−2,…,Ai+1+1 between Ai and Ai+1.
Return to step 1.
Print the sequence when the procedure ends.
*/

func main() {

	var n int
	fmt.Scan(&n)

	series := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&series[i])
	}

	var newseries []int

	for i := 1; i < n; i++ {

		if series[i-1] > series[i] {

			diff := series[i-1] - series[i]

			if diff == 1 {
				newseries = append(newseries, series[i-1])
			} else {

				num := series[i-1]

				for num > series[i] {
					newseries = append(newseries, num)
					num = num - 1
				}
			}
		}

		if series[i-1] < series[i] {

			diff := series[i] - series[i-1]

			if diff == 1 {
				newseries = append(newseries, series[i-1])
			} else {

				num := series[i-1]

				for num < series[i] {
					newseries = append(newseries, num)
					num = num + 1
				}
			}
		}
	}
	newseries = append(newseries, series[n-1])
	fmt.Println(newseries)
}

// 5 4 1 3 5

// 5 4 3 2 1 2 3 4

// Print length of 3n+1 sequence
//
//	if n is odd next n is 3n+1
//	if n is even next n is n/2
//	n can be 1 <= n <= 10^5
package main

import (
	"fmt"
)

func count3n_1Sequence(n int, sequence *[]int) {

	if n <= 1 {
		return
	}

	if n%2 == 0 {
		n = n / 2
	} else {
		n = (3 * n) + 1
	}

	*sequence = append(*sequence, n)
	count3n_1Sequence(n, sequence)
}

func main() {

	var n int
	var sequence []int
	fmt.Println("Enter value of n")
	fmt.Scanln(&n)
	sequence = append(sequence, n)
	count3n_1Sequence(n, &sequence)
	fmt.Printf("Length of 3n+1 sequence is %v", len(sequence))

}

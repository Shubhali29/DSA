// find square root of x which should be higher positive number near by, i.e y^2 <= x
// range - 0 <= x <= 10^9
// search space = 0 <= y <= 10^5 because 10^5 * 10^5 will be greater than 10^9
package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func CheckSqrt(x, y int) bool {

	squareOfy := y * y

	if squareOfy <= x {
		return true
	}
	return false
}

func main() {
	x := 37

	// serach space
	left := 0
	right := 100000
	ans := 0
	for left <= right {

		mid := left + (right-left)/2

		if CheckSqrt(x, mid) {
			ans = max(ans, mid)
			left = left + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Printf("Square root of %d  is %d.", x, ans)
}

package main

import (
	"fmt"
	"strconv"
)

/*
Given an integer n(1<=n<=10), print n rows of the given pattern in minimum length.

1

2 3

6 5 4

7 8 9 10

*/

func main() {

	var num int

	fmt.Scan(&num)
	digit := 1
	fmt.Print(digit)
	fmt.Println()
	for i := 2; i <= num; i++ {
		pattern := ""
		for j := 1; j <= i; j++ {
			if i%2 == 0 {
				pattern = pattern + strconv.Itoa(digit+1) + " "
			} else {
				pattern = strconv.Itoa(digit+1) + " " + pattern
			}
			digit = digit + 1
		}
		fmt.Print(pattern)
		fmt.Println()
	}
}

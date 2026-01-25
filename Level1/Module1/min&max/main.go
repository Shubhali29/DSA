package main

import (
	"fmt"
)

func main() {

	var a, b, c int
	var max, min int
	fmt.Scan(&a, &b, &c)

	if a >= b && a >= c {
		max = a
		if b < c {
			min = b
		} else {
			min = c
		}
	} else if b >= a && b >= c {
		max = b
		if a < c {
			min = a
		} else {
			min = c
		}
	} else {
		max = c
		if a < b {
			min = a
		} else {
			min = b
		}
	}

	fmt.Println(min, max)

}

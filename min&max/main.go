package main

import (
	"fmt"
)

func main() {

	var a, b, c int

	fmt.Println("Enter three numbers : ")
	fmt.Scanln(&a, &b, &c)
	var min, max int
	if a >= b && a >= c {
		max = a
		if b >= c {
			min = c
		} else {
			min = b
		}
	} else if b >= a && b >= c {
		max = b
		if a >= c {
			min = c
		} else {
			min = a
		}
	} else {
		max = c
		if a >= b {
			min = b
		} else {
			min = a
		}
	}
	fmt.Printf("Min: %d, Max: %d\n", min, max)
}

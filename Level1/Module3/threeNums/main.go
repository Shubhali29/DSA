package main

import (
	"fmt"
)

/*
Given two numbers K and S. Determine how many different values of X,Yand Z such that (0≤X,Y,Z≤K) and X+Y+Z=S.

Input
Only one line containing two numbers K and S (0≤K≤3000,0≤S≤3K).

Output
Print the answer required above.
*/

//  4
// 0, 1, 2, 3, 4

func main() {

	var k, s int
	fmt.Scan(&k, &s)

	count := 0

	for x := 0; x <= s; x++ {
		for y := 0; y <= s; y++ {
			for z := 0; z <= s; z++ {
				sum := x + y + z

				if sum == s {
					count++
				}
			}

		}
	}
	fmt.Println(count)
}

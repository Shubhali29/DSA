package main

import (
	"fmt"
)

// Print Pattern
// if n = 3
//  ***
//   *
//  ***

// n = 5
// *****
//    *
//   *
//  *
// *****

/* n = 6
******
    *
   *
  *
 *
******
*/

func main() {
	var n int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
	for i := 2; i < n; i++ {

		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		fmt.Print("*")
		fmt.Println()
	}
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
}

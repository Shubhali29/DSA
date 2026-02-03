package main

import (
	"fmt"
)

/*
Problem Statement
We have five cards with integers
A,
B,
C,
D, and
E written on them, one on each card.

This set of five cards is called a Full house if and only if the following condition is satisfied:

the set has three cards with a same number written on them, and two cards with another same number written on them.
Determine whether the set is a Full house.

Constraints
1≤A,B,C,D,E≤13
Not all of
A,
B,
C,
D, and
E are the same.
All values in input are integers.

*/

func insertInMap(fullHouse map[int]int, a int) {
	if _, exists := fullHouse[a]; exists {
		fullHouse[a] = fullHouse[a] + 1
	} else {
		fullHouse[a] = 1
	}

}
func main() {

	var a, b, c, d, e int

	fmt.Scan(&a, &b, &c, &d, &e)

	fullHouse := make(map[int]int)

	insertInMap(fullHouse, a)
	insertInMap(fullHouse, b)
	insertInMap(fullHouse, c)
	insertInMap(fullHouse, d)
	insertInMap(fullHouse, e)

	if len(fullHouse) > 2 || len(fullHouse) < 2 {
		fmt.Println("No")
		return
	} else {

		for _, value := range fullHouse {

			if value != 2 && value != 3 {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
}

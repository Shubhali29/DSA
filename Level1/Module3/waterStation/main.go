package main

import (
	"fmt"
	"math"
)

/*
There is an ultramarathon course totaling
100km. Water stations are set up every
5km along the course, including the start and goal, for a total of
21.

Takahashi is at the
Nkm point of this course. Find the position of the nearest water station to him.

Under the constraints of this problem, it can be proven that the nearest water station is uniquely determined.
*/
func main() {

	var n int
	fmt.Scan(&n)

	remainder := n % 5

	if remainder == 0 {
		// Already at water station
		fmt.Println(n)
	}

	half := int(math.Ceil(5 / 2))

	if remainder >= half {
		waterStation := (n - remainder) + 5
		fmt.Println(waterStation)
	} else {
		waterStation := (n - remainder)
		fmt.Println(waterStation)
	}
}

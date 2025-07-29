// Sum of list iterative and by recursion

package main

import (
	"fmt"
)

func iterative(list []int) int {

	sum := 0

	for _, value := range list {
		sum += value
	}

	return sum
}

func recursion(list []int, index int) int {

	if index == (len(list) - 1) {
		return list[index]
	}
	return recursion(list, index+1) + list[index]
}
func main() {

	list := []int{1, 2, 3, 4, 5}

	sum := iterative(list)
	fmt.Printf("Sum is %d", sum)

	sum = recursion(list, 0)
	fmt.Printf("\n Sum is %d", sum)
}

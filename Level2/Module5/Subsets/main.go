//
// 123
// []
// []
// []
// [], [3]

package main

import (
	"fmt"
)

func subsets(list []int, startIndex int, subset []int, result *[][]int) {

	if startIndex == len(list) {
		return
	}

	// Do not include the number
	subsets(list, startIndex+1, subset, result)
	// Include the number
	subset = append(subset, list[startIndex])
	*result = append(*result, subset)
	subsets(list, startIndex+1, subset, result)

}

func main() {

	var len int
	fmt.Println("Enter length of array")
	fmt.Scanln(&len)
	list := make([]int, len)
	for i := 0; i < len; i++ {
		fmt.Scanln(&list[i])
	}
	var result [][]int
	fmt.Printf("List is : %v", list)
	result = append(result, []int{})
	subsets(list, 0, []int{}, &result)
	fmt.Println("All subsets of given list ", result)
}

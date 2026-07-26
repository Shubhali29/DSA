package main

import (
	"fmt"
)

func main() {

	array := []int{1, 2, 3, 4, 5, 6, 6, 7, 8}
	target := 6

	if exist, index := BinarySearch(array, target); exist {
		fmt.Printf("%d is found in list at index %d", target, index)
	} else {
		fmt.Printf("%d is not found in list", target)
	}
}

func BinarySearch(list []int, target int) (bool, int) {

	index := -1
	exist := false
	left := 0
	right := len(list) - 1

	for left <= right {
		mid := left + (right-left)/2
		if target == list[mid] {
			exist = true
			index = mid
			break
		} else if target > list[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return exist, index
}

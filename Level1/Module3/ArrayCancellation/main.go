package main

import (
	"fmt"
)

/*
You're given an array a
 of n
 integers, such that a1+a2+⋯+an=0
.

In one operation, you can choose two different indices i
 and j
 (1≤i,j≤n
), decrement ai
 by one and increment aj
 by one. If i<j
 this operation is free, otherwise it costs one coin.

How many coins do you have to spend in order to make all elements equal to 0
?

Input
Each test contains multiple test cases. The first line contains the number of test cases t
 (1≤t≤5000
). Description of the test cases follows.

The first line of each test case contains an integer n
 (1≤n≤105
)  — the number of elements.

The next line contains n
 integers a1,…,an
 (−109≤ai≤109
). It is given that ∑ni=1ai=0
.

It is guaranteed that the sum of n
 over all test cases does not exceed 105
.

Output
For each test case, print the minimum number of coins we have to spend in order to make all elements equal to 0
.

*/

func main() {

	var testCase int

	fmt.Scan(&testCase)
	var result []int
	for i := 1; i <= testCase; i++ {

		var n int
		fmt.Scan(&n)

		array := make([]int, n)

		for i := 0; i < n; i++ {
			fmt.Scan(&array[i])
		}
		coinCount := makeZeroArray(array, n)

		result = append(result, coinCount)

	}

	for i := 0; i < testCase; i++ {
		fmt.Println(result[i])
	}
}

func makeZeroArray(array []int, length int) int {
	coinCount := 0
	i := 0
	j := length - 1

	for j >= 0 && i < length {
		fmt.Println(array)
		if array[i] == 0 && array[j] == 0 {
			i++
			j--
			continue
		} else if array[i] == 0 {
			i++
			continue
		} else if array[j] == 0 {
			j--
			continue
		}

		if i < j {
			if array[i] > 0 && array[j] < 0 {
				value := array[i] + array[j]
				array[i] = value
				array[j] = 0
				j--
				if i == j {
					j--
				}
			} else if array[i] <= 0 {
				i++
			} else if array[j] >= 0 {
				j--
			}
		} else if i > j {
			if array[i] > 0 && array[j] < 0 {
				temp := array[i]
				if array[i] > (-array[j]) {
					temp = -array[j]
				}
				array[j] = array[i] + array[j]
				array[i] = 0
				coinCount = coinCount + temp
				i++
				if i == j {
					i++
				}
			} else if array[i] <= 0 {
				i++
			} else if array[j] >= 0 {
				j--
			}
		}

	}
	return coinCount
}

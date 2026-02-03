package main

import (
	"fmt"
)

/*
There is pasta consisting of
N noodles at Takahashi's home. The length of the
i-th noodle is
A
i
​
 .
Takahashi has a meal plan for the next
M days. On the
i-th day, he is going to choose a pasta noodle of length exactly
B
i
​
  and eat it. If no such noodle is available on any day, his plan fails. Additionally, he cannot eat the same noodle on multiple days.

Can Takahashi accomplish his meal plan?

*/

func main() {

	var n, m int
	fmt.Scan(&n, &m)

	noodles := make([]int, n)
	days := make([]int, m)

	for i := 0; i < n; i++ {
		fmt.Scan(&noodles[i])
	}

	for j := 0; j < m; j++ {
		fmt.Scan(&days[j])
	}

	noodlesMap := make(map[int]int)

	for i := 0; i < n; i++ {
		if _, exists := noodlesMap[noodles[i]]; exists {
			noodlesMap[noodles[i]] += 1
		} else {
			noodlesMap[noodles[i]] = 1
		}
	}

	for j := 0; j < m; j++ {
		value, exists := noodlesMap[days[j]]
		if (!exists) || (exists && value == 0) {
			fmt.Println("No")
			return
		}
		noodlesMap[days[j]] = noodlesMap[days[j]] - 1
	}

	fmt.Println("Yes")
}

// Find next smallest element from left
// for Next greatest monotic increasing stack is used
package main

import (
	"fmt"
)

type monotonic struct {
	stack []int
}

func (m *monotonic) push(val int) {
	m.stack = append(m.stack, val)
}

func (m *monotonic) pop() {

	if !m.empty() {
		m.stack = m.stack[:(m.size() - 1)]
	}
}

func (m *monotonic) size() int {
	return len(m.stack)
}

func (m *monotonic) empty() bool {
	if m.size() == 0 {
		return true
	}
	return false
}

func (m *monotonic) top() int {

	if m.empty() {
		return -1
	}
	return m.stack[m.size()-1]
}

func NextSmallestElementFromLeft() {

	numbers := []int{2, 9, 3, 4, 5, 1}
	// [-1, 2, 2, 3, 4, -1]
	result := make([]int, len(numbers))
	stack := monotonic{stack: []int{}}
	for i, val := range numbers {

		for !stack.empty() && stack.top() >= val {
			stack.pop()
		}

		if stack.empty() {
			result[i] = -1
		} else {
			result[i] = stack.top()
		}

		stack.push(val)
	}

	fmt.Println("Final result : ", result)

}

func NextSmallestElementFromRight() {

	numbers := []int{2, 9, 3, 4, 5, 1}
	// [1, 3 , 1, 1, 1, -1]
	result := make([]int, len(numbers))
	stack := monotonic{stack: []int{}}
	for i := len(numbers) - 1; i >= 0; i-- {

		for !stack.empty() && stack.top() >= numbers[i] {
			stack.pop()
		}

		if stack.empty() {
			result[i] = -1
		} else {
			result[i] = stack.top()
		}

		stack.push(numbers[i])
	}

	fmt.Println("Final result : ", result)

}

func main() {

	NextSmallestElementFromLeft()
	NextSmallestElementFromRight()
}

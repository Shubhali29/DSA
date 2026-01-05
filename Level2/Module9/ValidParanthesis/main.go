// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.
package main

import (
	"fmt"
)

type stack struct {
	container []rune
}

func (m *stack) push(s rune) {
	m.container = append(m.container, s)
}

func (m *stack) pop() {

	if !m.empty() {
		m.container = m.container[:(m.size() - 1)]
	}
}

func (m *stack) size() int {
	return len(m.container)
}

func (m *stack) empty() bool {
	if m.size() == 0 {
		return true
	}
	return false
}

func (m *stack) top() rune {

	if m.empty() {
		return 0
	}
	return m.container[m.size()-1]
}

func isValid(s string) bool {

	parenthesis := stack{
		container: []rune{},
	}
	for _, value := range s {
		if value == '(' || value == '[' || value == '{' {
			parenthesis.push(value)
		} else if value == ')' {
			topEle := parenthesis.top()
			if topEle != '(' {
				return false
			}
			parenthesis.pop()
		} else if value == '}' {
			topEle := parenthesis.top()
			if topEle != '{' {
				return false
			}
			parenthesis.pop()
		} else if value == ']' {
			topEle := parenthesis.top()
			if topEle != '[' {
				return false
			}
			parenthesis.pop()
		}

	}

	if parenthesis.empty() {
		return true
	}
	return false
}

func main() {

	input := "([)]"

	// [{}]
	if isValid(input) {
		fmt.Println("Valid")
	} else {
		fmt.Println("Not Valid")
	}
}

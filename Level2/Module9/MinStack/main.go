// Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

// Implement the MinStack class:

// MinStack() initializes the stack object.
// void push(int val) pushes the element val onto the stack.
// void pop() removes the element on the top of the stack.
// int top() gets the top element of the stack.
// int getMin() retrieves the minimum element in the stack.
// You must implement a solution with O(1) time complexity for each function.

package main

type MinStack struct {
	stack []int
	min   []int
}

func Constructor() MinStack {
	object := MinStack{
		stack: []int{},
		min:   []int{},
	}

	return object
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

	if len(this.min) == 0 || val <= this.min[len(this.min)-1] {
		this.min = append(this.min, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) != 0 {
		topEle := this.Top()
		if len(this.min) > 0 && topEle == this.min[len(this.min)-1] {
			this.min = this.min[:len(this.min)-1]
		}
		this.stack = this.stack[:len(this.stack)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return -1
	}

	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.min) > 0 {
		return this.min[len(this.min)-1]
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

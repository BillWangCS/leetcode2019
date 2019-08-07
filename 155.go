/*
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
getMin() -- Retrieve the minimum element in the stack.


Example:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> Returns -3.
minStack.pop();
minStack.top();      --> Returns 0.
minStack.getMin();   --> Returns -2.
 */

package main

import "fmt"

type MinStack struct {
	min int
	stack []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	s := make([]int, 0)
	min := 1 << 31 - 1
	return MinStack{min, s}
}


func (this *MinStack) Push(x int)  {
	if x < this.min {
		this.min = x
	}
	this.stack = append(this.stack, x)
}


func (this *MinStack) Pop()  {
	l := len(this.stack)
	if l == 0 {
		return
	} else {
		if this.Top() == this.min {
			this.stack = this.stack[:l-1]
			l--
			this.min = 1 << 31 - 1
			for i := 0; i < l; i++ {
				if this.stack[i] < this.min {
					this.min = this.stack[i]
				}
			}
		} else {
			this.stack = this.stack[:l-1]
		}
	}
}

func heapify(nums *[]int, current int, size int) {
	if current < size {
		left := 2 * current + 1
		right := 2 * current + 2
		max := current
		if left < size {
			if (*nums)[left] > (*nums)[max] {
				max = left
			}
		}
		if right < size {
			if (*nums)[right] > (*nums)[max] {
				max = right
			}
		}
		if current != max {
			(*nums)[max], (*nums)[current] = (*nums)[current], (*nums)[max]
			heapify(nums, max, size)
		}
	}
}

func maxHeapify(nums *[]int, size int) {
	for i := len(*nums) - 1; i >= 0; i-- {
		heapify(nums, i, size)
	}
}

func heapSort(nums []int) []int {
	l := len(nums)
	maxHeapify(&nums, l)
	fmt.Println(nums)
	for i := 0; i < l; i++ {
		nums[l-1-i], nums[0] = nums[0], nums[l-i-1]
		heapify(&nums, 0, l-1-i)
	}
	return nums
}

func main() {
	fmt.Println(heapSort([]int{6,4,2,5,7,1}))
}


func (this *MinStack) Top() int {
	l := len(this.stack)
	if l == 0 {
		return -1 << 32
	} else {
		return this.stack[len(this.stack)-1]
	}
}


func (this *MinStack) GetMin() int {
	return this.min
}

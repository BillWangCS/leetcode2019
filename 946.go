/*
Given two sequences pushed and popped with distinct values,
return true if and only if this could have been the result of
a sequence of push and pop operations on an initially empty stack.



Example 1:

Input: pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
Output: true
Explanation: We might do the following sequence:
push(1), push(2), push(3), push(4), pop() -> 4,
push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
Example 2:

Input: pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
Output: false
Explanation: 1 cannot be popped before 2.


Note:

0 <= pushed.length == popped.length <= 1000
0 <= pushed[i], popped[i] < 1000
pushed is a permutation of popped.
pushed and popped have distinct values.
 */

package main

import "fmt"

type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(elem int) {
	*s = append(*s, elem)
}

func (s *Stack) Pop() int {
	l := len(*s)
	x := (*s)[l-1]
	*s = (*s)[:l-1]
	return x
}

func (s *Stack) Peek() int {
	return (*s)[len(*s)-1]
}

func validateStackSequences(pushed []int, popped []int) bool {
	stack := Stack{}
	cur := 0
	for _, v := range pushed {
		stack.Push(v)
		for !stack.IsEmpty() && stack.Peek() == popped[cur] {
			stack.Pop()
			cur++
		}
	}
	return stack.IsEmpty()
}

func main() {
	fmt.Println(validateStackSequences([]int{1,2,3,4,5}, []int{4,3,5,1,2}))
}
/*
Given a list of daily temperatures T, return a list such that, for each day in the input,
tells you how many days you would have to wait until a warmer temperature.
If there is no future day for which this is possible, put 0 instead.

For example, given the list of temperatures
T = [73, 74, 75, 71, 69, 72, 76, 73], your output should be [1, 1, 4, 2, 1, 1, 0, 0].

Note: The length of temperatures will be in the range [1, 30000].
Each temperature will be an integer in the range [30, 100].
 */

package main

import (
	"fmt"
	"sort"
)

/* Better answer
type Stack []int

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new integer onto the stack
func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

// Pop: remove and return top element of stack, return false if stack is empty
func (s *Stack) Pop() (int) {

	i := len(*s) - 1
	x := (*s)[i]
	*s = (*s)[:i]

	return x
}

// Peek: return top element of stack, return false if stack is empty
func (s *Stack) Peek() (int) {

	i := len(*s) - 1
	x := (*s)[i]

	return x
}

func dailyTemperatures(T []int) []int {
    stk := Stack{}
    n := len(T)
    ans := make([]int, n)
    for i:=0; i < n; i++{

        for !stk.IsEmpty() && T[stk.Peek()] < T[i]{
            cur := stk.Pop()
            ans[cur] = i - cur
        }
        stk.Push(i)
    }

    return ans
}
 */

func removeElements(A *[]int, pos []int) {
	sort.Ints(pos)
	for i := len(pos) - 1; i >= 0; i-- {
		j := pos[i]
		if j == len(*A) - 1 {
			*A = (*A)[:j]
		} else {
			tmp1 := (*A)[:j]
			tmp2 := (*A)[j+1:]
			*A = append(tmp1, tmp2...)
		}
	}
}

func dailyTemperatures(T []int) []int {
	l := len(T)
	res := make([]int, l)
	toDefine := make([]int, 0)
	toDefine = append(toDefine, 0)
	for i := 1; i < l; i++ {
		toRemove := make([]int, 0)
		for j := 0 ; j < len(toDefine); j++ {
			if T[toDefine[j]] < T[i] {
				res[toDefine[j]] = i - toDefine[j]
				toRemove = append(toRemove, j)
			}
		}
		removeElements(&toDefine, toRemove)
		toDefine = append(toDefine, i)
	}
	return res
}

func main() {
	fmt.Println(dailyTemperatures([]int{89,62,70,58,47,47,46,76,100,70}))
}
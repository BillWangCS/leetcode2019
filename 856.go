/*
Given a balanced parentheses string S, compute the score of the string based on the following rule:

() has score 1
AB has score A + B, where A and B are balanced parentheses strings.
(A) has score 2 * A, where A is a balanced parentheses string.


Example 1:

Input: "()"
Output: 1
Example 2:

Input: "(())"
Output: 2
Example 3:

Input: "()()"
Output: 2
Example 4:

Input: "()(())"
2*(1+2*1)
Output: 6


Note:

S is a balanced parentheses string, containing only ( and ).
2 <= S.length <= 50
 */

package main

import "fmt"

type Stack []parenthesesInfo

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(elem parenthesesInfo) {
	*s = append(*s, elem)
}

func (s *Stack) Pop() parenthesesInfo {
	l := len(*s)
	x := (*s)[l-1]
	*s = (*s)[:l-1]
	return x
}

func (s *Stack) Peek() parenthesesInfo {
	return (*s)[len(*s)-1]
}

type parenthesesInfo struct {
	index int
	lastIndex int
}

func scoreOfParentheses(S string) int {
	res := 0
	stk := Stack{}
	l := len(S)
	parenthesesArr := make([]parenthesesInfo, l)
	for i, v := range S {
		if v == '(' {
			stk.Push(parenthesesInfo{index:i})
		} else {
			if !stk.IsEmpty() {
				elem := stk.Pop()
				elem.lastIndex = i
				parenthesesArr[elem.index] = elem
			}
		}
	}
	for i := 0; i < l; i++ {
		if parenthesesArr[i].lastIndex != 0 {
			if parenthesesArr[i].lastIndex - parenthesesArr[i].index == 1 {
				res += 1
			} else {
				res += 2 * scoreOfParentheses(S[parenthesesArr[i].index+1:parenthesesArr[i].lastIndex])
			}
			i = parenthesesArr[i].lastIndex
		}
	}
	return res
}

func main() {
	fmt.Println(scoreOfParentheses("(()(()))"))
}
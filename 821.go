/*
Given a string S and a character C,
return an array of integers representing the shortest distance from the character C in the string.

Example 1:

Input: S = "loveleetcode", C = 'e'
Output: [3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0]


Note:

S string length is in [1, 10000].
C is a single character, and guaranteed to be in string S.
All letters in S and C are lowercase.
 */

package main

import "fmt"

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

func shortestToChar(S string, C byte) []int {
	res := make([]int, 0)
	if len(S) == 0 {
		return res
	}
	pos := make([]int, 0)
	for i := 0; i < len(S); i++ {
		if S[i] == C {
			pos = append(pos, i)
		}
	}
	fmt.Println(pos)
	if len(pos) == 1 {
		p := pos[0]
		for i := 0; i < len(S); i++ {
			res = append(res, abs(i-p))
		}
	} else {
		l := 0
		c := 1
		for i := 0; i < len(S); i++ {
			if i > pos[c] {
				if c < len(pos)-1 {
					c++
				}
				if l < len(pos)-1 {
					l++
				}
			}
			min := pos[l]
			max := pos[c]
			res = append(res, minInt(abs(i-min), abs(i-max)))
		}
	}
	return res
}

func main() {
	fmt.Println(shortestToChar("bbaa", 'b'))
}
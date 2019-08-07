/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true
Example 2:

Input: "()[]{}"
Output: true
Example 3:

Input: "(]"
Output: false
Example 4:

Input: "([)]"
Output: false
Example 5:

Input: "{[]}"
Output: true
 */

package main

import "fmt"

func isValid(s string) bool {
	st := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		l := len(st)
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			st = append(st, byte(s[i]))
		} else {
			if l == 0 {
				return false
			}
			if st[l-1] == '{' && s[i] == '}' || st[l-1] == '[' && s[i] == ']' || st[l-1] == '(' && s[i] == ')' {
				st = st[:l-1]
			} else {
				return false
			}
		}
	}
	if len(st) > 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isValid("()"))
}
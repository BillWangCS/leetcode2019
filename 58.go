/*
Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the length of last word in the string.

If the last word does not exist, return 0.

Note: A word is defined as a character sequence consists of non-space characters only.

Example:

Input: "Hello World"
Output: 5
 */

package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	strs := strings.Split(s, " ")
	if len(strs) == 0 {
		return 0
	}
	return len(strs[len(strs)-1])
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
}
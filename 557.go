/*
Given a string, you need to reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.

Example 1:
Input: "Let's take LeetCode contest"
Output: "s'teL ekat edoCteeL tsetnoc"
Note: In the string, each word is separated by single space and there will not be any extra space in the string.
 */

package main

import (
	"fmt"
	"strings"
)

func reverseWord(s string) string {
	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}

func reverseWords(s string) string {
	strs := strings.Split(s, " ")
	for i, str := range strs {
		strs[i] = reverseWord(str)
	}
	return strings.Join(strs, " ")
}

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}
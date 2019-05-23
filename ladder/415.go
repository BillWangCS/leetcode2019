/*
Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

Example
Example 1:

Input: "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama"
Example 2:

Input: "race a car"
Output: false
Explanation: "raceacar"
Challenge
O(n) time without extra memory.

Notice
Have you consider that the string might be empty? This is a good question to ask during an interview.
 */

package main

import (
	"fmt"
	"strings"
)

func isAlphanumeric(c byte) bool {
	if (c <= 'z' && c >= 'a') || (c >= '0' && c <= '9') {
		return true
	}
	return false
}

func isPalindrome (s string) bool {
	// write your code here
	i := 0
	j := len(s) - 1
	s = strings.ToLower(s)
	for i < j {
		for i < j && !isAlphanumeric(byte(s[i])) {
			i++
		}
		for i < j && !isAlphanumeric(byte(s[j])) {
			j--
		}
		if s[i] != s[j] {
			return false
		}
		if i < j {
			i++
			j--
		}
	}
	return true
}

func main() {
	//fmt.Println(isAlphanumeric(''))
	fmt.Println(isPalindrome(" "))
}

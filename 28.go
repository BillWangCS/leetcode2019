/*
Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1
Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().
 */

package main

import "fmt"

func strStr(haystack string, needle string) int {
	l := len(needle)
	if l == 0 {
		return 0
	}
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			if i+l <= len(haystack) && haystack[i:i+len(needle)] == needle {
				return i
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("hello", "llo"))
}
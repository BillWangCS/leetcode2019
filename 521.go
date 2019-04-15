/*
Given a group of two strings, you need to find the longest uncommon subsequence of this group of two strings.
The longest uncommon subsequence is defined as the longest subsequence of
one of these strings and this subsequence should not be any subsequence of the other strings.

A subsequence is a sequence that can be derived from one sequence by deleting some characters without
changing the order of the remaining elements.
Trivially, any string is a subsequence of itself and an empty string is a subsequence of any string.

The input will be two strings, and the output needs to be the length of the longest uncommon subsequence.
If the longest uncommon subsequence doesn't exist, return -1.

Example 1:
Input: "aba", "cdc"
Output: 3
Explanation: The longest uncommon subsequence is "aba" (or "cdc"),
because "aba" is a subsequence of "aba",
but not a subsequence of any other strings in the group of two strings.
Note:

Both strings' lengths will not exceed 100.
Only letters from a ~ z will appear in input strings.
 */

package main

import "fmt"

func findLUSlength(a string, b string) int {
	if len(a) > len(b) {
		return len(a)
	} else if len(b) > len(a) {
		return len(b)
	}
	if a == b {
		return -1
	}
	head := 0
	tail := len(a) - 1
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			head = i
			break
		}
	}
	for i := tail; i >= 0; i-- {
		if a[i] != b[i] {
			tail = i
			break
		}
	}
	return tail - head + 1
}

func main() {
	fmt.Println(findLUSlength("aba", "cdc"))
}
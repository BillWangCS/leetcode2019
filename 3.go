/*
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
 */

package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	longest := 0
	dict := make(map[byte]bool)
	str := ""
	l := 0
	r := 0
	for r < len(s) {
		if !dict[byte(s[r])] {
			dict[byte(s[r])] = true
			r++
			if r - l > longest {
				longest = r - l
				str = s[l:r]
			}
		} else {
			for l < r {
				if s[l] == s[r] {
					break
				}
				dict[byte(s[l])] = false
				l++
			}
			if l < r {
				l++
			}
			r++
		}
	}
	fmt.Println(str)
	return longest
}

func main() {
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
}

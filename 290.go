/*
Given a pattern and a string str, find if str follows the same pattern.

Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in str.

Example 1:

Input: pattern = "abba", str = "dog cat cat dog"
Output: true
Example 2:

Input:pattern = "abba", str = "dog cat cat fish"
Output: false
Example 3:

Input: pattern = "aaaa", str = "dog cat cat dog"
Output: false
Example 4:

Input: pattern = "abba", str = "dog dog dog dog"
Output: false
Notes:
You may assume pattern contains only lowercase letters, and str contains lowercase letters that may be separated by a single space.
 */

package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, str string) bool {
	dict1 := make(map[uint8]string)
	dict2 := make(map[string]uint8)
	str = strings.Trim(str, " ")
	strs := strings.Split(str, " ")
	if len(strs) != len(pattern) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		if _, ok := dict1[pattern[i]]; !ok {
			dict1[pattern[i]] = strs[i]
		} else if dict1[pattern[i]] != strs[i]{
			return false
		}
	}
	for i := 0; i < len(pattern); i++ {
		if _, ok := dict2[strs[i]]; !ok {
			dict2[strs[i]] = pattern[i]
		} else if dict2[strs[i]] != pattern[i]{
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dodg"))
}
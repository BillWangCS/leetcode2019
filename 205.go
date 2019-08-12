/*
Given two strings s and t, determine if they are isomorphic.

Two strings are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters.
No two characters may map to the same character but a character may map to itself.

Example 1:

Input: s = "egg", t = "add"
Output: true
Example 2:

Input: s = "foo", t = "bar"
Output: false
Example 3:

Input: s = "paper", t = "title"
Output: true
Note:
You may assume both s and t have the same length.
 */

package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	dict := make(map[string]string)
	for i := 0; i < len(s); i++ {
		if _, ok := dict[string(s[i])]; !ok {
			dict[string(s[i])] = string(t[i])
		} else {
			if string(t[i]) != dict[string(s[i])] {
				return false
			}
		}
	}
	dict2 := make(map[string]string)
	for i := 0; i < len(s); i++ {
		if _, ok := dict2[string(t[i])]; !ok {
			dict2[string(t[i])] = string(s[i])
		} else {
			if string(s[i]) != dict2[string(t[i])] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isIsomorphic("foo", "baa"))
}
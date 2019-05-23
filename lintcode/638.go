/*
Given two strings s and t, determine if they are isomorphic.

Two strings are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters.
No two characters may map to the same character but a character may map to itself.

Example
Example 1:

Input : s = "egg", t = "add"
Output : true
Explanation :
e -> a, g -> d.
Example 2:

Input : s = "foo", t = "bar"
Output : false
Explanation :
No solution.
 */

package main

import "fmt"

func isIsomorphic (s string, t string) bool {
	// write your code here
	ls := len(s)
	lt := len(t)
	if ls != lt {
		return false
	}
	dicts := make(map[byte]byte)
	for i := 0; i < ls; i++ {
		if _, ok := dicts[byte(s[i])]; !ok {
			dicts[byte(s[i])] = byte(t[i])
		} else if dicts[byte(s[i])] != byte(t[i]) {
			return false
		}
	}
	dictt := make(map[byte]byte)
	for i := 0; i < lt; i++ {
		if _, ok := dictt[byte(t[i])]; !ok {
			dictt[byte(t[i])] = byte(s[i])
		} else if dictt[byte(t[i])] != byte(s[i]) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isIsomorphic("NYLM", "NYLN"))
}
/*
Given two strings s and t , write a function to determine if t is an anagram of s.

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false
Note:
You may assume the string contains only lowercase alphabets.
 */

package main

func isAnagram(s string, t string) bool {
	ls := len(s)
	lt := len(t)
	if ls != lt {
		return false
	}
	dicts := make(map[uint8]int)
	for i := 0; i < ls; i++ {
		dicts[s[i]]++
	}
	dictt := make(map[uint8]int)
	for i := 0; i < lt; i++ {
		dictt[t[i]]++
	}
	for i := 0; i < ls; i++ {
		if dictt[s[i]] != dicts[s[i]] {
			return false
		}
	}
	for i := 0; i < lt; i++ {
		if dictt[t[i]] != dicts[t[i]] {
			return false
		}
	}
	return true
}

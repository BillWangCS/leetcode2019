/*
Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.

Examples:

s = "leetcode"
return 0.

s = "loveleetcode",
return 2.
Note: You may assume the string contain only lowercase letters.
 */

package main

import "fmt"

func firstUniqChar(s string) int {
	dict := make(map[byte]int)
	for _, v := range s {
		dict[byte(v)]++
	}
	for i, v := range s {
		if dict[byte(v)] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
}
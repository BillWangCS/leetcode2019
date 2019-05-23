/*
Given a string which consists of lowercase or uppercase letters,
find the length of the longest palindromes that can be built with those letters.

This is case sensitive, for example "Aa" is not considered a palindrome here.

Input : s = "abccccdd"
Output : 7
Explanation :
One longest palindrome that can be built is "dccaccd", whose length is `7`.
 */

package main

import "fmt"

func longestPalindrome (s string) int {
	res := 0
	dict := make(map[string]int)
	for _, v := range s {
		dict[string(v)]++
	}
	flag := false
	for _, v := range dict {
		if v % 2 == 0 {
			res += v
		} else {
			flag = true
			res += v - 1
		}
	}
	if flag {
		res+=1
	}
	return res
}

func main() {
	fmt.Println(longestPalindrome("NTrQdQGgwtxqRTSBOitAXUkwGLgUHtQOmYMwZlUxqZysKpZxRoehgirdMUgy"))
}
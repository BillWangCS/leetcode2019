/*
Given a string S, find the longest palindromic substring in S.
You may assume that the maximum length of S is 1000, and there exists one unique longest palindromic substring.

Example
Example 1:

Input:"abcdzdcab"
Output:"cdzdc"
Example 2:

Input:"aba"
Output:"aba"
Challenge
O(n2) time is acceptable. Can you do it in O(n) time.
 */

package main

import "fmt"

func palindromeInfo(pos1 int, pos2 int, s string) (int, string) {
	p := pos2
	f := pos1
	for f >= 0 && p < len(s) {
		if s[p] != s[f] {
			break
		}
		p++
		f--
	}
	fmt.Println(pos1, pos2, f, p)
	return p - f - 1, s[f+1:p]
}

func longestPalindrome (s string) string {
	// write your code here
	if len(s) == 1 {
		return s
	}
	res := ""
	max := 0
	for i := 0; i < len(s); i++ {
		l := 0
		str := ""
		if i + 1 < len(s) && s[i] == s[i+1] {
			m := 0
			for s[i] == s[i+m] {
				m++
				if i + m == len(s) {
					break
				}
			}
			l, str = palindromeInfo(i, i+m-1, s)
		} else {
			l, str = palindromeInfo(i, i, s)
		}
		if l > max {
			max = l
			res = str
		}
	}
	return res
}

func main() {
	fmt.Println(longestPalindrome("edccccd"))
}
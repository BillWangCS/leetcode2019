/*
Given a string, your task is to count how many palindromic substrings in this string.

The substrings with different start indexes or end indexes are counted as different substrings even they consist of same characters.

Example 1:

Input: "abc"
Output: 3
Explanation: Three palindromic strings: "a", "b", "c".


Example 2:

Input: "aaa"
Output: 6
Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".


Note:

The input string length won't exceed 1000.
 */

 /*
 Approach #1: Expand Around Center [Accepted]

Intuition

Let N be the length of the string. The middle of the palindrome could be in one of 2N - 1 positions:
 either at letter or between two letters.

For each center, let's count all the palindromes that have this center.
 Notice that if [a, b] is a palindromic interval (meaning S[a], S[a+1], ..., S[b] is a palindrome), then [a+1, b-1] is one too.

Algorithm

For each possible palindrome center, let's expand our candidate palindrome on the interval [left, right] as long as we can.
 The condition for expanding is left >= 0 and right < N and S[left] == S[right].
 That means we want to count a new palindrome S[left], S[left+1], ..., S[right].
  */

 /*
 class Solution {
    public int countSubstrings(String S) {
        int N = S.length(), ans = 0;
        for (int center = 0; center <= 2*N-1; ++center) {
            int left = center / 2;
            int right = left + center % 2;
            while (left >= 0 && right < N && S.charAt(left) == S.charAt(right)) {
                ans++;
                left--;
                right++;
            }
        }
        return ans;
    }
}
  */

package main

import "fmt"

func countSubstrings (s string) int {
	// write your code here
	l := len(s)
	res := 0
	for center := 0; center <= 2 * l - 1; center++ {
		left := center / 2
		right := left + center%2
		for left >= 0 && right < l && s[left] == s[right] {
			res++
			left--
			right++
		}
	}
	return res
}

func main() {
	fmt.Println(countSubstrings("baaabaaab"))
}
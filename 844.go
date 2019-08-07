/*
Given two strings S and T, return if they are equal when both are typed into empty text editors. # means a backspace character.

Example 1:

Input: S = "ab#c", T = "ad#c"
Output: true
Explanation: Both S and T become "ac".
Example 2:

Input: S = "ab##", T = "c#d#"
Output: true
Explanation: Both S and T become "".
Example 3:

Input: S = "a##c", T = "#a#c"
Output: true
Explanation: Both S and T become "c".
Example 4:

Input: S = "a#c", T = "b"
Output: false
Explanation: S becomes "c" while T becomes "b".
Note:

1 <= S.length <= 200
1 <= T.length <= 200
S and T only contain lowercase letters and '#' characters.
Follow up:

Can you solve it in O(N) time and O(1) space?
 */

package main

import "fmt"

func backspaceCompare(S string, T string) bool {
	s1 := ""
	s2 := ""
	for i := 0; i < len(S); i++ {
		if string(S[i]) == "#" {
			if len(s1) == 0 {
				continue
			}
			s1 = s1[:len(s1)-1]
		} else {
			s1 += string(S[i])
		}
	}
	for i := 0; i < len(T); i++ {
		if string(T[i]) == "#" {
			if len(s2) == 0 {
				continue
			}
			s2 = s2[:len(s2)-1]
		} else {
			s2 += string(T[i])
		}
	}
	return s1 == s2
}

func main() {
	fmt.Println(backspaceCompare("a##c", "#a#c"))
}
/*
Given a string S,
return the "reversed" string where all characters that are not a letter stay in the same place,
and all letters reverse their positions.



Example 1:

Input: "ab-cd"
Output: "dc-ba"
Example 2:

Input: "a-bC-dEf-ghIj"
Output: "j-Ih-gfE-dCba"
Example 3:

Input: "Test1ng-Leet=code-Q!"
Output: "Qedo1ct-eeLg=ntse-T!"


Note:

S.length <= 100
33 <= S[i].ASCIIcode <= 122
S doesn't contain \ or "
 */

package main

import (
	"fmt"
	"strings"
)

func isAlpha(s string) bool {
	ss := strings.ToLower(s)
	if ss >= "a" && ss <= "z" {
		return true
	}
	return false
}

func reverseOnlyLetters(S string) string {
	strs := strings.Split(S, "")
	i := 0
	j := len(strs) - 1
	for i <= j {
		for i <= j && !isAlpha(strs[i]) {
			i++
		}
		for i <= j && !isAlpha(strs[j]) {
			j--
		}
		if i <= j {
			strs[i], strs[j] = strs[j], strs[i]
			i++
			j--
		}
	}
	return strings.Join(strs, "")
}

func main() {
	fmt.Println(isAlpha("]"))
	fmt.Println(reverseOnlyLetters("7_28]"))
}
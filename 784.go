/*
Given a string S, we can transform every letter individually to be lowercase or uppercase to create another string.
Return a list of all possible strings we could create.

Examples:
Input: S = "a1b2"
Output: ["a1b2", "a1B2", "A1b2", "A1B2"]

Input: S = "3z4"
Output: ["3z4", "3Z4"]

Input: S = "12345"
Output: ["12345"]
Note:

S will be a string with length between 1 and 12.
S will consist only of letters or digits.
 */

package main

import (
	"fmt"
	"strings"
)

func dfs_784(res *[]string, index int, strs []string) {
	*res = append(*res, strings.Join(strs, ""))
	for i := index; i < len(strs); i++ {
		if strs[i] >= "a" && strs[i] <= "z" {
			tmp := strs[i]
			strs[i] = strings.ToUpper(strs[i])
			dfs_784(res, i + 1, strs)
			strs[i] = tmp
		} else if strs[i] >= "A" && strs[i] <= "Z" {
			tmp := strs[i]
			strs[i] = strings.ToLower(strs[i])
			dfs_784(res, i + 1, strs)
			strs[i] = tmp
		}
	}
}

func letterCasePermutation(S string) []string {
	res := make([]string, 0)
	strs := strings.Split(S, "")
	dfs_784(&res, 0, strs)
	return res
}

func main() {
	fmt.Println(letterCasePermutation("a1b2"))
}
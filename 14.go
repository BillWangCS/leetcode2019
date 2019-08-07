/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
Note:

All given inputs are in lowercase letters a-z.
 */

package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	res := ""
	if len(strs) == 0 {
		return res
	}
	for j := 0; j < len(strs[0]); j++ {
		for i := 1; i< len(strs); i++ {
			if len(strs[i]) <= j || strs[i][j] != strs[0][j] {
				return res
			}
		}
		res += string(strs[0][j])
	}
	return res
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
}

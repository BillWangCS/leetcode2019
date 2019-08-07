/*
Given a string and an integer k, you need to reverse the first k characters for every 2k characters counting
from the start of the string. If there are less than k characters left,
reverse all of them. If there are less than 2k but greater than or equal to k characters,
then reverse the first k characters and left the other as original.
Example:
Input: s = "abcdefg", k = 2
Output: "bacdfeg"
Restrictions:
The string consists of lower English letters only.
Length of the given string and k will in the range [1, 10000]
 */

package main

import "fmt"

func reverseArr(A *[]string) {
	i := 0
	j := len(*A) - 1
	for i <= j {
		(*A)[i], (*A)[j] = (*A)[j], (*A)[i]
		i++
		j--
	}
}

func reverseStr(s string, k int) string {
	res := ""
	arr := make([]string, 0)
	arrs := make([][]string, 0)
	for i := 0; i < len(s); i++ {
		if i % k == 0 {
			if len(arr) != 0 {
				arrs = append(arrs, arr)
			}
			arr = make([]string, 0)
			arr = append(arr, string(s[i]))
		} else {
			arr = append(arr, string(s[i]))
		}
	}
	if len(arr) != 0 {
		arrs = append(arrs, arr)
	}
	for i := 0; i < len(arrs); i++ {
		if i % 2 == 0 {
			reverseArr(&arrs[i])
		}
	}
	for i := 0; i < len(arrs); i++ {
		for j := 0; j < len(arrs[i]); j++ {
			res += arrs[i][j]
		}
	}
	return res
}

func main() {
	fmt.Println(reverseStr("abcbababc", 3))
}
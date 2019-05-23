/*
In a string S of lowercase letters, these letters form consecutive groups of the same character.

For example, a string like S = "abbxxxxzyy" has the groups "a", "bb", "xxxx", "z" and "yy".

Call a group large if it has 3 or more characters.  We would like the starting and ending positions of every large group.

The final answer should be in lexicographic order.



Example 1:

Input: "abbxxxxzzy"
Output: [[3,6]]
Explanation: "xxxx" is the single large group with starting  3 and ending positions 6.
Example 2:

Input: "abc"
Output: []
Explanation: We have "a","b" and "c" but no large group.
Example 3:

Input: "abcdddeeeeaabbbcd"
Output: [[3,5],[6,9],[12,14]]


Note:  1 <= S.length <= 1000
 */

package main

import "fmt"

func largeGroupPositions(S string) [][]int {
	res := make([][]int, 0)
	l := len(S)
	if l == 0 {
		return res
	}
	i := 0
	j := 0
	for i < l {
		x := S[i]
		cnt := 0
		for j = i; j < l; j++ {
			if x == S[j] {
				cnt++
			} else {
				break
			}
		}
		if cnt >= 3 {
			tmp := make([]int, 0)
			tmp = append(tmp, i)
			tmp = append(tmp, j-1)
			res = append(res, tmp)
		}
		i = j
	}
	return res
}

func main() {
	fmt.Println(largeGroupPositions("abcdddeeeeaabbb"))
}
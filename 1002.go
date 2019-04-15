/*
Given an array A of strings made only from lowercase letters,
return a list of all characters that show up in all strings within the list (including duplicates).
For example, if a character occurs 3 times in all strings but not 4 times,
you need to include that character three times in the final answer.

You may return the answer in any order.



Example 1:

Input: ["bella","label","roller"]
Output: ["e","l","l"]
Example 2:

Input: ["cool","lock","cook"]
Output: ["c","o"]


Note:

1 <= A.length <= 100
1 <= A[i].length <= 100
A[i][j] is a lowercase letter
 */

package main

import "fmt"

func commonChars(A []string) []string {
	res := make([]string, 0)
	strMap := make(map[string]map[string]int)
	for _, v := range A {
		cMap := make(map[string]int)
		for _, c := range v {
			cMap[string(c)] ++
		}
		strMap[v] = cMap
	}

	for i := 'a'; i <= 'z'; i++ {
		min := 101
		for _, v := range A {
			if strMap[v][string(i)] < min {
				min = strMap[v][string(i)]
			}
		}
		for j := 0; j < min; j++ {
			res = append(res, string(i))
		}
	}
	return res
}

func main() {
	fmt.Println(commonChars([]string{"cool","lock","cook"}))
}

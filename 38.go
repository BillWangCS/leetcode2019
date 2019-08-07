/*
The count-and-say sequence is the sequence of integers with the first five terms as following:

1.     1
2.     11
3.     21
4.     1211
5.     111221
1 is read off as "one 1" or 11.
11 is read off as "two 1s" or 21.
21 is read off as "one 2, then one 1" or 1211.

Given an integer n where 1 ≤ n ≤ 30, generate the nth term of the count-and-say sequence.

Note: Each term of the sequence of integers will be represented as a string.



Example 1:

Input: 1
Output: "1"
Example 2:

Input: 4
Output: "1211"
 */

package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	l := 1
	s := "1"
	current := ""
	for l < n {
		i := 0
		j := 0
		cnt := 0
		for j < len(s) {
			if s[j] == s[i] {
				cnt++
			} else {
				current += strconv.Itoa(cnt) + string(s[i])
				i = j
				cnt = 0
				continue
			}
			j++
		}
		current += strconv.Itoa(cnt) + string(s[i])
		s = current
		current = ""
		l++
	}
	return s
}

func main() {
	fmt.Println(countAndSay(5))
}
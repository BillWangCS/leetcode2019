/*
A message containing letters from A-Z is being encoded to numbers using the following mapping:

'A' -> 1
'B' -> 2
...
'Z' -> 26
Given an encoded message containing digits, determine the total number of ways to decode it.

Example
Example 1:

Input: "12"
Output: 2
Explanation: It could be decoded as AB (1 2) or L (12).
Example 2:

Input: "10"
Output: 1
 */

package main

import "fmt"

//Better solution 爬梯子问题的变种
func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	count := 1
	prvCount := 1
	for i := 0; i < len(s); i++ {
		newCount := 0
		if i > 0 && (s[i-1] == '1' || (s[i-1] == '2' && s[i] < '7')) {
			newCount += prvCount
		}
		if s[i] > '0' {
			newCount += count
		}
		prvCount = count
		count = newCount
	}
	return count
}

/*func numDecodingsHelper (s string) (int, bool) {
	// write your code here
	res := 0
	flag1 := false
	flag2 := false
	if len(s) == 0 {
		return 0, false
	}
	if len(s) == 1 {
		if s == "0" {
			return 0, false
		} else {
			return 1, true
		}
	}
	if len(s) == 2 {
		if s == "10" || s == "20" {
			return 1, true
		} else if s > "10" && s<= "26" {
			return 2, true
		}
	}
	if s[:2] >= "10" && s[:2] <= "26" {
		r, flag := numDecodingsHelper(s[2:])
		if flag {
			res += r
		}
		flag1 = flag
	}
	if s[:1] >= "1" && s[:1] <= "9" && s[:2] != "10" && s[:2] != "20" {
		r, flag := numDecodingsHelper(s[1:])
		if flag {
			res += r
		}
		flag2 = flag
	}
	return res, flag1 || flag2
}

func numDecodings (s string) int {
	res, flag := numDecodingsHelper(s)
	if !flag {
		return 0
	}
	return res
}
*/
func main() {
	//s := "1001"
	//fmt.Println(s[:2]>="10", s[:2]<="26")
	fmt.Println(numDecodings("12300"))
}

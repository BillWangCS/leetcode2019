/*
Given an integer, return its base 7 string representation.

Example 1:
Input: 100
Output: "202"
Example 2:
Input: -7
Output: "-10"
Note: The input will be in range of [-1e7, 1e7].
 */

package main

import (
	"fmt"
	"strconv"
)

func convertToBase7(num int) string {
	str := ""
	if num == 0 {
		return "0"
	}
	flag := false
	if num < 0 {
		num = -num
		flag = true
	}
	for num > 0 {
		str = strconv.Itoa(num % 7) + str
		num = num / 7
	}
	if flag {
		str = "-" + str
	}
	return str
}

func main() {
	fmt.Println(convertToBase7(-99))
}

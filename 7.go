/*
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output: 321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21
Note:
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range:
[−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
 */

package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var n int64
	n = int64(x)
	var t int64
	flag := false
	if n == math.MinInt32 {
		return 0
	}
	if n < 0 {
		t = -n
		flag = true
	} else {
		t = n
	}
	var sum int64
	fmt.Println(t)
	for t > 0 {
		tmp := t % 10
		sum = 10 * sum + tmp
		t = t / 10
	}
	if flag {
		sum = -1 * sum
	}
	if sum > math.MaxInt32 || sum < math.MinInt32 {
		return 0
	}
	return int(sum)
}

func main() {
	fmt.Println(reverse(-100))
}
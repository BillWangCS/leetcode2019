/*
Given a number N, return a string consisting of "0"s and "1"s that represents its value in base -2 (negative two).

The returned string must have no leading zeroes, unless the string is "0".



Example 1:

Input: 2
Output: "110"
Explantion: (-2) ^ 2 + (-2) ^ 1 = 2
Example 2:

Input: 3
Output: "111"
Explantion: (-2) ^ 2 + (-2) ^ 1 + (-2) ^ 0 = 3
Example 3:

Input: 4
Output: "100"
Explantion: (-2) ^ 2 = 4


Note:

0 <= N <= 10^9
 */

package main

import "fmt"

func getBits(n int) string {
	res := ""
	if n == 0 {
		return "0"
	}
	for n > 0 {
		if n % 2 == 0 {
			res = "0" + res
		} else {
			res = "1" + res
		}
		n /= 2
	}
	return res
}

func baseNeg2(N int) string {
	negBits := make([]int, 0)
	var i uint32
	for i = 1; i < 33; i+=2 {
		negBits = append(negBits, 1<<i)
	}
	for _, v := range negBits {
		if v&N != 0 {
			N += v<<1
		}
	}
	return getBits(N)
}

func main() {
	fmt.Println(baseNeg2(6))
}
/*
Given a positive integer, output its complement number.
The complement strategy is to flip the bits of its binary representation.

Note:
The given integer is guaranteed to fit within the range of a 32-bit signed integer.
You could assume no leading zero bit in the integer’s binary representation.
Example 1:
Input: 5
Output: 2
Explanation: The binary representation of 5 is 101 (no leading zero bits), and its complement is 010.
So you need to output 2.
Example 2:
Input: 1
Output: 0
Explanation: The binary representation of 1 is 1 (no leading zero bits), and its complement is 0.
So you need to output 0.
 */

package main

import (
	"fmt"
	"strconv"
)

func findComplement(num int) int {
	res := 0
	s := ""
	for num > 0 {
		n := num % 2
		if n == 1 {
			n = 0
		} else {
			n = 1
		}
		s = strconv.Itoa(n) + s
		num = num >> 1
	}
	d := 0
	for _, v := range s {
		if v == '1' {
			d = 1
		} else {
			d = 0
		}
		res = res << 1 + d
	}
	return res
}

func main() {
	fmt.Println(findComplement(5))
}
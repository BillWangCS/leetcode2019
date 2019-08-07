/*
Given two non-negative integers num1 and num2 represented as string, return the sum of num1 and num2.

Note:

The length of both num1 and num2 is < 5100.
Both num1 and num2 contains only digits 0-9.
Both num1 and num2 does not contain any leading zero.
You must not use any built-in BigInteger library or convert the inputs to integer directly.
 */

package main

import (
	"fmt"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	l := len(num1)
	s := len(num2)
	c := 0
	res := ""
	if l < s {
		s, l = l, s
	}
	for i := 1; i <= l; i++ {
			sum := 0
			if i <= len(num1) {
				n1, _ := strconv.Atoi(string(num1[len(num1)-i]))
				sum += n1
			}
			if i <= len(num2) {
				n2, _ := strconv.Atoi(string(num2[len(num2)-i]))
				sum += n2
			}
			sum += c
			res = strconv.Itoa(sum%10) + res
			c = sum / 10
	}
	if c > 0 {
		res = strconv.Itoa(c) + res
	}
	return res
}

func main() {
	fmt.Println(addStrings("999", "1"))
}

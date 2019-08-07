/*
Given two binary strings, return their sum (also a binary string).

The input strings are both non-empty and contains only characters 1 or 0.

Example 1:

Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101"
 */

package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	c := 0
	res := ""
	for i >= 0 || j >= 0 {
		s := 0
		if i >= 0 {
			 if a[i] == '1' {
			 	s += 1
			 }
		}
		if j >= 0 {
			if b[j] == '1' {
				s += 1
			}
		}
		s += c
		c = s / 2
		res = strconv.Itoa(s%2) + res
		i--
		j--
	}
	if c > 0 {
		res = "1" + res
	}
	return res
}

func main() {
	fmt.Println(addBinary("1010", "1011"))
}
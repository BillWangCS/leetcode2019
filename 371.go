/*
Calculate the sum of two integers a and b, but you are not allowed to use the operator + and -.

Example 1:

Input: a = 1, b = 2
Output: 3
Example 2:

Input: a = -2, b = 3
Output: 1
 */

package main

import "fmt"

func getSum(a int, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	res := 0
	carry := a&b << 1
	sum := a^b
	for carry != 0 {
		res = sum^carry
		carry = sum&carry << 1
		sum = res
	}
	return sum
}

func main() {
	fmt.Println(getSum(1,2))
}
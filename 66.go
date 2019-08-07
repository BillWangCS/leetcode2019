/*
Given a non-empty array of digits representing a non-negative integer, plus one to the integer.

The digits are stored such that the most significant digit is at the head of the list,
and each element in the array contain a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.

Example 1:

Input: [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Example 2:

Input: [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
 */

package main

import "fmt"

func plusOne(digits []int) []int {
	res := make([]int, 0)
	c := 0
	flag := false
	for i := len(digits) - 1; i >= 0; i-- {
		if !flag {
			tmp := digits[i] + c + 1
			c = tmp / 10
			digits[i] = tmp % 10
			flag = true
		} else {
			tmp := c + digits[i]
			fmt.Println(tmp)
			if tmp >= 10 {
				c = tmp / 10
				digits[i] = tmp % 10
			} else {
				digits[i] = tmp
				c = 0
				break
			}
		}
	}
	if c > 0 {
		res = append(res, 1)
		res = append(res, digits...)
		return res
	}
	return digits
}

func main() {
	fmt.Println(plusOne([]int{9,9,9,9}))
}
/*
Given a positive integer, check whether it has alternating bits: namely,
if two adjacent bits will always have different values.

Example 1:
Input: 5
Output: True
Explanation:
The binary representation of 5 is: 101
Example 2:
Input: 7
Output: False
Explanation:
The binary representation of 7 is: 111.
Example 3:
Input: 11
Output: False
Explanation:
The binary representation of 11 is: 1011.
Example 4:
Input: 10
Output: True
Explanation:
The binary representation of 10 is: 1010.
 */

package main

import "fmt"

func hasAlternatingBits(n int) bool {
	last := n & 1
	n = n >> 1
	for n > 0 {
		if last ^ (n & 1) == 0 {
			return false
		}
		last = n & 1
		n = n >> 1
	}
	return true
}

func main() {
	fmt.Println(hasAlternatingBits(11))
}
/*
The Hamming distance between two integers is the number of positions at which the corresponding bits are different.

Given two integers x and y, calculate the Hamming distance.

Note:
0 ≤ x, y < 231.

Example:

Input: x = 1, y = 4

Output: 2

Explanation:
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑

The above arrows point to positions where the corresponding bits are different.
 */

package main

import "fmt"

func hammingDistance(x int, y int) int {
	res := 0
	n := x^y
	for n > 0 {
		if n % 2 == 1 {
			res ++
		}
		n = n >> 1
	}
	return res
}

func main() {
	fmt.Println(hammingDistance(1, 4))
}

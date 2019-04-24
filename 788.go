/*
X is a good number if after rotating each digit individually by 180 degrees,
we get a valid number that is different from X.  Each digit must be rotated - we cannot choose to leave it alone.

A number is valid if each digit remains a digit after rotation. 0, 1,
and 8 rotate to themselves; 2 and 5 rotate to each other; 6 and 9 rotate to each other,
and the rest of the numbers do not rotate to any other number and become invalid.

Now given a positive number N, how many numbers X from 1 to N are good?

Example:
Input: 10
Output: 4
Explanation:
There are four good numbers in the range [1, 10] : 2, 5, 6, 9.
Note that 1 and 10 are not good numbers, since they remain unchanged after rotating.
Note:

N  will be in range [1, 10000].
 */

package main

import "fmt"

func isGood(n int) bool {
	t := n
	for t > 0 {
		if t % 10 == 3 || t % 10 == 4 || t % 10 == 7 {
			return false
		}
		t = t / 10
	}
	for n > 0 {
		if n % 10 == 2 || n % 10 == 5 || n % 10 == 6 || n % 10 == 9 {
			return true
		}
		n = n / 10
	}
	return false
}

func rotatedDigits(N int) int {
	res := 0
	for i := 1; i <= N; i++ {
		if isGood(i) {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(isGood(36))
}
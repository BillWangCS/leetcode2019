/*
Given an array A of integers,
return true if and only if we can partition the array into three non-empty parts with equal sums.

Formally, we can partition the array if we can find indexes
i+1 < j with (A[0] + A[1] + ... + A[i] == A[i+1] + A[i+2] + ... + A[j-1] == A[j] + A[j-1] + ... + A[A.length - 1])



Example 1:

Input: [0,2,1,-6,6,-7,9,1,2,0,1]
Output: true
Explanation: 0 + 2 + 1 = -6 + 6 - 7 + 9 + 1 = 2 + 0 + 1
Example 2:

Input: [0,2,1,-6,6,7,9,-1,2,0,1]
Output: false
Example 3:

Input: [3,3,6,5,-2,2,5,1,-9,4]
Output: true
Explanation: 3 + 3 = 6 = 5 - 2 + 2 + 5 + 1 - 9 + 4


Note:

3 <= A.length <= 50000
-10000 <= A[i] <= 10000
 */

package main

import "fmt"

func canThreePartsEqualSum(A []int) bool {
	sum := 0
	for _, v := range A {
		sum += v
	}
	if sum % 3 != 0 {
		return false
	}
	l := len(A)
	subSum := sum / 3
	i := 0
	j := l - 1
	s1 := 0
	for i < l {
		s1 += A[i]
		if s1 == subSum {
			break
		}
		i++
	}
	s2 := 0
	for j >= 0 {
		s2 += A[j]
		if s2 == subSum {
			break
		}
		j--
	}
	fmt.Println(i, j, subSum)
	if i + 1 < j {
		return true
	}
	return false
}

func main() {
	fmt.Println(canThreePartsEqualSum([]int{3,3,6,5,-2,2,5,1,-9,4}))
}
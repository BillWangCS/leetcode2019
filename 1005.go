/*
Given an array A of integers, we must modify the array in the following way:
we choose an i and replace A[i] with -A[i], and we repeat this process K times in total.
(We may choose the same index i multiple times.)

Return the largest possible sum of the array after modifying it in this way.



Example 1:

Input: A = [4,2,3], K = 1
Output: 5
Explanation: Choose indices (1,) and A becomes [4,-2,3].
Example 2:

Input: A = [3,-1,0,2], K = 3
Output: 6
Explanation: Choose indices (1, 2, 2) and A becomes [3,1,0,2].
Example 3:

Input: A = [2,-3,-1,5,-4], K = 2
Output: 13
Explanation: Choose indices (1, 4) and A becomes [2,3,-1,5,4].


Note:

1 <= A.length <= 10000
1 <= K <= 10000
-100 <= A[i] <= 100
 */

package main

import (
	"fmt"
	"sort"
)

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func largestSumAfterKNegations(A []int, K int) int {
	if len(A) == 0 {
		return 0
	}
	positive := make([]int, 0)
	negative := make([]int, 0)
	for _, v := range A {
		if v < 0 {
			negative = append(negative, v)
		} else {
			positive = append(positive, v)
		}
	}
	sort.Ints(negative)
	for i := 0; i < len(negative); i++ {
		if K <= 0 {
			break
		}
		negative[i] = -negative[i]
		K--
	}
	sum := 0
	for i := 0; i < len(positive); i++ {
		sum += positive[i]
	}
	for i := 0; i < len(negative); i++ {
		sum += negative[i]
	}
	fmt.Println(positive, negative, sum)
	if K > 0 {
		if K % 2 == 0 {
			return sum
		} else {
			sort.Ints(positive)
			sort.Ints(negative)
			if len(negative) == 0 {
				sum -= 2*positive[0]
			} else if len(positive) == 0 {
				sum -= 2*negative[0]
			} else {
				sum -= 2*minInt(positive[0], negative[0])
			}
		}
	}
	return sum
}

func main() {
	fmt.Println(largestSumAfterKNegations([]int{-2,5,0,2,-2}, 3))
}

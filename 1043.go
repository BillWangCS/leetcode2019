/*
Given an integer array A, you partition the array into (contiguous) subarrays of length at most K.
After partitioning, each subarray has their values changed to become the maximum value of that subarray.

Return the largest sum of the given array after partitioning.



Example 1:

Input: A = [1,15,7,9,2,5,10], K = 3
Output: 84
Explanation: A becomes [15,15,15,9,10,10,10]


Note:

1 <= K <= A.length <= 500
0 <= A[i] <= 10^6
 */

package main

import (
	"fmt"
)
//贪心，要更新的时候和前一个需要这个元素的方案比较，谁的和大录取谁
//p-1 p p+k-1 is the biggest
//p-k p p-k+1 is the biggest

func maxSumAfterPartitioning(A []int, K int) int {

	return res
}

func main() {
	fmt.Println(maxSumAfterPartitioning([]int{1,4,1,5,7,3,6,1,9,9,3}, 4))
}
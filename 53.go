/*
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
 */

package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	res := 0
	max := math.MinInt32
	sum := 0
	for i := 0; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
		if sum + nums[i] > res {
			res = sum + nums[i]
		}
		if sum + nums[i] > 0 {
			sum += nums[i]
		} else {
			sum = 0
		}
	}
	if max < 0 {
		res = max
	}
	return res
}

func main() {
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}

/*
Given an array nums of n integers and an integer target,
find three integers in nums such that the sum is closest to target.
Return the sum of the three integers. You may assume that each input would have exactly one solution.

Example:

Given array nums = [-1, 2, 1, -4], and target = 1.

The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
 */

package main

import (
	"fmt"
	"sort"
)

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	l := len(nums)
	res := 1 << 31 - 1
	resSum := 0
	for i := 0; i < l; i++ {
		j := 0
		k := l - 1
		t := target - nums[i]
		min := 1 << 31 - 1
		minSum := 0
		for j < k {
			if j == i {
				j++
				continue
			}
			if k == i {
				k--
				continue
			}
			sum := nums[j] + nums[k]
			if sum > t {
				k--
			} else if sum == t {
				return sum + nums[i]
			} else {
				j++
			}
			if abs(sum - t) < min {
				min = abs(sum - t)
				minSum = sum + nums[i]
			}
		}
		if min < res {
			res = min
			resSum = minSum
		}
	}
	return resSum
}

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 2))
}
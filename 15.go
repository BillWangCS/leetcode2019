/*
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0?
Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
 */

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	dict := make(map[string]bool)
	for i := 0; i < len(nums); i++ {
		j := 0
		k := len(nums) - 1
		for j < k {
			if j != i && k != i && nums[i] + nums[j] + nums[k] == 0 {
				tmp := []int{nums[i], nums[j], nums[k]}
				sort.Ints(tmp)
				s := ""
				for i := 0; i < 3; i++ {
					s += strconv.Itoa(tmp[i])
				}
				if !dict[s] {
					res = append(res, tmp)
					dict[s] = true
				}
				j++
			}
			if nums[i] + nums[j] + nums[k] < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return res
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
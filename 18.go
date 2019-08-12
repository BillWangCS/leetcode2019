/*
Given an array nums of n integers and an integer target, are there elements a, b, c,
and d in nums such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.

Note:

The solution set must not contain duplicate quadruplets.

Example:

Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.

A solution set is:
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
 */

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	l := len(nums)
	res := make([][]int, 0)
	dict := make(map[string]bool)
	for i := 0; i < l; i++ {
		for j := i+1; j < l; j++ {
			s := nums[i] + nums[j]
			m := 0
			n := l - 1
			for m < n {
				if m == i || m == j {
					m++
					continue
				}
				if n == i || n == j {
					n--
					continue
				}
				if nums[m] + nums[n] + s == target {
					tmp := []int{nums[i], nums[j], nums[m], nums[n]}
					sort.Ints(tmp)
					str := ""
					for x := 0; x < 4; x++ {
						str += strconv.Itoa(tmp[x])
					}
					if !dict[str] {
						res = append(res, tmp)
						dict[str] = true
					}
					m++
					n--
				} else if nums[m] + nums[n] + s < target {
					m++
				} else {
					n--
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}

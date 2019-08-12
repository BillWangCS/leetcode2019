/*
Given an array of integers and an integer k, find out whether there are two distinct indices i and j
in the array such that nums[i] = nums[j] and the absolute difference between i and j is at most k.

Example 1:

Input: nums = [1,2,3,1], k = 3
Output: true
Example 2:

Input: nums = [1,0,1,1], k = 1
Output: true
Example 3:

Input: nums = [1,2,3,1,2,3], k = 2
Output: false
 */

package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	minDist := 1 << 31 - 1
	dict := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		if _, ok := dict[nums[i]]; !ok {
			dict[nums[i]] = i
		} else {
			dist := i - dict[nums[i]]
			if dist < minDist {
				minDist = dist
			}
			dict[nums[i]] = i
		}
	}
	if minDist <= k {
		return true
	}
	return false
}

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1,2,3,1,2,3}, 2))
}
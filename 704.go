/*
Given a sorted (in ascending order) integer array nums of n elements and a target value,
write a function to search target in nums. If target exists, then return its index, otherwise return -1.


Example 1:

Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4

Example 2:

Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1


Note:

You may assume that all elements in nums are unique.
n will be in the range [1, 10000].
The value of each element in nums will be in the range [-9999, 9999].
 */

package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left + 1 < right {
		mid := (left + right)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}
	if nums[left] == target {
		return left
	}
	if nums[right] == target {
		return right
	}
	return -1
}

func main() {
	fmt.Println(search([]int{-1,0,3,5,9,12}, 100))
}
/*
Given a mountain sequence of n integers which increase firstly and then decrease, find the mountain top.

Example
Example 1:

Input: nums = [1, 2, 4, 8, 6, 3]
Output: 8
Example 2:

Input: nums = [10, 9, 8, 7],
Output: 10
 */

package main

import "fmt"

func mountainSequence (nums []int) int {
	// write your code here
	left, right := 0, len(nums) - 1
	if len(nums) == 1 {
		return nums[0]
	}
	if nums[left] > nums[left+1] {
		return nums[left]
	}
	if nums[right] > nums[right-1] {
		return nums[right]
	}
	for left + 1 < right {
		mid := (left + right) / 2
		fmt.Println(mid)
		if nums[mid+1] < nums[mid] && nums[mid] > nums[mid-1] {
			return nums[mid]
		}
		if nums[mid] > nums[mid-1] && nums[mid+1] > nums[mid] {
			left = mid
		} else if nums[mid] < nums[mid-1] && nums[mid+1] < nums[mid] {
			right = mid
		}
	}
	return -1
}

func main() {
	fmt.Println(mountainSequence([]int{1,2,3,4,5,10,4,3,2,1}))
}

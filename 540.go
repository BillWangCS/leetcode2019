/*
Given a sorted array consisting of only integers where every element appears twice except for one element which appears once.
Find this single element that appears only once.



Example 1:

Input: [1,1,2,2,3,4,4,8,8]
Output: 2
Example 2:

Input: [3,3,4,7,7,11,11]
Output: 10


Note: Your solution should run in O(log n) time and O(1) space.
 */

package main

import "fmt"

func singleNonDuplicate(nums []int) int {
	left := 0
	mid := 0
	right := len(nums) - 1
	for left + 1 < right {
		mid = (left + right) / 2
		fmt.Println(mid)
		if nums[mid] == nums[mid-1] {
			if mid & 1 == 0 {
				right = mid
			} else {
				left = mid
			}
		} else if nums[mid] == nums[mid+1] {
			 if mid & 1 == 0 {
			 	left = mid
			 } else {
			 	right = mid
			 }
		} else {
			return nums[mid]
		}
	}
	if left == 0 {
		if nums[right] == nums[right+1] {
			return nums[left]
		} else {
			return nums[right+1]
		}
	}
	if right == len(nums) - 1 {
		if nums[left] == nums[left-1] {
			return nums[right]
		} else {
			return nums[left-1]
		}
	}
	if nums[left] == nums[left-1] {
		return nums[right]
	}
	return nums[left]
}

func main() {
	fmt.Println(singleNonDuplicate([]int{1,2,2,3,3}))
}

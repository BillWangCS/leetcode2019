/*
Find the last position of a target number in a sorted array. Return -1 if target does not exist.

Example
Example 1:

Input: nums = [1,2,2,4,5,5], target = 2
Output: 2
Example 2:

Input: nums = [1,2,2,4,5,5], target = 6
Output: -1
 */

package main

import "fmt"

func lastPosition(nums []int, target int) int {
	// write your code here
	left, right := 0, len(nums) - 1
	pos := 0
	for left + 1 < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			pos = mid
			break
		}
		if nums[mid] > target {
			right = mid
		} else {
			left = mid
		}
	}
	if nums[left] == target {
		pos = left
	} else if nums[right] == target {
		pos = right
	}
	fmt.Println(pos)
	if nums[pos] != target {
		return -1
	}
	for i := pos; i < len(nums); i++ {
		if i == len(nums) - 1 && target == nums[i] {
			return i
		}
		if target != nums[i] {
			return i-1
		}
	}
	return -1
}

func main() {
	fmt.Println(lastPosition([]int{1,2,2,4,5,5}, 5))
}

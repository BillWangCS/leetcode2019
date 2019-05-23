/*
Given an array with n integers, your task is to check if it could become non-decreasing by modifying at most 1 element.

We define an array is non-decreasing if array[i] <= array[i + 1] holds for every i (1 <= i < n).

Example 1:
Input: [4,2,3]
Output: True
Explanation: You could modify the first 4 to 1 to get a non-decreasing array.
Example 2:
Input: [4,2,1]
Output: False
Explanation: You can't get a non-decreasing array by modify at most one element.
Note: The n belongs to [1, 10,000].
 */

package main

import "fmt"

func isOrderedArr(nums []int) bool {
	for i := 0; i+1 < len(nums); i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}

func checkPossibility(nums []int) bool {
	l := len(nums)
	if l == 0 || l == 1 {
		return true
	}
	if isOrderedArr(nums) {
		return true
	}
	for i := 0; i + 1 < l; i++ {
		if nums[i] > nums[i+1] {
			tmp := make([]int, 0)
			tmp2 := make([]int, 0)
			tmp = append(tmp, nums[:i]...)
			tmp = append(tmp, nums[i+1:]...)
			tmp2 = append(tmp2, nums[:i+1]...)
			if i + 2 < l {
				tmp2 = append(tmp2, nums[i+2:]...)
			}
			fmt.Println(tmp, tmp2)
			if isOrderedArr(tmp) || isOrderedArr(tmp2) && len(tmp2) > 0 {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(checkPossibility([]int{1,2,4,5,3}))
}
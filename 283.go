/*
Given an array nums, write a function to move all 0's
to the end of it while maintaining the relative order of the non-zero elements.

Example:

Input: [0,1,0,3,12]
Output: [1,3,12,0,0]
Note:

You must do this in-place without making a copy of the array.
Minimize the total number of operations.
 */

package main

func moveZeroes(nums []int)  {
	cur := 0
	l := len(nums)
	for i := 0; i < l; i++ {
		if nums[i] != 0{
			nums[cur] = nums[i]
			cur++
		}
	}
	for i := cur; i < l; i++ {
		nums[i] = 0
	}
}

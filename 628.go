/*
Given an integer array, find three numbers whose product is maximum and output the maximum product.

Example 1:

Input: [1,2,3]
Output: 6


Example 2:

Input: [1,2,3,4]
Output: 24


Note:

The length of the given array will be in range [3,104] and all elements are in the range [-1000, 1000].
Multiplication of any three numbers in the input won't exceed the range of 32-bit signed integer.
 */

package main

import (
	"fmt"
	"sort"
)

func maxInteger(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	return maxInteger(nums[0]*nums[1]*nums[len(nums)-1], nums[len(nums)-1]*nums[len(nums)-2]*nums[len(nums)-3])
}

func main() {
	fmt.Println(maximumProduct([]int{-4,-3,-2,-1,60}))
}
/*
Given an integer array nums, find the contiguous subarray within an array (containing at least one number)
which has the largest product.

Example 1:

Input: [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.
Example 2:

Input: [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
 */

package main

import "fmt"

func getProduct(nums []int) int {
	cnt := 0
	p := 1
	firstFlag := true
	first := -1
	last := -1
	for i := 0; i < len(nums); i++ {
		p *= nums[i]
		if nums[i] < 0 {
			cnt++
			if firstFlag {
				first = i
				firstFlag = false
			}
			last = i
		}
	}
	if cnt % 2 == 0 {
		return p
	} else {
		p1 := 1
		if cnt == 1 && len(nums) == 1 {
			return nums[first]
		}
		for i := 0; i < last; i++ {
			p1 *= nums[i]
		}
		p2 := 1
		for i := first; i < len(nums); i++ {
			p2 *= nums[i]
		}
		p2 = p2 / nums[first]
		if p1 >= p2 {
			return p1
		} else {
			return p2
		}
	}
}

func maxProduct(nums []int) int {
	arrs := make([][]int, 0)
	res := -1<<31
	tmp := make([]int, 0)
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			tmp = append(tmp, nums[i])
		}
		if nums[i] == 0 {
			cnt++
			if len(tmp) > 0 {
				arrs = append(arrs, tmp)
			}
			tmp = make([]int, 0)
		}
	}
	if len(tmp) > 0 {
		arrs = append(arrs, tmp)
	}
	fmt.Println(arrs)
	for i := 0; i < len(arrs); i++ {
		p := getProduct(arrs[i])
		fmt.Println(p, arrs[i])
		if p > res {
			res = p
		}
	}
	if cnt > 0 && res < 0 {
		res = 0
	}
	return res
}

func main() {
	fmt.Println(maxProduct([]int{-1,1,2,1}))
}
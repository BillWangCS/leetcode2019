/*
Given an array nums of n integers where n > 1,
return an array output such that output[i] is equal to the product of all the elements of nums except nums[i].

Example:

Input:  [1,2,3,4]
Output: [24,12,8,6]
Note: Please solve it without division and in O(n).

Follow up:
Could you solve it with constant space complexity?
(The output array does not count as extra space for the purpose of space complexity analysis.)
 */

package main

import "fmt"

func productExceptSelf(nums []int) []int {
	res := make([]int, 0)
	if len(nums) == 0 {
		return res
	}
	product := 1
	cnt := 0
	for _, v := range nums {
		if v == 0{
			cnt++
			continue
		}
		product *= v
	}
	if cnt > 1 {
		product = 0
	}
	for _, v := range nums {
		if v == 0 {
			res = append(res, product)
		} else {
			if cnt > 0 {
				res = append(res, 0)
			} else {
				res = append(res, product/v)
			}
		}
	}
	return res
}

func main() {
	fmt.Println(productExceptSelf([]int{1,0}))
}
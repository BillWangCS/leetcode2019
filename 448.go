/*
Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

Find all the elements of [1, n] inclusive that do not appear in this array.

Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

Example:

Input:
[4,3,2,7,8,2,3,1]

Output:
[5,6]
 */

package main

import "fmt"


func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

func findDisappearedNumbers(nums []int) []int {
	res := make([]int, 0)
	for _, v := range nums {
		v = abs(v) - 1
		if nums[v] > 0 {
			nums[v] = -nums[v]
		}
	}
	for i, v := range nums {
		if v > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func main() {
	fmt.Println(findDisappearedNumbers([]int{4,3,2,7,8,2,3,1}))
}
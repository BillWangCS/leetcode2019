/*
Given an array of integers, 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

Find all the elements that appear twice in this array.

Could you do it without extra space and in O(n) runtime?

Example:
Input:
[4,3,2,7,8,2,3,1]

Output:
[2,3]
 */

package main

import "fmt"

func findDuplicates(nums []int) []int {
	res := make([]int, 0)
	l := len(nums)
	for i := 0; i < l; i++ {
		index := (nums[i]-1) % l
		nums[index] += l
	}
	for i := 0; i < l; i++ {
		if nums[i] > 2*l {
			res = append(res, i + 1)
		}
	}
	fmt.Println(nums)
	return res
}

func main() {
	fmt.Println(findDuplicates([]int{4,3,2,7,8,2,3,1}))
}
/*
Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.

Example 1:

Input: [3,0,1]
Output: 2
Example 2:

Input: [9,6,4,2,3,5,7,0,1]
Output: 8
 */

package main

import "fmt"

func missingNumber(nums []int) int {
	l := len(nums)
	s := (l+1) * l / 2
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return s - sum
}

func main() {
	fmt.Println(missingNumber([]int{9,6,4,2,3,5,7,0,1}))
}
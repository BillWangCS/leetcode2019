/*
Given an array of integers nums, sort the array in ascending order.



Example 1:

Input: [5,2,3,1]
Output: [1,2,3,5]
Example 2:

Input: [5,1,1,2,0,0]
Output: [0,0,1,1,2,5]


Note:

1 <= A.length <= 10000
-50000 <= A[i] <= 50000
 */

package main

import "fmt"

func sortHelper(nums *[]int, start, end int) {
	fmt.Println(start, end, *nums)
	if start >= end {
		return
	}
	left := start
	right := end
	pivot := (*nums)[(left + right)/2]
	for left <= right {
		for left <= right && (*nums)[left] < pivot {
			left++
		}
		for left <= right && (*nums)[right] > pivot {
			right--
		}
		if left <= right {
			(*nums)[left], (*nums)[right] = (*nums)[right], (*nums)[left]
			left++
			right--
		}
	}
	sortHelper(nums, left, end)
	sortHelper(nums, start, right)
}

func sortArray(nums []int) []int {
	sortHelper(&nums, 0, len(nums)-1)
	return nums
}

func main() {
	fmt.Println(sortArray([]int{5,1,1,2,0,0}))
}
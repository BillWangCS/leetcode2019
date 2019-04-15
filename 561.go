/*
Given an array of 2n integers, your task is to group these integers into n pairs of integer,
say (a1, b1), (a2, b2), ..., (an, bn) which makes sum of min(ai, bi) for all i from 1 to n as large as possible.

Example 1:
Input: [1,4,3,2]

Output: 4
Explanation: n is 2, and the maximum sum of pairs is 4 = min(1, 2) + min(3, 4).
Note:
n is a positive integer, which is in the range of [1, 10000].
All the integers in the array will be in the range of [-10000, 10000].
 */

package main

import "fmt"

//import "sort"

func quickSort2(A []int, start, end int) {
	if start >= end {
		return
	}
	left, right := start, end
	pivot := A[(start + end) / 2]
	for left <= right {
		for left <= right && A[left] < pivot {
			left++
		}
		for left <= right && A[right] > pivot {
			right--
		}
		if left <= right {
			A[left], A[right] = A[right], A[left]
			left++
			right--
		}
	}
	quickSort2(A, start, right)
	quickSort2(A, left, end)
	return
}

func arrayPairSum(nums []int) int {
	l := len(nums)
	quickSort2(nums, 0, l-1)
	fmt.Println(nums)
	res := 0
	for i := 0; i < l; i += 2 {
		res += nums[i]
	}
	if l % 2 == 1 {
		res += nums[l-1]
	}
	return res
}

func main() {
	fmt.Println(arrayPairSum([]int{-470, 66, -4835, -5623}))
}
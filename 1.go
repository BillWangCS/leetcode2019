/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
 */

package main

import (
	"fmt"
)

type elem struct {
	val int
	index int
}

func quickSort(A *[]elem, start, end int) {
	if start >= end {
		return
	}
	left, right := start, end
	pivot := (*A)[(start+end)/2].val
	for left <= right {
		for left <= right && (*A)[left].val < pivot {
			left++
		}
		for left <= right && (*A)[right].val > pivot {
			right--
		}
		if left <= right {
			(*A)[left], (*A)[right] = (*A)[right], (*A)[left]
			left++
			right--
		}
	}
	quickSort(A, start, right)
	quickSort(A, left, end)
}

func twoSum(nums []int, target int) []int {
	res := make([]int, 0)
	arr := make([]elem, 0)
	for i := 0; i < len(nums); i++ {
		arr = append(arr, elem{nums[i], i})
	}
	i := 0
	j := len(nums) - 1
	quickSort(&arr, i, j)
	fmt.Println(arr)
	for i < j {
		if arr[i].val + arr[j].val == target {
			res = append(res, arr[i].index)
			res = append(res, arr[j].index)
			break
		}
		if arr[i].val + arr[j].val < target {
			i++
		}
		if arr[i].val + arr[j].val > target {
			j--
		}
	}
	return res
}

func main() {
	fmt.Println(twoSum([]int{-10,-1,-18,-19}, -19))
}

/*
[-10,-1,-18,-19]
-19
 */
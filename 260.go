/*
Given an array of numbers nums, in which exactly two elements appear only once
and all the other elements appear exactly twice. Find the two elements that appear only once.

Example:

Input:  [1,2,1,3,2,5]
Output: [3,5]
Note:

The order of the result is not important. So in the above example, [5, 3] is also correct.
Your algorithm should run in linear runtime complexity. Could you implement it using only constant space complexity?
 */

package main

import "fmt"

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

func singleNumber(nums []int) []int {
	l := len(nums)
	res := make([]int, 0)
	if l == 0 {
		return res
	}
	s := 0
	for i := 0; i < l; i++ {
		s ^= nums[i]
	}
	fmt.Println(s)
	var cnt uint
	for abs(s) > 0 {
		if s & 1 == 1 {
			break
		}
		s = s >> 1
		cnt++
	}
	fmt.Println(cnt)
	x := 0
	y := 0
	for i := 0; i < l; i++ {
		if (nums[i] >> cnt) & 1 == 0 {
			x ^= nums[i]
		} else {
			y ^= nums[i]
		}
	}
	res = append(res, x)
	res = append(res, y)
	return res
}

func main() {
	fmt.Println(singleNumber([]int{43772400,1674008457,1779561093,744132272,1674008457,448610617,1779561093,124075538,-1034600064,49040018,612881857,390719949,-359290212,-812493625,124732,-1361696369,49040018,-145417756,-812493625,2078552599,1568689850,865876872,865876872,-1471385435,1816352571,1793963758,2078552599,-1034600064,1475115274,-119634980,124732,661111294,-1813882010,1568689850,448610617,1347212898,-1293494866,612881857,661111294,-1361696369,1816352571,-1813882010,-359290212,1475115274,1793963758,1347212898,43772400,-1471385435,124075538,-1293494866,-119634980,390719949}))
}
/*
Given two arrays, write a function to compute their intersection.

Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [4,9]
Note:

Each element in the result should appear as many times as it shows in both arrays.
The result can be in any order.
Follow up:

What if the given array is already sorted? How would you optimize your algorithm?
What if nums1's size is small compared to nums2's size? Which algorithm is better?
What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?
 */

package main

import "fmt"

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func intersect(nums1 []int, nums2 []int) []int {
	dict1 := make(map[int]int)
	dict2 := make(map[int]int)
	res := make([]int, 0)
	for _, v := range nums1 {
		dict1[v]++
	}
	for _, v := range nums2 {
		dict2[v]++
	}
	for k, v := range dict1 {
		if _, ok := dict2[k]; ok {
			n := minInt(dict2[k],v)
			for i := 0; i < n; i++ {
				res = append(res, k)
			}
		}
	}
	return res
}

func main() {
	fmt.Println(intersect([]int{4,9,5}, []int{9,4,9,8,4}))
}
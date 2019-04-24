/*
Given two arrays, write a function to compute their intersection.

Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]
Note:

Each element in the result must be unique.
The result can be in any order.
 */

package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	dictNums := make(map[int]int)
	for _, num := range nums1 {
		dictNums[num]++
	}
	for _, v := range nums2 {
		if dictNums[v] > 0 {
			res = append(res, v)
			/*follow up
Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2, 2]
 */
			//dictNums[v]--
			dictNums[v] = 0

		}
	}
	return res
}

func main() {
	fmt.Println(intersection([]int{4, 5, 9, 4}, []int{9, 4, 9, 8, 4}))
}

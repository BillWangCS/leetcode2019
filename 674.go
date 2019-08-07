/*
Given an unsorted array of integers, find the length of longest continuous increasing subsequence (subarray).

Example 1:
Input: [1,3,5,4,7]
Output: 3
Explanation: The longest continuous increasing subsequence is [1,3,5], its length is 3.
Even though [1,3,5,7] is also an increasing subsequence, it's not a continuous one where 5 and 7 are separated by 4.
Example 2:
Input: [2,2,2,2,2]
Output: 1
Explanation: The longest continuous increasing subsequence is [2], its length is 1.
Note: Length of the array will not exceed 10,000.
 */

package main

import "fmt"

func findLengthOfLCIS(nums []int) int {
	max := 0
	l := 1
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > l {
			cnt++
		} else {
			cnt = 1
		}
		l = nums[i]
		if cnt > max {
			max = cnt
		}
	}
	return max
}

func main() {
	fmt.Println(findLengthOfLCIS([]int{2,2,2,2,2,2,2}))
}
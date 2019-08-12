/*
Given an array, rotate the array to the right by k steps, where k is non-negative.

Example 1:

Input: [1,2,3,4,5,6,7] and k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]
Example 2:

Input: [-1,-100,3,99] and k = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]
Note:

Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.
Could you do it in-place with O(1) extra space?
 */

package main

import "fmt"

func rotate(nums []int, k int) {
	l := len(nums)
	x := k % l
	if x == 0 {
		return
	}
	i := 0
	tmp := make([]int, 0)
	for i < l-x {
		tmp = append(tmp, nums[i])
		i++
	}
	j := 0
	for i := l - x; i < l; i ++ {
		nums[j] = nums[i]
		j++
	}
	i = 0
	for j < l {
		nums[j] = tmp[i]
		j++
		i++
	}
	fmt.Println(nums)
	return
}

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}

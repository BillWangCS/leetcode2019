/*
Given a non-empty array of integers, every element appears three times except for one, which appears exactly once. Find that single one.

Note:

Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

Example 1:

Input: [2,2,3,2]
Output: 3
Example 2:

Input: [0,1,0,1,0,1,99]
Output: 99
 */

package main

import "fmt"

func singleNumber(nums []int) int {
	res := 0
	var i uint
	for i = 0; i <= 64; i++ {
		cnt := 0
		for j := 0; j < len(nums); j++ {
			if 1<<i & nums[j] != 0 {
				cnt++
			}
		}
		if cnt % 3 != 0 {
			fmt.Println(1<<i)
			res += 1<<i
		}
	}
	return res
}

func main() {
	fmt.Println(singleNumber([]int{0,1,0,1,0,1,-99}))
}
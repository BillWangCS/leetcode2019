/*
We define a harmounious array as an array where the difference between its maximum value and its minimum value is exactly 1.

Now, given an integer array, you need to find the length of its longest harmonious subsequence among all its possible subsequences.

Example 1:

Input: [1,3,2,2,5,2,3,7]
Output: 5
Explanation: The longest harmonious subsequence is [3,2,2,2,3].


Note: The length of the input array will not exceed 20,000.
 */

package main

import "fmt"

func findLHS(nums []int) int {
	res := 0
	dict := make(map[int]int)
	l := len(nums)
	for i := 0; i< l; i++ {
		dict[nums[i]]++
	}
	for i := 0; i < l; i++ {
		if _, ok := dict[nums[i] + 1]; ok {
			if res < dict[nums[i]] + dict[nums[i] + 1] {
				res = dict[nums[i]] + dict[nums[i] + 1]
			}
		}
	}
	return res
}

func main() {
	fmt.Println(findLHS([]int{1,3,2,2,5,2,3,7,2}))
}
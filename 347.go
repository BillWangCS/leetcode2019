/*
Given a non-empty array of integers, return the k most frequent elements.

Example 1:

Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
Example 2:

Input: nums = [1], k = 1
Output: [1]
Note:

You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
 */

package main

import (
	"fmt"
	"sort"
)

func quickK(A *[]freqInfo, start, end int, K int) {
	if start >= end {
		return
	}
	left, right := start, end
	pivot := (*A)[(end - start) / 2 + start]
	for left <= right {
		for left <= right && (*A)[left].freq > pivot.freq {
			left++
		}
		for left <= right && (*A)[right].freq < pivot.freq {
			right--
		}
		if left <= right {
			(*A)[left], (*A)[right] = (*A)[right], (*A)[left]
			left++
			right--
		}
	}
	if K >= left {
		quickK(A, left, end, K)
	} else if K <= right {
		quickK(A, start, right, K)
	}
	return
}

type freqInfo struct {
	val int
	freq int
}

func topKFrequent(nums []int, k int) []int {
	res := make([]int, 0)
	dict := make(map[int]int)
	freqInfoArr := make([]freqInfo, 0)
	for i := 0; i < len(nums); i++ {
		dict[nums[i]]++
	}
	for k, v := range dict {
		tmp := freqInfo{k, v}
		freqInfoArr = append(freqInfoArr, tmp)
	}
	fmt.Println(freqInfoArr)
	quickK(&freqInfoArr, 0, len(freqInfoArr)-1, k)
	fmt.Println(freqInfoArr)
	for i := 0; i < k; i++ {
		res = append(res, freqInfoArr[i].val)
	}
	return res
}

func main() {
	fmt.Println(topKFrequent([]int{1,1,1,2,2,3}, 2))
}
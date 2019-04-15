/*
Given a set of distinct integers, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.

Example:

Input: nums = [1,2,3]
Output:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
 */

package main

import (
	"fmt"
	"sort"
)

//       BFS commented for subsets

//func inArray(A []int, n int) bool {
//	for _, v := range A {
//		if v == n {
//			return true
//		}
//	}
//	return false
//}
//

//func subsets(nums []int) [][]int {
//	res := make([][]int, 0)
//	path := make([]int, 0)
//	queue := make([][]int, 0)
//	start := make([]int, 0)
//	queue = append(queue, start)
//	sort.Ints(nums)
//	for len(queue) > 0 {
//		l := len(queue)
//		for i := 0; i < l; i++ {
//			path = queue[0]
//			x := make([]int, 0)
//			x = append(x, path...)
//			res = append(res, x)
//			for j := 0; j < len(nums); j++ {
//				if !inArray(path, nums[j]) && (len(path) > 0 && nums[j] > path[len(path)-1]) || len(path) == 0 {
//					path = append(path, nums[j])
//					x := make([]int, 0)
//					x = append(x, path...)
//					queue = append(queue, x)
//					path = path[:len(path)-1]
//				}
//			}
//			queue = queue[1:]
//		}
//	}
//	return res
//}

func subsets2(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	sort.Ints(nums)
	dfs(nums, 0, path, &res)
	return res
}

func dfs(nums []int, index int, path []int, res *[][]int) {
	x := make([]int, 0)
	x = append(x, path...)
	*res = append(*res, x)
	for i := index; i < len(nums); i++ {
		//for subsets2 dedup
		if i > index && nums[i-1] == nums[i] {
			continue
		}
		path = append(path, nums[i])
		dfs(nums, i+1, path, res)
		path = path[:len(path)-1]
	}
}

func main() {
	fmt.Println(subsets([]int{1, 2, 2}))
}

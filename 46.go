/*
Given a collection of distinct integers, return all possible permutations.

Example:

Input: [1,2,3]
Output:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
 */

package main

import "fmt"

func isNotIn(results [][]int, path []int) bool {
	for _, v := range results {
		flag := false
		for i := 0; i < len(path); i++ {
			if path[i] != v[i] {
				flag = true
				break
			}
		}
		if !flag {
			return false
		}
	}
	return true
}

func dfs_46(nums []int, path *[]int, visited *map[int]bool, results *[][]int) {
	if len(*path) == len(nums) {
		if isNotIn(*results, *path) {
			tmp := make([]int, 0)
			tmp = append(tmp, *path...)
			*results = append(*results, tmp)
		}
	}
	for i := 0; i< len(nums); i++ {
		if !(*visited)[i] {
			*path = append(*path, nums[i])
			(*visited)[i] = true
			dfs_46(nums, path, visited, results)
			*path = (*path)[:len(*path)-1]
			(*visited)[i] = false
		}
	}
}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	visited := make(map[int]bool)
	dfs_46(nums, &path, &visited, &res)
	return res
}

func main() {
	fmt.Println(permute([]int{1,2}))
}
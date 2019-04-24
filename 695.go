/*
Given a non-empty 2D array grid of 0's and 1's, an island is a group of 1's (representing land)
connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.

Find the maximum area of an island in the given 2D array. (If there is no island, the maximum area is 0.)

Example 1:

[[0,0,1,0,0,0,0,1,0,0,0,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,1,1,0,1,0,0,0,0,0,0,0,0],
 [0,1,0,0,1,1,0,0,1,0,1,0,0],
 [0,1,0,0,1,1,0,0,1,1,1,0,0],
 [0,0,0,0,0,0,0,0,0,0,1,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,0,0,0,0,0,0,1,1,0,0,0,0]]
Given the above grid, return 6. Note the answer is not 11, because the island must be connected 4-directionally.
Example 2:

[[0,0,0,0,0,0,0,0]]
Given the above grid, return 0.
Note: The length of each dimension in the given grid does not exceed 50.
 */

package main

import "fmt"

func getRoot(father []int, u int) int {
	if father[u] == u {
		return u
	}
	return getRoot(father, father[u])
}

func merge(father *[]int, u, v int) {
	fu := getRoot(*father, u)
	fv := getRoot(*father, v)
	if fu != fv {
		(*father)[fu] = fv
	}
}

func maxAreaOfIsland(grid [][]int) int {
	father := make([]int, 0)
	cnt := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				father = append(father, cnt)
			} else {
				father = append(father, -1)
			}
			cnt++
		}
	}
	fmt.Println(father)
	cnt = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				if i + 1 < len(grid) && grid[i+1][j] == 1 {
					merge(&father, cnt, cnt + len(grid[0]))
				}

				if j + 1 < len(grid[0]) && grid[i][j+1] == 1 {
					merge(&father, cnt, cnt + 1)
				}
			}
			cnt++
		}
	}
	visited := make(map[int]int)
	for i := 0; i < len(father); i++ {
		if father[i] != -1 {
			visited[getRoot(father, father[i])]++
		}
	}
	res := 0
	for _, v := range visited {
		if v > res {
			res = v
		}
	}
	return res
}

func main() {
	fmt.Println(maxAreaOfIsland([][]int{[]int{0,0,1,0,0,0,0,1,0,0,0,0,0}, []int{0,0,0,0,0,0,0,1,1,1,0,0,0}, []int{0,1,1,0,1,0,0,0,0,0,0,0,0}, []int{0,1,0,0,1,1,0,0,1,0,1,0,0}, []int{0,1,0,0,1,1,0,0,1,1,1,0,0}, []int{0,0,0,0,0,0,0,0,0,0,1,0,0}, []int{0,0,0,0,0,0,0,1,1,1,0,0,0}, []int{0,0,0,0,0,0,0,1,1,0,0,0,0}}))
}

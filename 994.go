/*
In a given grid, each cell can have one of three values:

the value 0 representing an empty cell;
the value 1 representing a fresh orange;
the value 2 representing a rotten orange.
Every minute, any fresh orange that is adjacent (4-directionally) to a rotten orange becomes rotten.

Return the minimum number of minutes that must elapse until no cell has a fresh orange.  If this is impossible, return -1 instead.



Example 1:



Input: [[2,1,1],[1,1,0],[0,1,1]]
Output: 4
Example 2:

Input: [[2,1,1],[0,1,1],[1,0,1]]
Output: -1
Explanation:  The orange in the bottom left corner (row 2, column 0) is never rotten, because rotting only happens 4-directionally.
Example 3:

Input: [[0,2]]
Output: 0
Explanation:  Since there are already no fresh oranges at minute 0, the answer is just 0.


Note:

1 <= grid.length <= 10
1 <= grid[0].length <= 10
grid[i][j] is only 0, 1, or 2.
 */

package main

import "fmt"

func dfs_994(grid [][]int, i, j, step int, minStep *int, visited *[][]bool) {
	if grid[i][j] == 0 || (*visited)[i][j] {
		return
	}
	if grid[i][j] == 2 {
		if step < *minStep {
			*minStep = step
		}
		return
	}
	(*visited)[i][j] = true
	if i-1 >= 0 {
		dfs_994(grid, i-1, j, step+1, minStep, visited)
	}
	if i+1 < len(grid) {
		dfs_994(grid, i+1, j, step+1, minStep, visited)
	}
	if j-1 >= 0 {
		dfs_994(grid, i, j-1, step+1, minStep, visited)
	}
	if j+1 < len(grid[0]) {
		dfs_994(grid, i, j+1, step+1, minStep, visited)
	}
	(*visited)[i][j] = false
}

func orangesRotting(grid [][]int) int {
	res := 0
	if len(grid) == 0 {
		return -1
	}
	visited := make([][]bool, 0)
	for i := 0; i < len(grid); i++ {
		tmp := make([]bool, 0)
		for j :=0; j < len(grid[0]); j++ {
			tmp = append(tmp, false)
		}
		visited = append(visited, tmp)
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				minStep := 1<<31
				dfs_994(grid, i, j, 0, &minStep, &visited)
				if minStep > res {
					res = minStep
				}
			}
		}
	}
	if res == 1<<31 {
		return -1
	}
	return res
}

func main() {
	fmt.Println(orangesRotting([][]int{[]int{0,2}}))
	//fmt.Println(dfs_994([][]int{[]int{0,2}, []int{0,1}, []int{0,1}, []int{1,1}, []int{1,1}, []int{1,1}}, 5, 0))
}
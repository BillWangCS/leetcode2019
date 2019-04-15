/*
On a N * N grid, we place some 1 * 1 * 1 cubes.

Each value v = grid[i][j] represents a tower of v cubes placed on top of grid cell (i, j).

Return the total surface area of the resulting shapes.



Example 1:

Input: [[2]]
Output: 10
Example 2:

Input: [[1,2],[3,4]]
Output: 34
Example 3:

Input: [[1,0],[0,2]]
Output: 16
Example 4:

Input: [[1,1,1],[1,0,1],[1,1,1]]
Output: 32
Example 5:

Input: [[2,2,2],[2,1,2],[2,2,2]]
Output: 46


Note:

1 <= N <= 50
0 <= grid[i][j] <= 50
 */

package main

import "fmt"

func surfaceArea(grid [][]int) int {
	res := 0
	r := len(grid)
	c := len(grid[0])
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == 0 {
				continue
			}
			up := 0
			down := 0
			left := 0
			right := 0
			if i == 0 {
				up = grid[i][j]
			} else {
				if grid[i][j] > grid[i-1][j] {
					up = grid[i][j] - grid[i-1][j]
				}
			}
			if i == r - 1 {
				down = grid[i][j]
			} else {
				if grid[i][j] > grid[i+1][j] {
					down = grid[i][j] - grid[i+1][j]
				}
			}
			if j == 0 {
				left = grid[i][j]
			} else {
				if grid[i][j] > grid[i][j-1] {
					left = grid[i][j] - grid[i][j-1]
				}
			}
			if j == c - 1 {
				right = grid[i][j]
			} else {
				if grid[i][j] > grid[i][j+1] {
					right = grid[i][j] - grid[i][j+1]
				}
			}
			res = 2 + res + up + down + left + right
		}
	}
	return res
}

func main() {
	fmt.Println(surfaceArea([][]int{[]int{1,1,1}, []int{1,0,1}, []int{1,1,1}}))
}
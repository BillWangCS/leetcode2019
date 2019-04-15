/*
You are given a map in form of a two-dimensional integer grid where 1 represents land and 0 represents water.

Grid cells are connected horizontally/vertically (not diagonally). The grid is completely surrounded by water,
and there is exactly one island (i.e., one or more connected land cells).

The island doesn't have "lakes" (water inside that isn't connected to the water around the island).
One cell is a square with side length 1. The grid is rectangular, width and height don't exceed 100.
Determine the perimeter of the island.



Example:

Input:
[[0,1,0,0],
 [1,1,1,0],
 [0,1,0,0],
 [1,1,0,0]]

Output: 16

Explanation: The perimeter is the 16 yellow stripes in the image below:

 */

package main

import "fmt"

func islandPerimeter(grid [][]int) int {
	res := 0
	r := len(grid)
	if r == 0 {
		return 0
	}
	c := len(grid[0])
	if c == 0 {
		return 0
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == 0 {
				continue
			}
			if i - 1 < 0 {
				res++
			}
			if j - 1 < 0 {
				res++
			}
			if i - 1 >= 0 && grid[i-1][j] == 0 {
				res++
			}
			if j - 1 >= 0 && grid[i][j-1] == 0 {
				res++
			}
			if i + 1 == r {
				res++
			}
			if j + 1 == c {
				res++
			}
			if j + 1 < c && grid[i][j+1] == 0 {
				res++
			}
			if i + 1 < r && grid[i+1][j] == 0 {
				res++
			}
		}
	}
	return res
}

func main() {
	fmt.Println(islandPerimeter([][]int{[]int{0,1,0,0}, []int{1,1,1,0}, []int{0,1,0,0}, []int{1,1,0,0}}))
}
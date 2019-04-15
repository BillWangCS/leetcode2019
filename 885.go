/*
On a 2 dimensional grid with R rows and C columns, we start at (r0, c0) facing east.

Here, the north-west corner of the grid is at the first row and column,
and the south-east corner of the grid is at the last row and column.

Now, we walk in a clockwise spiral shape to visit every position in this grid.

Whenever we would move outside the boundary of the grid,
we continue our walk outside the grid (but may return to the grid boundary later.)

Eventually, we reach all R * C spaces of the grid.

Return a list of coordinates representing the positions of the grid in the order they were visited.



Example 1:

Input: R = 1, C = 4, r0 = 0, c0 = 0
Output: [[0,0],[0,1],[0,2],[0,3]]




Example 2:

Input: R = 5, C = 6, r0 = 1, c0 = 4
Output: [[1,4],[1,5],[2,5],[2,4],[2,3],[1,3],[0,3],[0,4],[0,5],
[3,5],[3,4],[3,3],[3,2],[2,2],[1,2],[0,2],[4,5],[4,4],[4,3],
[4,2],[4,1],[3,1],[2,1],[1,1],[0,1],[4,0],[3,0],[2,0],[1,0],[0,0]]




Note:

1 <= R <= 100
1 <= C <= 100
0 <= r0 < R
0 <= c0 < C

 */

package main

import "fmt"

func maxInteger(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
	direct := 0
	res := make([][]int, 0)
	directMap := make(map[int]int)
	directMap[0] = 1
	directMap[1] = 1
	directMap[2] = 2
	directMap[3] = 2
	x := 0
	y := 0
	step := directMap[direct]
	for step <= maxInteger(R, C)*2 {
		if direct == 0 {
			step = directMap[direct]
			for n := 0; n < step; n++ {
				if (y+c0 >= 0 && y+c0 <= C-1) && (x+r0 >= 0 && x+r0 <= R-1) {
					pos := []int{x+r0, y+c0}
					res = append(res, pos)
				}
				y++
			}
			directMap[direct] += 2
			direct = (direct + 1) % 4
		}
		if direct == 1 {
			step = directMap[direct]
			for n := 0; n < step; n++ {
				if (y+c0 >= 0 && y+c0 <= C-1) && (x+r0 >= 0 && x+r0 <= R-1) {
					pos := []int{x+r0, y+c0}
					res = append(res, pos)
				}
				x++
			}
			directMap[direct] += 2
			direct = (direct + 1) % 4
		}
		if direct == 2 {
			step = directMap[direct]
			for n := 0; n < step; n++ {
				if (y+c0 >= 0 && y+c0 <= C-1) && (x+r0 >= 0 && x+r0 <= R-1) {
					pos := []int{x+r0, y+c0}
					res = append(res, pos)
				}
				y--
			}
			directMap[direct] += 2
			direct = (direct + 1) % 4
		}
		if direct == 3 {
			step = directMap[direct]
			for n := 0; n < step; n++ {
				if (y+c0 >= 0 && y+c0 <= C-1) && (x+r0 >= 0 && x+r0 <= R-1) {
					pos := []int{x+r0, y+c0}
					res = append(res, pos)
				}
				x--
			}
			directMap[direct] += 2
			direct = (direct + 1) % 4
		}
	}
	return res
}

func main() {
	fmt.Println(spiralMatrixIII(10, 3, 0, 1))
}
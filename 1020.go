/*
Given a 2D array A, each cell is 0 (representing sea) or 1 (representing land)

A move consists of walking from one land square 4-directionally to another land square, or off the boundary of the grid.

Return the number of land squares in the grid for which we cannot walk off the boundary of the grid in any number of moves.



Example 1:

Input: [[0,0,0,0],
		[1,0,1,0],
		[0,1,1,0],
		[0,0,0,0]]
Output: 3
Explanation:
There are three 1s that are enclosed by 0s, and one 1 that isn't enclosed because its on the boundary.
Example 2:

Input: [[0,1,1,0],[0,0,1,0],[0,0,1,0],[0,0,0,0]]
Output: 0
Explanation:
All 1s are either on the boundary or can reach the boundary.


Note:

1 <= A.length <= 500
1 <= A[i].length <= 500
0 <= A[i][j] <= 1
All rows have the same size.
 */

package main

import "fmt"

func dfs_1020(A *[][]int, i, j int) {
	if i < 0 || i >= len(*A) || j < 0 || j >= len((*A)[0]) || (*A)[i][j] == -2 || (*A)[i][j] == 0 {
		return
	}
	(*A)[i][j] = -2
	dfs_1020(A, i-1, j)
	dfs_1020(A, i+1, j)
	dfs_1020(A, i, j-1)
	dfs_1020(A, i, j+1)
}

func numEnclaves(A [][]int) int {
	if len(A) < 3 {
		return 0
	}
	if len(A[0]) < 3 {
		return 0
	}
	cntUp := 0
	cntDown := len(A) - 1
	cntLeft := 0
	cntRight := len(A[0]) - 1
	for i := cntUp; i <= cntDown; i++ {
		if A[i][cntLeft] == 1 {
			dfs_1020(&A, i, cntLeft)
		}
	}
	cntLeft++
	for i := cntUp; i <= cntDown; i++ {
		if A[i][cntRight] == 1 {
			dfs_1020(&A, i, cntRight)
		}
	}
	cntRight--
	for i := cntLeft; i <= cntRight; i++ {
		if A[cntUp][i] == 1 {
			dfs_1020(&A, cntUp, i)
		}
	}
	cntUp++
	for i := cntLeft; i <= cntRight; i++ {
		if A[cntDown][i] == 1 {
			dfs_1020(&A, cntDown, i)
		}
	}
	cntDown--
	res := 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {
			if A[i][j] == 1 {
				res++
			}
		}
	}
	return res
}

func main() {
	fmt.Println(numEnclaves([][]int{[]int{0,0,0,1,1,1,0,1,0,0},
									[]int{1,1,0,0,0,1,0,1,1,1},
									[]int{0,0,0,1,1,1,0,1,0,0},
									[]int{0,1,1,0,0,0,1,0,1,0},
									[]int{0,1,1,1,1,1,0,0,1,0},
									[]int{0,0,1,0,1,1,1,1,0,1},
									[]int{0,1,1,0,0,0,1,1,1,1},
									[]int{0,0,1,0,0,1,0,1,0,1},
									[]int{1,0,1,0,1,1,0,0,0,0},
									[]int{0,0,0,0,1,1,0,0,0,1}}))
}
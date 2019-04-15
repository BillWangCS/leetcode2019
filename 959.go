/*
In a N x N grid composed of 1 x 1 squares, each 1 x 1 square consists of a /, \, or blank space.
These characters divide the square into contiguous regions.

(Note that backslash characters are escaped, so a \ is represented as "\\".)

Return the number of regions.



Example 1:

Input:
[
  " /",
  "/ "
]
Output: 2
Explanation: The 2x2 grid is as follows:

Example 2:

Input:
[
  " /",
  "  "
]
Output: 1
Explanation: The 2x2 grid is as follows:

Example 3:

Input:
[
  "\\/",
  "/\\"
]
Output: 4
Explanation: (Recall that because \ characters are escaped, "\\/" refers to \/, and "/\\" refers to /\.)
The 2x2 grid is as follows:

Example 4:

Input:
[
  "/\\",
  "\\/"
]
Output: 5
Explanation: (Recall that because \ characters are escaped, "/\\" refers to /\, and "\\/" refers to \/.)
The 2x2 grid is as follows:

Example 5:

Input:
[
  "//",
  "/ "
]
Output: 3
Explanation: The 2x2 grid is as follows:



Note:

1 <= grid.length == grid[0].length <= 30
grid[i][j] is either '/', '\', or ' '.
 */

package main

import "fmt"

type cell struct {
	left int
	up int
	right int
	down int
}

func getRoot(u int, father []int) int {
	if father[u] == u {
		return u
	} else {
		father[u] = getRoot(father[u], father)
		return father[u]
	}
}

func merge(u, v int, father *[]int) {
	fu := getRoot(u, *father)
	fv := getRoot(v, *father)
	if fu != fv {
		(*father)[fu] = fv
	}
}


func regionsBySlashes(grid []string) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	cnt := 0
	table := make([][]cell, 0)
	father := make([]int, 0)
	for _, s := range grid {
		row := make([]cell, 0)
		for _, v := range s {
			c := cell{cnt, cnt+1, cnt+2, cnt+3}
			father = append(father, cnt)
			father = append(father, cnt+1)
			father = append(father, cnt+2)
			father = append(father, cnt+3)
			row = append(row, c)
			cnt += 4
			v = v
		}
		table = append(table, row)
	}
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[0]); j++ {
			if i - 1 >= 0 {
				merge(table[i][j].up, table[i - 1][j].down, &father)
			}
			if j - 1 >= 0 {
				merge(table[i][j].left, table[i][j - 1].right, &father)
			}
			if byte(grid[i][j]) == '\\' {
				merge(table[i][j].down, table[i][j].left, &father)
				merge(table[i][j].right, table[i][j].up, &father)
			} else if byte(grid[i][j]) == '/' {
				merge(table[i][j].up, table[i][j].left, &father)
				merge(table[i][j].down, table[i][j].right, &father)
			} else {
				merge(table[i][j].up, table[i][j].left, &father)
				fmt.Println(father)
				merge(table[i][j].up, table[i][j].right, &father)
				fmt.Println(father)
				merge(table[i][j].up, table[i][j].down, &father)
			}
		}
	}
	visited := make(map[int]bool)
	for i := 0; i < len(father); i++ {
		visited[getRoot(father[i], father)] = true
	}
	fmt.Println(father)
	return len(visited)
}


func main() {
	fmt.Println(regionsBySlashes([]string{"/\\",
		"\\/"}))
}
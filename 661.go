/*
Given a 2D integer matrix M representing the gray scale of an image,
you need to design a smoother to make the gray scale of each cell becomes the average gray scale (rounding down)
of all the 8 surrounding cells and itself. If a cell has less than 8 surrounding cells, then use as many as you can.

Example 1:
Input:
[[1,1,1],
 [1,0,1],
 [1,1,1]]
Output:
[[0, 0, 0],
 [0, 0, 0],
 [0, 0, 0]]
Explanation:
For the point (0,0), (0,2), (2,0), (2,2): floor(3/4) = floor(0.75) = 0
For the point (0,1), (1,0), (1,2), (2,1): floor(5/6) = floor(0.83333333) = 0
For the point (1,1): floor(8/9) = floor(0.88888889) = 0
Note:
The value in the given matrix is in the range of [0, 255].
The length and width of the given matrix are in the range of [1, 150].
 */

package main

import "fmt"

func imageSmoother(M [][]int) [][]int {
	r := len(M)
	res := make([][]int, 0)
	if r == 0 {
		return res
	}
	c := len(M[0])
	for i := 0; i < r; i++ {
		row := make([]int, 0)
		for j := 0; j < c; j++ {
			cnt := 0
			v := M[i][j]
			cnt++
			if i-1 >= 0 {
				v += M[i-1][j]
				cnt++
				if j-1 >= 0 {
					v += M[i-1][j-1]
					cnt++
				}
				if j+1 < c {
					v += M[i-1][j+1]
					cnt++
				}
			}
			if i+1 < r {
				fmt.Println(i+1, j)
				v += M[i+1][j]
				cnt++
				if j-1 >= 0 {
					v += M[i+1][j-1]
					cnt++
				}
				if j+1 < c {
					v += M[i+1][j+1]
					cnt++
				}
			}
			if j-1 >= 0 {
				v += M[i][j-1]
				cnt++
			}
			if j+1 < c {
				v += M[i][j+1]
				cnt++
			}
			row = append(row, v/cnt)
		}
		res = append(res, row)
	}
	return res
}

func main() {
	fmt.Println(imageSmoother([][]int{[]int{2,3}}))
}
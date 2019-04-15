/*
Given a matrix A, return the transpose of A.

The transpose of a matrix is the matrix flipped over it's main diagonal,
switching the row and column indices of the matrix.



Example 1:

Input: [[1,2,3],
		[4,5,6],
		[7,8,9]]
		[1,4,7],
		[2,5,8],
		[3,6,9]]
Output: [[1,4,7],[2,5,8],[3,6,9]]
Example 2:

Input: [[1,2,3],
		[4,5,6]]
Output: [[1,4],[2,5],[3,6]]


Note:

1 <= A.length <= 1000
1 <= A[0].length <= 1000
 */

package main

import "fmt"

func transpose(A [][]int) [][]int {
	l := len(A)
	row := make([]int, 0)
	res := make([][]int, 0)
	if l == 0 {
		return res
	}
	for j := 0; j < len(A[0]); j++ {
		row = nil
		for i := 0; i < l; i++ {
			row = append(row, A[i][j])
		}
		res = append(res, row)
	}
	return res
}

func main() {
	fmt.Println(transpose([][]int{[]int{1,2,3},[]int{4,5,6}}))
}
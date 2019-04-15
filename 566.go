/*
In MATLAB, there is a very useful function called 'reshape',
which can reshape a matrix into a new one with different size but keep its original data.

You're given a matrix represented by a two-dimensional array,
and two positive integers r and c representing the row number and column number of the wanted reshaped matrix,
respectively.

The reshaped matrix need to be filled with all the elements of the original matrix
in the same row-traversing order as they were.

If the 'reshape' operation with given parameters is possible and legal,
output the new reshaped matrix; Otherwise, output the original matrix.

Example 1:
Input:
nums =
[[1,2],
 [3,4]]
r = 1, c = 4
Output:
[[1,2,3,4]]
Explanation:
The row-traversing of nums is [1,2,3,4]. The new reshaped matrix is a 1 * 4 matrix,
fill it row by row by using the previous list.
Example 2:
Input:
nums =
[[1,2],
 [3,4]]
r = 2, c = 4
Output:
[[1,2],
 [3,4]]
Explanation:
There is no way to reshape a 2 * 2 matrix to a 2 * 4 matrix. So output the original matrix.
 */

package main

import "fmt"

func matrixReshape(nums [][]int, r int, c int) [][]int {
	res := make([][]int, 0)
	traversalArr := make([]int, 0)
	r1 := len(nums)
	c1 := len(nums[0])
	if r1 == 0 || r == 0 || c == 0 || c1 == 0 {
		return nums
	}
	if r * c != r1 * c1 {
		return nums
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[0]); j++ {
			traversalArr = append(traversalArr, nums[i][j])
		}
	}
	cur := 0
	for i := 0; i < r; i++ {
		row := make([]int, 0)
		for j := 0; j < c; j++ {
			row = append(row, traversalArr[cur])
			cur++
		}
		res = append(res, row)
	}
	return res
}

func main() {
	fmt.Println(matrixReshape([][]int{[]int{1,2}, []int{3,4}}, 2, 4))
}
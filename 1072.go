/*
Given a matrix consisting of 0s and 1s, we may choose any number of columns in the matrix and flip every cell in that column.  Flipping a cell changes the value of that cell from 0 to 1 or from 1 to 0.

Return the maximum number of rows that have all values equal after some number of flips.



Example 1:

Input: [[0,1],[1,1]]
Output: 1
Explanation: After flipping no values, 1 row has all values equal.
Example 2:

Input: [[0,1],[1,0]]
Output: 2
Explanation: After flipping values in the first column, both rows have equal values.
Example 3:

Input: [[0,0,0],[0,0,1],[1,1,0]]
Output: 2
Explanation: After flipping values in the first two columns, the last two rows have equal values.


Note:

1 <= matrix.length <= 300
1 <= matrix[i].length <= 300
All matrix[i].length's are equal
matrix[i][j] is 0 or 1
 */

package main

import "fmt"

func maxEqualRowsAfterFlips(matrix [][]int) int {
	r := len(matrix)
	c := len(matrix[0])
	max := 0
	if c == 1 {
		return r
	}
	for i := 0; i < r; i++ {
		cnt := 0
		for j := 0; j < r; j++ {
			if i == j {
				continue
			}
			s1 := 1
			s2 := 0
			for k := 0; k < c; k++ {
				tmp := matrix[i][k] ^ matrix[j][k]
				s2 |= tmp
				s1 &= tmp
			}
			if s1 == 1 {
				cnt++
			}
			if s2 == 0 {
				cnt++
			}
		}
		if cnt > max {
			max = cnt
		}
	}
	return max+1
}

func main() {
	fmt.Println(maxEqualRowsAfterFlips([][]int{[]int{0,0,0}, []int{0,0,1}, []int{1,1,0}}))
}

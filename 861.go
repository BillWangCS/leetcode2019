/*
We have a two dimensional matrix A where each value is 0 or 1.

A move consists of choosing any row or column, and toggling each value in that row or column:
changing all 0s to 1s, and all 1s to 0s.

After making any number of moves, every row of this matrix is interpreted as a binary number,
and the score of the matrix is the sum of these numbers.

Return the highest possible score.



Example 1:

Input: [
[1,1,0,0],
[1,0,1,0],
[1,1,0,0]
]
Output: 39
Explanation:
Toggled to [[1,1,1,1],[1,0,0,1],[1,1,1,1]].
0b1111 + 0b1001 + 0b1111 = 15 + 9 + 15 = 39


Note:

1 <= A.length <= 20
1 <= A[0].length <= 20
A[i][j] is 0 or 1.
 */

package main

import "fmt"

func toggleArray(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			arr[i] = 1
		} else {
			arr[i] = 0
		}
	}
	return arr
}

func sumFunc(A []int) int {
	sum := 0
	for i := 0; i < len(A); i++ {
		sum = 2 * sum + A[i]
	}
	return sum
}

func matrixScore(A [][]int) int {
	res := 0
	l := len(A)
	if l == 0 {
		return 0
	}
	for i := 0; i < l; i++ {
		arr := make([]int, 0)
		arr = append(arr, A[i]...)
		if sumFunc(arr) > sumFunc(toggleArray(A[i])) {
			A[i] = arr
		}

	}
	cl := len(A[0])
	fmt.Println("aaa", A[0])
	for j := 0; j < cl; j++ {
		column := make([]int, 0)
		cnt := 0
		for i := 0; i < l; i++ {
			if A[i][j] == 1 {
				cnt++
			}
			column = append(column, A[i][j])
		}
		fmt.Println(cnt)
		if cnt <= l/2 {
			toggleArray(column)
			for x, v := range column {
				A[x][j] = v
			}
		}
	}
	fmt.Println(A)
	for i := 0; i < l; i++ {
		res += sumFunc(A[i])
	}
	return res
}

func main() {
	fmt.Println(matrixScore([][]int{[]int{1,1,0,0},[]int{1,0,1,0},[]int{1,1,0,0}}))
}
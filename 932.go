/*
Given a square array of integers A, we want the minimum sum of a falling path through A.

A falling path starts at any element in the first row, and chooses one element from each row.
The next row's choice must be in a column that is different from the previous row's column by at most one.



Example 1:

Input: [[1,2,3],[4,5,6],[7,8,9]]
Output: 12
Explanation:
The possible falling paths are:
[1,4,7], [1,4,8], [1,5,7], [1,5,8], [1,5,9]
[2,4,7], [2,4,8], [2,5,7], [2,5,8], [2,5,9], [2,6,8], [2,6,9]
[3,5,7], [3,5,8], [3,5,9], [3,6,8], [3,6,9]
The falling path with the smallest sum is [1,4,7], so the answer is 12.



Note:

1 <= A.length == A[0].length <= 100
-100 <= A[i][j] <= 100
 */

package main

import (
	"fmt"
	"math"
)

/*
超时的DFS解法
func dfs_932(layer int, pos int, A [][]int, path *[]int, sums *[]int) {
	if layer == len(A) - 1 {
		*path = append(*path, A[layer][pos])
		sum := 0
		for i := 0; i < len(*path); i++ {
			sum += (*path)[i]
		}
		*sums = append(*sums, sum)
		*path = (*path)[:len(*path)-1]
		return
	}
	if pos > 0 {
		*path = append(*path, A[layer][pos])
		dfs_932(layer + 1, pos-1, A, path, sums)
		*path = (*path)[:len(*path)-1]
	}
	*path = append(*path, A[layer][pos])
	dfs_932(layer + 1, pos, A, path, sums)
	*path = (*path)[:len(*path)-1]
	if pos+1 < len(A[0]) {
		*path = append(*path, A[layer][pos])
		dfs_932(layer + 1, pos+1, A, path, sums)
		*path = (*path)[:len(*path)-1]
	}
}

func maxFallingSum(pos int, A [][]int) int {
	res := 100000000
	path := make([]int, 0)
	sums := make([]int, 0)
	dfs_932(0, pos, A, &path, &sums)
	for i := 0; i < len(sums); i++ {
		if res > sums[i] {
			res = sums[i]
		}
	}
	return res
}

func minFallingPathSum(A [][]int) int {
	res := 100000000
	for i := 0; i < len(A[0]); i++ {
		s := maxFallingSum(i, A)
		fmt.Println(s)
		if res > s {
			res = s
		}
	}
	return res
}
*/
func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minFallingPathSum(A [][]int) int {
	res := math.MaxInt32
	for i := len(A) - 2; i >= 0; i-- {
		for j := 0; j < len(A[0]); j++ {
			min := A[i+1][j]
			if j > 0 {
				min = minInt(A[i+1][j-1], min)
			}
			if j + 1 < len(A[0]) {
				min = minInt(A[i+1][j+1], min)
			}
			A[i][j] += min
		}
	}
	for i := 0; i < len(A); i++ {
		if res > A[0][i] {
			res = A[0][i]
		}
	}
	return res
}

func main() {
	fmt.Println(minFallingPathSum([][]int{[]int{1,2,3}, []int{4,5,6}, []int{7,8,9}}))
	//{[]int{-19,-1,-96,48,-94,36,16,55,-42,37,-59,6,-32,96,95,-58,13,-34,94,85},[]int{17,44,36,-29,84,80,-34,50,-99,64,13,91,-27,25,-36,57,20,98,-100,-72},[]int{-92,-75,86,90,-4,90,64,56,50,-63,10,-15,90,-66,-66,32,-69,-78,1,60},[]int{21,51,-47,-43,-14,99,44,90,8,11,99,-62,57,59,69,50,-69,32,85,13},[]int{-28,90,12,-18,23,61,-55,-97,6,89,36,26,26,-1,46,-50,79,-45,89,86},[]int{-85,-10,49,-10,2,62,41,92,-67,85,86,27,89,-50,77,55,22,-82,-94,-98},[]int{-50,53,-23,55,25,-22,76,-93,-7,66,-75,42,-35,-96,-5,4,-92,13,-31,-100},[]int{-62,-78,8,-92,86,69,90,-37,81,97,53,-45,34,19,-19,-39,-88,-75,-74,-4},[]int{29,53,-91,65,-92,11,49,26,90,-31,17,-84,12,63,-60,-48,40,-49,-48,88},[]int{100,-69,80,11,-93,17,28,-94,52,64,-86,30,-9,-53,-8,-68,-33,31,-5,11},[]int{9,64,-31,63,-84,-15,-30,-10,67,2,98,73,-77,-37,-96,47,-97,78,-62,-17},[]int{-88,-38,-22,-90,54,42,-29,67,-85,-90,-29,81,52,35,13,61,-18,-94,61,-62},[]int{-23,-29,-76,-30,-65,23,31,-98,-9,11,75,-1,-84,-90,73,58,72,-48,30,-81},[]int{66,-33,91,-6,-94,82,25,-43,-93,-25,-69,10,-71,-65,85,28,-52,76,25,90},[]int{-3,78,36,-92,-52,-44,-66,-53,-55,76,-7,76,-73,13,-98,86,-99,-22,61,100},[]int{-97,65,2,-93,56,-78,22,56,35,-24,-95,-13,83,-34,-51,-73,2,7,-86,-19},[]int{32,94,-14,-13,-6,-55,-21,29,-21,16,67,100,77,-26,-96,22,-5,-53,-92,-36},[]int{60,93,-79,76,-91,43,-95,-16,74,-21,85,43,21,-68,-32,-18,18,100,-43,1},[]int{87,-31,26,53,26,51,-61,92,-65,17,-41,27,-42,-14,37,-46,46,-31,-74,23},[]int{-67,-14,-20,-85,42,36,56,9,11,-66,-59,-55,5,64,-29,77,47,44,-33,-77}}
}

/*
1,2,3
4,5,6
7,8,9
 */
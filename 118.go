/*
Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.


In Pascal's triangle, each number is the sum of the two numbers directly above it.

Example:

Input: 5
Output:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
 */

package main

import "fmt"

func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{[]int{1}}
	}
	if numRows == 2 {
		return [][]int{[]int{1}, []int{1,1}}
	}
	last := []int{1, 1}
	res := make([][]int, 0)
	if numRows == 0 {
		return res
	}
	res = append(res, []int{1})
	res = append(res, []int{1,1})
	for i := 2; i < numRows; i++ {
		tmp := make([]int, 0)
		tmp = append(tmp, 1)
		for j := 0; j+1 < len(last); j++ {
			tmp = append(tmp, last[j]+last[j+1])
		}
		tmp = append(tmp, 1)
		last = tmp
		res = append(res, tmp)
	}
	return res
}

func main() {
	fmt.Println(generate(5))
}
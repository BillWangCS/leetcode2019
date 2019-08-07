/*
Given a non-negative index k where k ≤ 33, return the kth index row of the Pascal's triangle.

Note that the row index starts from 0.


In Pascal's triangle, each number is the sum of the two numbers directly above it.

Example:

Input: 3
Output: [1,3,3,1]
Follow up:

Could you optimize your algorithm to use only O(k) extra space?
 */

package main

import "fmt"

func getRow(rowIndex int) []int {
	arr := make([]int, 0)
	arr = append(arr, 1)
	for i := 1; i <= rowIndex; i++ {
		tmp := make([]int, 0)
		tmp = append(tmp, 1)
		for j := 0; j + 1 < len(arr); j++ {
			tmp = append(tmp, arr[j] + arr[j+1])
		}
		tmp = append(tmp, 1)
		arr = tmp
	}
	return arr
}

func main() {
	fmt.Println(getRow(4))
}

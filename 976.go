/*
Given an array A of positive lengths, return the largest perimeter of a triangle with non-zero area,
formed from 3 of these lengths.

If it is impossible to form any triangle of non-zero area, return 0.



Example 1:

Input: [2,1,2]
Output: 5
Example 2:

Input: [1,2,1]
Output: 0
Example 3:

Input: [3,2,3,4]
Output: 10
Example 4:

Input: [3,6,2,3]
Output: 8


Note:

3 <= A.length <= 10000
1 <= A[i] <= 10^6
 */

package main

import (
	"fmt"
	"sort"
)

func isTriangle(a, b, c int) bool {
	arr := []int{a, b, c}
	sort.Ints(arr)
	if arr[0] + arr[1] > arr[2] {
		return true
	}
	return false
}

func largestPerimeter(A []int) int {
	sort.Ints(A)
	c := len(A) - 1
	b := c - 1
	for c > 1 {
		b = c - 1
		a := b - 1
		if A[a] > A[c] - A[b] {
			return A[a] + A[b] + A[c]
		}
		c--
	}
	return 0
}

/*
Break through time limit exceeded.
func largestPerimeter(A []int) int {
	res := 0
	c := len(A) - 1
	b := c - 1
	a := b - 1
	sort.Ints(A)
	for c >= 2 {
		b = c - 1
		a = b - 1
		if A[a] + A[b] + A[c] < res {
			break
		}
		for b >= 1 {
			a = b - 1
			if A[a] + A[b] + A[c] < res {
				break
			}
			for a >= 0 {
				if A[a] + A[b] + A[c] < res {
					break
				}
				if isTriangle(A[a], A[b], A[c]) {
					res = A[a] + A[b] + A[c]
				}
				a--
			}
			b--
		}
		c--
	}
	return res
}
*/

func main() {
	fmt.Println(largestPerimeter([]int{3,2,3,4}))
}
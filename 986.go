/*
Given two lists of closed intervals, each list of intervals is pairwise disjoint and in sorted order.

Return the intersection of these two interval lists.

(Formally, a closed interval [a, b] (with a <= b) denotes the set of real numbers x with a <= x <= b.
The intersection of two closed intervals is a set of real numbers that is either empty,
or can be represented as a closed interval.  For example, the intersection of [1, 3] and [2, 4] is [2, 3].)



Example 1:



Input: A = [[0,2],[5,10],[13,23],[24,25]], B = [[1,5],[8,12],[15,24],[25,26]]
Output: [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
Reminder: The inputs and the desired output are lists of Interval objects, and not arrays or lists.


Note:

0 <= A.length < 1000
0 <= B.length < 1000
0 <= A[i].start, A[i].end, B[i].start, B[i].end < 10^9
NOTE: input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.
 */

package main

import "fmt"

func maxInteger(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func getIntersection(A []int, B []int) []int {
	res := make([]int, 0)
	left := maxInteger(A[0], B[0])
	right := minInt(A[len(A)-1], B[len(B)-1])
	res = append(res, left)
	res = append(res, right)
	return res
}

func intervalIntersection(A [][]int, B [][]int) [][]int {
	res := make([][]int, 0)
	i := 0
	j := 0
	for i < len(A) && j < len(B) {
		intersection := getIntersection(A[i], B[j])
		if intersection[0] <= intersection[1] {
			res = append(res, intersection)
		}
		if A[i][len(A[i])-1] > B[j][len(B[j])-1] {
			j++
		} else if A[i][len(A[i])-1] < B[j][len(B[j])-1] {
			i++
		} else {
			i++
			j++
		}
	}
	return res
}

func main() {
	fmt.Println(intervalIntersection([][]int{[]int{0,2}, []int{5,10}, []int{13,23}, []int{24,25}}, [][]int{[]int{1,5}, []int{8,12}, []int{15,24}, []int{25,26}}))
}
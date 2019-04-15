/*
Let's call an array A a mountain if the following properties hold:

A.length >= 3
There exists some 0 < i < A.length - 1 such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]
Given an array that is definitely a mountain, return any i such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1].

Example 1:

Input: [0,1,0]
Output: 1
Example 2:

Input: [0,2,1,0]
Output: 1
Note:

3 <= A.length <= 10000
0 <= A[i] <= 10^6
A is a mountain, as defined above.
 */

package main

import "fmt"

func peakIndexInMountainArray(A []int) int {
	start, end := 0, len(A)-1
	for start + 1 < end {
		mid := start + (end - start)/2
		if A[mid] < A[mid+1] {
			start = mid
		} else if A[mid] > A[mid+1] {
			end = mid
		} else {
			return -1
		}
	}
	if A[start] > A[start+1] {
		return start
	}
	if A[end] > A[end+1] {
		return end
	}
	return -1
}

func main() {
	fmt.Println(peakIndexInMountainArray([]int{3,4,5,1}))
}
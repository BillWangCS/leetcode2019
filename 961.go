/*
In a array A of size 2N, there are N+1 unique elements, and exactly one of these elements is repeated N times.

Return the element repeated N times.



Example 1:

Input: [1,2,3,3]
Output: 3
Example 2:

Input: [2,1,2,5,3,2]
Output: 2
Example 3:

Input: [5,1,5,2,5,3,5,4]
[1,2,5,5,5,3,5,4]
Output: 5


Note:

4 <= A.length <= 10000
0 <= A[i] < 10000
A.length is even
 */
package main

import "fmt"

func quickFind(A[]int, start, end ,pos int) int {
	if start >= end {
		return -1
	}
	left, right := start, end
	pivot := A[(start + end)/2]
	for left <= right {
		for left <= right && A[left] < pivot {
			left++
		}
		for left <= right && A[right] > pivot {
			right --
		}
		if left <= right {
			A[left], A[right] = A[right], A[left]
			left++
			right--
		}
	}
	if pos <= right {
		quickFind(A, start, right, pos)
	}
	if pos >= left {
		quickFind(A, left, end, pos)
	}
	return A[pos]
}

func repeatedNTimes(A []int) int {
	res := quickFind(A, 0, len(A) - 1, (len(A)-1)/2)
	if res == A[len(A)/2] || res == A[(len(A)-1)/2-1] {
		return res
	} else {
		if A[len(A)/2] == A[len(A)/2 +1] {
			return A[len(A)/2]
		}
	}
	return res
}

func main() {
	fmt.Println(repeatedNTimes([]int{5,9,3,3}))
}


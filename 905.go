/*
Given an array A of non-negative integers, return an array consisting of all the even elements of A, followed by all the odd elements of A.

You may return any answer array that satisfies this condition.



Example 1:

Input: [3,1,2,4]
Output: [2,4,3,1]
The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.


Note:

1 <= A.length <= 5000
0 <= A[i] <= 5000
 */

package main

import "fmt"

func quicksort(A []int, start, end int) {
	if start >= end {
		return
	}
	left, right := start, end
	for left <= right {
		for left <= right && A[left]%2 == 1 {
			left++
		}
		for left <= right && A[right]%2 == 0 {
			right--
		}
		if left <= right {
			A[left], A[right] = A[right], A[left]
			left++
			right--
		}
	}
	return
}

func sortArrayByParity(A []int) []int {
	l := len(A)
	if l <= 1 {
		return A
	}
	quicksort(A, 0, l-1)
	fmt.Println("quick", A)
	left, right := 1, l-1
	if l % 2 == 0 {
		left, right = 1, l-2
	}
	for left <= right {
		A[left], A[right] = A[right], A[left]
		left += 2
		right -= 2
	}
	return A
}

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
}

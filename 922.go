/*
Given an array A of non-negative integers, half of the integers in A are odd, and half of the integers are even.

Sort the array so that whenever A[i] is odd, i is odd; and whenever A[i] is even, i is even.

You may return any answer array that satisfies this condition.



Example 1:

Input: [4,2,5,7]
2,4,6,1,3,5
2,5,6,8,1,4,5,7

Output: [4,5,2,7]
Explanation: [4,7,2,5], [2,5,4,7], [2,7,4,5] would also have been accepted.


Note:

2 <= A.length <= 20000
A.length % 2 == 0
0 <= A[i] <= 1000
 */

package main

func partition(A []int, start, end int) {
	left, right := start, end
	for left <= right {
		for left <= right && A[left]%2 == 0 {
			left++
		}
		for left <= right && A[right]%2 == 1 {
			right--
		}
		if left <= right {
			A[left], A[right] = A[right], A[left]
			left++
			right--
		}
	}
}

func sortArrayByParityII(A []int) []int {
	l := len(A)
	partition(A, 0, l-1)
	for i := 1; i < l-i-1; i += 2 {
		A[i], A[l-i-1] = A[l-i-1], A[i]
	}
	return A
}

func main() {
	sortArrayByParityII([]int{4,2,5,7})
}
/*
Given target, a non-negative integer k and an integer array A sorted in ascending order,
find the k closest numbers to target in A, sorted in ascending order by the difference between the number and target.
Otherwise, sorted in ascending order by number if the difference is same.

Example
Example 1:

Input: A = [1, 2, 3], target = 2, k = 3
Output: [2, 1, 3]
Example 2:

Input: A = [1, 4, 6, 8], target = 3, k = 3
Output: [4, 1, 6]
Challenge
O(logn + k) time

Notice
The value k is a non-negative integer and will always be smaller than the length of the sorted array.
Length of the given array is positive and will not exceed
1
0
4
10
​4
​​
Absolute value of elements in the array will not exceed
1
0
4
10
​4
 */

package main

import "fmt"

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

func reverse(A []int) {
	l := len(A)
	for i := 0; i < l/2; i++ {
		j := l - i - 1
		A[i], A[j] = A[j], A[i]
	}
}

func kClosestNumbers (A []int, target int, k int) []int {
	// write your code here
	left, right := 0, len(A) - 1
	if target <= A[0] {
		return A[:k]
	}
	if target >= A[right] {
		tmp := A[right-k+1:]
		reverse(tmp)
		return tmp
	}
	pos := -1
	for left + 1 < right {
		mid := (left + right) / 2
		if A[mid] == target || A[mid] < target && A[mid+1] > target {
			pos = mid
			break
		}
		if A[mid] < target {
			left = mid
		} else if A[mid] > target {
			right = mid
		}
	}
	if pos == -1 {
		if target == A[left] || target > A[left] && target < A[right] {
			pos = left
		} else if target == A[right] {
			pos = right
		}
	}
	res := make([]int, 0)
	if pos != -1 {
		l := pos
		r := pos+1
		i := 0
		if pos == len(A) - 1 {
			tmp := A[pos-k+1:]
			reverse(tmp)
			return tmp
		}
		for i < k {
			if l >= 0 && (r >= len(A) || abs(target - A[l]) <= abs(target - A[r])) {
				res = append(res, A[l])
				l--
				i++
			} else {
				fmt.Println(target, l, r)
				res = append(res, A[r])
				r++
				i++
			}
		}
		fmt.Println(l, r)
	}
	return res
}

func main() {
	fmt.Println(kClosestNumbers([]int{1,4,6,10,20}, 21, 4))
}
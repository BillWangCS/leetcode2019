/*
For a non-negative integer X, the array-form of X is an array of its digits in left to right order.
For example, if X = 1231, then the array form is [1,2,3,1].

Given the array-form A of a non-negative integer X, return the array-form of the integer X+K.



Example 1:

Input: A = [1,2,0,0], K = 34
Output: [1,2,3,4]
Explanation: 1200 + 34 = 1234
Example 2:

Input: A = [2,7,4], K = 181
Output: [4,5,5]
Explanation: 274 + 181 = 455
Example 3:

Input: A = [2,1,5], K = 806
Output: [1,0,2,1]
Explanation: 215 + 806 = 1021
Example 4:

Input: A = [9,9,9,9,9,9,9,9,9,9], K = 1
Output: [1,0,0,0,0,0,0,0,0,0,0]
Explanation: 9999999999 + 1 = 10000000000


Note：

1 <= A.length <= 10000
0 <= A[i] <= 9
0 <= K <= 10000
If A.length > 1, then A[0] != 0
 */

package main

import "fmt"

func reverseArr(A *[]int) {
	i := 0
	j := len(*A) - 1
	for i <= j {
		(*A)[i], (*A)[j] = (*A)[j], (*A)[i]
		i++
		j--
	}
}

func maxInteger(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func addToArrayForm(A []int, K int) []int {
	B := make([]int, 0)
	res := make([]int, 0)
	for K > 0 {
		n := K % 10
		B = append(B, n)
		K = K / 10
	}
	reverseArr(&B)
	la := len(A)
	lb := len(B)
	l := maxInteger(la, lb)
	c := 0
	s := 0
	for i := 1; i <= l; i++ {
		if i > la {
			s = B[lb-i] + c
		} else if i > lb {
			s = A[la-i] + c
		} else {
			s = A[la-i] + B[lb-i] + c
		}
		if s >= 10 {
			c = 1
		} else {
			c = 0
		}
		res = append(res, s%10)
	}
	if c > 0 {
		res = append(res, c)
	}
	reverseArr(&res)
	return res
}

func main() {
	fmt.Println(addToArrayForm([]int{9,9,9,9,9,9,9,9,9,9}, 1))
}
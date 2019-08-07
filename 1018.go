/*
Given an array A of 0s and 1s, consider N_i: the i-th subarray from A[0] to A[i] interpreted
as a binary number (from most-significant-bit to least-significant-bit.)

Return a list of booleans answer, where answer[i] is true if and only if N_i is divisible by 5.

Example 1:

Input: [0,1,1]
Output: [true,false,false]
Explanation:
The input numbers in binary are 0, 01, 011; which are 0, 1, and 3 in base-10.
Only the first number is divisible by 5, so answer[0] is true.
Example 2:

Input: [1,1,1]
Output: [false,false,false]
Example 3:

Input: [0,1,1,1,1,1]
Output: [true,false,false,false,true,false]
Example 4:

Input: [1,1,1,0,1]
Output: [false,false,false,false,false]


Note:

1 <= A.length <= 30000
A[i] is 0 or 1
 */

package main

import "fmt"

//http://www.dewen.net.cn/q/1827/判断一个二进制表示数能否被5整除的算法

func prefixesDivBy5(A []int) []bool {
	res := make([]bool, 0)
	sum1 := 0
	sum2 := 0
	flag1 := true
	flag2 := true
	for i := 0; i < len(A); i++ {
		if i % 2 == 0 {
			if flag1 {
				sum1 += A[i]
				flag1 = !flag1
			} else {
				sum1 -= A[i]
				flag1 = !flag1
			}
		} else {
			if flag2 {
				sum2 += A[i]
				flag2 = !flag2
			} else {
				sum2 -= A[i]
				flag2 = !flag2
			}
		}
		res = append(res, (2*sum1 + sum2) % 5 == 0)
	}
	return res
}

func main() {
	fmt.Println(prefixesDivBy5([]int{0,1,1,1,1,1}))
}
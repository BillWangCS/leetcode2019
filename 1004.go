/*
Given an array A of 0s and 1s, we may change up to K values from 0 to 1.

Return the length of the longest (contiguous) subarray that contains only 1s.



Example 1:

Input: A = [1,1,1,0,0,0,1,1,1,1,0], K = 2
Output: 6
Explanation:
[1,1,1,0,0,1,1,1,1,1,1]
Bolded numbers were flipped from 0 to 1.  The longest subarray is underlined.
Example 2:

Input: A = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
Output: 10
Explanation:
[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
Bolded numbers were flipped from 0 to 1.  The longest subarray is underlined.


Note:

1 <= A.length <= 20000
0 <= K <= A.length
A[i] is 0 or 1
 */

package main

import "fmt"

//Method one: Time limit out
/*func countMaxOne(A []int) int {
	max := 0
	cnt := 0
	for i := 0; i < len(A); i++ {
		if A[i] == 1 {
			cnt++
			if cnt > max {
				max = cnt
			}
		} else {
			cnt = 0
		}
	}
	return max
}

func longestOnes(A []int, K int) int {
	res := 0
	l, r := 0, 0
	dict := make(map[int]bool)
	n := K
	for i := 0 ; i < len(A); i++ {
		if A[i] == 0 {
			l = i
			r = l
			for r < len(A) && n > 0 {
				if A[r] == 0 {
					A[r] = 1
					dict[r] = true
					n--
				}
				r++
			}
			x := countMaxOne(A)
			if x > res {
				res = x
			}
			for k, _ := range dict {
				if dict[k] {
					A[k] = 0
					dict[k] = false
				}
			}
			n = K
		}
	}
	return res
}*/

//Method Two
func longestOnes(A []int, K int) int {
	res := 0
	l, r := 0, 0
	cnt := 0
	for r < len(A) {
		if A[r] == 0 {
			cnt++
		}
		for cnt > K {
			if A[l] == 0 {
				cnt--
			}
			l++
		}
		if res < r-l+1 {
			res = r-l+1
		}
		r++
	}
	return res
}

func main() {
	fmt.Println(longestOnes([]int{1,1,1,0,0,0,1,1,1,1,0}, 2))
}
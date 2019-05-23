/*
We are given two strings, A and B.

A shift on A consists of taking string A and moving the leftmost character to the rightmost position. For example,
if A = 'abcde', then it will be 'bcdea' after one shift on A. Return True if and only if A can become B after some number of shifts on A.

Example 1:
Input: A = 'abcdecd', B = 'cdecdab'
Output: true

Example 2:
Input: A = 'abcde', B = 'abced'
Output: false
Note:

A and B will have length at most 100.
 */

package main

import (
	"fmt"
	"strings"
)

func rotateString(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	if len(A) == 0 && len(B) == 0 {
		return true
	}
	poses := make([]int, 0)
	for i := 0; i < len(A); i++ {
		if A[i] == B[0] {
			poses = append(poses, i)
		}
	}
	if len(poses) == 0 {
		return false
	}
	for _ ,pos := range poses {
		if strings.HasPrefix(B, A[pos:]) && strings.HasPrefix(A, B[len(A) - pos:]) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(rotateString("ohbrwzxvxe", "uornhegseo"))
}
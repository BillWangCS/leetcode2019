/*
Given a binary string S (a string consisting only of '0' and '1's) and a positive integer N,
return true if and only if for every integer X from 1 to N, the binary representation of X is a substring of S.



Example 1:

Input: S = "0110", N = 3
Output: true
Example 2:

Input: S = "0110", N = 4
Output: false


Note:

1 <= S.length <= 1000
1 <= N <= 10^9
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convertToBinaryStr(N int) string {
	str := ""
	for N > 0 {
		x := N % 2
		N = N >> 1
		str = strconv.Itoa(x) + str
	}
	return str
}

func queryString(S string, N int) bool {
	for i := 1; i <= N; i ++ {
		if !strings.Contains(S, convertToBinaryStr(i)) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(queryString("110101011011000011011111000000", 15))
}
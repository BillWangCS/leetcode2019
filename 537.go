/*
Given two strings representing two complex numbers.

You need to return a string representing their multiplication. Note i2 = -1 according to the definition.

Example 1:
Input: "1+1i", "1+1i"
Output: "0+2i"
Explanation: (1 + i) * (1 + i) = 1 + i2 + 2 * i = 2i, and you need convert it to the form of 0+2i.
Example 2:
Input: "1+-1i", "1+-1i"
Output: "0+-2i"
Explanation: (1 - i) * (1 - i) = 1 + i2 - 2 * i = -2i, and you need convert it to the form of 0+-2i.
Note:

The input strings will not have extra blank.
The input strings will be given in the form of a+bi,
where the integer a and b will both belong to the range of [-100, 100]. And the output should be also in this form.
 */


package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseComplex(s string) (a, b int) {
	strs := strings.Split(s, "+")
	a, _ = strconv.Atoi(strs[0])
	complexes := strings.Split(strs[1], "i")
	b, _ = strconv.Atoi(complexes[0])
	return a, b
}

func complexNumberMultiply(a string, b string) string {
	x, y := parseComplex(a)
	m, n := parseComplex(b)
	first := m*x-y*n
	last := x*n + m*y
	return strconv.Itoa(first) + "+" + strconv.Itoa(last) + "i"
}

func main() {
	fmt.Println(complexNumberMultiply("1+1i", "1+1i"))
}

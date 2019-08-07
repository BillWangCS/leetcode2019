/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:
(you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);
Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:

P     I    N
A   L S  I G
Y A   H R
P     I
 */

package main

import "fmt"

func convert(s string, numRows int) string {
	l := len(s)
	if numRows == 1 {
		return s
	}
	res := ""
	matrix := make([][]string, 0)
	for i := 0; i < numRows; i++ {
		matrix = append(matrix, make([]string, l))
	}
	m := 0
	n := 0
	i := 0
	for i < len(s) {
		if m == 0 {
			for m < numRows {
				fmt.Println(m, n, i)
				matrix[m][n] = string(s[i])
				m++
				i++
				if i >= l {
					break
				}
			}
			m--
		} else {
			m--
			n++
			matrix[m][n] = string(s[i])
			if m != 0 {
				i++
				if i >= l {
					break
				}
			}
		}
	}
	for i := 0; i < numRows; i++ {
		for j := 0; j < l; j++ {
			if matrix[i][j] != "" {
				res += matrix[i][j]
			}
		}
	}
	return res
}

func main() {
	fmt.Println(convert("AB", 1))
}

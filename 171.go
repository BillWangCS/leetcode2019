/*
Given a column title as appear in an Excel sheet, return its corresponding column number.

For example:

    A -> 1
    B -> 2
    C -> 3
    ...
    Z -> 26
    AA -> 27
    AB -> 28
    ...
Example 1:

Input: "A"
Output: 1
Example 2:

Input: "AB"
Output: 28
Example 3:

Input: "ZY"
Output: 701
 */

package main

import "fmt"

func titleToNumber(s string) int {
	sum := 0
	for _, v := range s {
		num := int(byte(v) - 'A') + 1
		sum = 26 * sum + num
	}
	return sum
}

func main() {
	fmt.Println(titleToNumber("ZY"))
}
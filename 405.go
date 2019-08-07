/*
Given an integer, write an algorithm to convert it to hexadecimal. For negative integer, two’s complement method is used.

Note:

All letters in hexadecimal (a-f) must be in lowercase.
The hexadecimal string must not contain extra leading 0s. If the number is zero,
it is represented by a single zero character '0'; otherwise,
the first character in the hexadecimal string will not be the zero character.
The given number is guaranteed to fit within the range of a 32-bit signed integer.
You must not use any method provided by the library which converts/formats the number to hex directly.
Example 1:

Input:
26

Output:
"1a"
Example 2:

Input:
-1

Output:
"ffffffff"
 */

package main

import (
	"fmt"
	"strconv"
)

func toHex(num int) string {
	res := ""
	if num == 0 {
		return "0"
	}
	x := uint32(num)
	for x != 0 {
		tmp := x & 15
		switch tmp {
		case 10:
			res = "a" + res
		case 11:
			res = "b" + res
		case 12:
			res = "c" + res
		case 13:
			res = "d" + res
		case 14:
			res = "e" + res
		case 15:
			res = "f" + res
		default:
			res = strconv.Itoa(int(tmp)) + res
		}
		x = x >> 4
	}
	return res
}

func main() {
	fmt.Println(toHex(-1))
}
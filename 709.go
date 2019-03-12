/*
Implement function ToLowerCase() that has a string parameter str, and returns the same string in lowercase.



Example 1:

Input: "Hello"
Output: "hello"
Example 2:

Input: "here"
Output: "here"
Example 3:

Input: "LOVELY"
Output: "lovely"
*/

package main

import "fmt"

func toLowerCase(str string) string {
	res := ""
	for _, v := range str {
		if v >= 65 && v <= 90 {
			v = v + 32
		}
		res = res + string(v)
	}
	return res
}

func main() {
	fmt.Println(toLowerCase("aAzZ"))
}
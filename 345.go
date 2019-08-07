/*
Write a function that takes a string as input and reverse only the vowels of a string.

Example 1:

Input: "hello"
Output: "holle"
Example 2:

Input: "leetcode"
Output: "leotcede"
Note:
The vowels does not include the letter "y".
 */

package main

import "fmt"

func reverseVowels(s string) string {
	dict := make(map[byte]bool)
	dict['a'] = true
	dict['e'] = true
	dict['i'] = true
	dict['o'] = true
	dict['u'] = true
	dict['A'] = true
	dict['E'] = true
	dict['I'] = true
	dict['O'] = true
	dict['U'] = true
	vowelStr := ""
	dictPos := make(map[int]bool)
	for i := 0; i < len(s); i++ {
		if dict[byte(s[i])] {
			vowelStr = string(s[i]) + vowelStr
			dictPos[i] = true
		}
	}
	res := ""
	j := 0
	for i := 0; i < len(s); i++ {
		if dictPos[i] {
			res += string(vowelStr[j])
			j++
		} else {
			res += string(s[i])
		}
	}
	return res
}

func main() {
	fmt.Println(reverseVowels("hello"))
}
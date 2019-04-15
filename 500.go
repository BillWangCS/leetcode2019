/*
Given a List of words, return the words that can be typed using letters of alphabet
on only one row's of American keyboard like the image below.






Example:

Input: ["Hello", "Alaska", "Dad", "Peace"]
Output: ["Alaska", "Dad"]


Note:

You may use one character in the keyboard more than once.
You may assume the input string will only contain letters of alphabet.
 */

package main

import (
	"fmt"
	"strings"
)

func findWords(words []string) []string {
	res := make([]string, 0)
	dict := make(map[rune]int)
	for _, v := range "qwertyuiop" {
		dict[v] = 1
	}
	for _, v := range "asdfghjkl" {
		dict[v] = 2
	}
	for _, v := range "zxcvbnm" {
		dict[v] = 3
	}
	for _, s := range words {
		x := 0
		flag := true
		for i, v := range strings.ToLower(s) {
			if i == 0 {
				x = dict[v]
			} else {
				if dict[v] != x {
					flag = false
					break
				}
			}
		}
		if flag {
			res = append(res, s)
		}
	}
	return res
}

func main() {
	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
}


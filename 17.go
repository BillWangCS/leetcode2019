/*
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.



Example:

Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
Note:

Although the above answer is in lexicographical order, your answer could be in any order you want.
 */

package main

import "fmt"

func dfs_17(digits string, pos int, str *string, res *[]string) {
	if pos == len(digits) {
		*res = append(*res, *str)
		return
	}
	switch(digits[pos]) {
	case '2':
		for i := 0; i < 3; i ++ {
			*str += string('a' + i)
			dfs_17(digits, pos+1, str, res)
			*str = (*str)[:len(*str)-1]
		}
	case '3':
		for i := 0; i < 3; i ++ {
			*str += string('d' + i)
			dfs_17(digits, pos+1, str, res)
			*str = (*str)[:len(*str)-1]
		}
	case '4':
		for i := 0; i < 3; i ++ {
			*str += string('g' + i)
			dfs_17(digits, pos+1, str, res)
			*str = (*str)[:len(*str)-1]
		}
	case '5':
		for i := 0; i < 3; i ++ {
			*str += string('j' + i)
			dfs_17(digits, pos+1, str, res)
			*str = (*str)[:len(*str)-1]
		}
	case '6':
		for i := 0; i < 3; i ++ {
			*str += string('m' + i)
			dfs_17(digits, pos+1, str, res)
			*str = (*str)[:len(*str)-1]
		}
	case '7':
		for i := 0; i < 4; i ++ {
			*str += string('p' + i)
			dfs_17(digits, pos+1, str, res)
			*str = (*str)[:len(*str)-1]
		}
	case '8':
		for i := 0; i < 3; i ++ {
			*str += string('t' + i)
			dfs_17(digits, pos+1, str, res)
			*str = (*str)[:len(*str)-1]
		}
	case '9':
		for i := 0; i < 4; i ++ {
			*str += string('w' + i)
			dfs_17(digits, pos+1, str, res)
			*str = (*str)[:len(*str)-1]
		}
	}
}

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	str := ""
	if len(digits) == 0 {
		return res
	}
	dfs_17(digits, 0, &str, &res)
	return res
}

func main() {
	fmt.Println(letterCombinations("23"))
}
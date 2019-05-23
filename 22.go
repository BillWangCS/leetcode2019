/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
 */

package main

import "fmt"

/*It is better understanding of the whole process, we can just change the left and right to a int to save time and space*/

func dfs_22(leftStack *[]string, left []string, right []string, n int, str *string, res *[]string) {
	if len(*str) == 2 * n {
		tmp := ""
		tmp += *str
		fmt.Println(tmp)
		*res = append(*res, tmp)
		return
	}
	if len(*leftStack) == 0 {
		fmt.Println("hhh1")
		if len(left) > 0 {
			fmt.Println("hhh2")
			*leftStack = append(*leftStack, "(")
			*str += "("
			fmt.Println("hhh", *str)
			dfs_22(leftStack, left[1:], right, n, str, res)
			*leftStack = (*leftStack)[:len(*leftStack)-1]
			*str = (*str)[0:len(*str)-1]
		} else {
			return
		}
	} else if len(*leftStack) == n && len(right) > 0 {
		fmt.Println("hhh3")
		*leftStack = (*leftStack)[:len(*leftStack)-1]
		*str += ")"
		fmt.Println("hhh4")
		dfs_22(leftStack, left, right[1:], n, str, res)
		*leftStack = append(*leftStack, "(")
		*str = (*str)[:len(*str)-1]
	} else {
		fmt.Println("hhh5")
		if len(left) > 0 {
			fmt.Println("hhh6")
			*leftStack = append(*leftStack, "(")
			*str += "("
			dfs_22(leftStack, left[1:], right, n, str, res)
			*leftStack = (*leftStack)[:len(*leftStack)-1]
			*str = (*str)[0:len(*str)-1]
		}
		if len(right) > 0 {
			fmt.Println("hhh7")
			*leftStack = (*leftStack)[:len(*leftStack)-1]
			*str += ")"
			dfs_22(leftStack, left, right[1:], n, str, res)
			*leftStack = append(*leftStack, "(")
			*str = (*str)[:len(*str)-1]
		}
	}
}

func generateParenthesis(n int) []string {
	left := make([]string, n)
	right := make([]string, n)
	for i := 0; i < n; i++ {
		left[i] = "("
		right[i] = ")"
	}
	res := make([]string, 0)
	leftStack := make([]string, 0)
	str := ""
	dfs_22(&leftStack, left, right, n, &str, &res)
	fmt.Println(len(res))
	return res
}

func main() {
	fmt.Println(generateParenthesis(3))
}
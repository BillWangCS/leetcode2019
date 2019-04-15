/*
We are given two sentences A and B.  (A sentence is a string of space separated words.
Each word consists only of lowercase letters.)

A word is uncommon if it appears exactly once in one of the sentences, and does not appear in the other sentence.

Return a list of all uncommon words.

You may return the list in any order.



Example 1:

Input: A = "this apple is sweet", B = "this apple is sour"
Output: ["sweet","sour"]
Example 2:

Input: A = "apple apple", B = "banana"
Output: ["banana"]


Note:

0 <= A.length <= 200
0 <= B.length <= 200
A and B both contain only spaces and lowercase letters.
 */

package main

import (
	"fmt"
	"strings"
)

func uncommonFromSentences(A string, B string) []string {
	res := make([]string, 0)
	dictA := make(map[string]int)
	dictB := make(map[string]int)
	strsA := strings.Split(A, " ")
	strsB := strings.Split(B, " ")
	for _, v := range strsA {
		dictA[v]++
	}
	for _, v := range strsB {
		dictB[v]++
	}
	for _, v := range strsA {
		if dictB[v] == 0 && dictA[v] == 1  {
			res = append(res, v)
		}
	}
	for _, v := range strsB {
		if dictA[v] == 0 && dictB[v] == 1 {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	fmt.Println(uncommonFromSentences("this apple is sweet", "this apple is sour"))
}

/*
You have a list of words and a pattern, and you want to know which words in words matches the pattern.

A word matches the pattern if there exists a permutation of letters p
so that after replacing every letter x in the pattern with p(x), we get the desired word.

(Recall that a permutation of letters is a bijection from letters to letters:
every letter maps to another letter, and no two letters map to the same letter.)

Return a list of the words in words that match the given pattern.

You may return the answer in any order.



Example 1:

Input: words = ["abc","deq","mee","aqq","dkd","ccc"], pattern = "abb"
Output: ["mee","aqq"]
Explanation: "mee" matches the pattern because there is a permutation {a -> m, b -> e, ...}.
"ccc" does not match the pattern because {a -> c, b -> c, ...} is not a permutation,
since a and b map to the same letter.


Note:

1 <= words.length <= 50
1 <= pattern.length = words[i].length <= 20
 */

package main

import "fmt"

func isMatched(s string, p string) bool {
	visitMap := make(map[uint8]uint8)
	visitMap2 := make(map[uint8]uint8)
	if len(s) != len(p) {
		return false
	}
	for i, _ := range p {
		if visitMap[p[i]] == 0 {
			if visitMap2[s[i]] == 0 {
				visitMap[p[i]] = s[i]
				visitMap2[s[i]] = p[i]
			} else if visitMap2[s[i]] != p[i] {
				return false
			}
		} else if visitMap[p[i]] != s[i]  {
			return false
		}
	}
	return true
}

func findAndReplacePattern(words []string, pattern string) []string {
	res := make([]string, 0)
	for _, v := range words {
		if isMatched(v, pattern) {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	fmt.Println(findAndReplacePattern([]string{"abc","deq","mee","aqq","dkd","ccc"}, "abb"))
}
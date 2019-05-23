/*
Given a word, you need to judge whether the usage of capitals in it is right or not.

We define the usage of capitals in a word to be right when one of the following cases holds:

All letters in this word are capitals, like "USA".
All letters in this word are not capitals, like "leetcode".
Only the first letter in this word is capital if it has more than one letter, like "Google".
Otherwise, we define that this word doesn't use capitals in a right way.
Example 1:
Input: "USA"
Output: True
Example 2:
Input: "FlaG"
Output: False
Note: The input will be a non-empty word consisting of uppercase and lowercase latin letters.
 */

package main

import "fmt"

func detectCapitalUse(word string) bool {
	l := len(word)
	if l == 1 {
		return true
	}
	if byte(word[0]) >= 'a' && byte(word[0]) <= 'z' {
		for i := 1; i < l; i++ {
			if word[i] < 'a' || word[i] > 'z' {
				return false
			}
		}
	} else {
		if byte(word[1]) >= 'A' && byte(word[1]) <= 'Z' {
			for i := 2; i < l; i++ {
				if word[i] > 'Z' || word[i] < 'A' {
					return false
				}
			}
		} else {
			for i := 2; i < l; i++ {
				if word[i] <= 'Z' && word[i] >= 'A' {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	fmt.Println(detectCapitalUse("leetcode"))
}
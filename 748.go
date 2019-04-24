/*
Find the minimum length word from a given dictionary words, which has all the letters from the string licensePlate.
Such a word is said to complete the given string licensePlate

Here, for letters we ignore case. For example, "P" on the licensePlate still matches "p" on the word.

It is guaranteed an answer exists. If there are multiple answers, return the one that occurs first in the array.

The license plate might have the same letter occurring multiple times. For example, given a licensePlate of "PP",
the word "pair" does not complete the licensePlate, but the word "supper" does.

Example 1:
Input: licensePlate = "1s3 PSt", words = ["step", "steps", "stripe", "stepple"]
Output: "steps"
Explanation: The smallest length word that contains the letters "S", "P", "S", and "T".
Note that the answer is not "step", because the letter "s" must occur in the word twice.
Also note that we ignored case for the purposes of comparing whether a letter exists in the word.
Example 2:
Input: licensePlate = "1s3 456", words = ["looks", "pest", "stew", "show"]
Output: "pest"
Explanation: There are 3 smallest length words that contains the letters "s".
We return the one that occurred first.
Note:
licensePlate will be a string with length in range [1, 7].
licensePlate will contain digits, spaces, or letters (uppercase or lowercase).
words will have a length in the range [10, 1000].
Every words[i] will consist of lowercase letters, and have length in range [1, 15].
 */

package main

import (
	"fmt"
	"strings"
)

func qualifiedLen(s string, dict map[string]int) int {
	sMap := make(map[string]int)
	for _, v := range s {
		sMap[string(v)]++
	}
	if s == "our" {
		fmt.Println(dict, sMap, s)
	}
	for k, v := range dict {
		if v > sMap[k] {
			return -1
		}
	}
	return len(s)
}

func shortestCompletingWord(licensePlate string, words []string) string {
	dictPlate := make(map[string]int)
	res := ""
	cnt := 1000000
	for _, v := range licensePlate {
		if (string(v) >= "a" && string(v) <= "z") || (string(v) >= "A" && string(v) <= "Z") {
			dictPlate[strings.ToLower(string(v))]++
		}
	}
	for _, s := range words {
		ql := qualifiedLen(s, dictPlate)
		fmt.Println(ql, s)
		if ql < cnt && ql > 0 {
			cnt = ql
			res = s
		}
	}
	return res
}

func main() {
	fmt.Println(shortestCompletingWord("Ar16259", []string{"nature","though","party","food","any","democratic","building","eat","structure","our"}))
}


/*
Given a paragraph and a list of banned words, return the most frequent word that is not in the list of banned words.
It is guaranteed there is at least one word that isn't banned, and that the answer is unique.

Words in the list of banned words are given in lowercase, and free of punctuation.
Words in the paragraph are not case sensitive.  The answer is in lowercase.



Example:

Input:
paragraph = "Bob hit a ball, the hit BALL flew far after it was hit."
banned = ["hit"]
Output: "ball"
Explanation:
"hit" occurs 3 times, but it is a banned word.
"ball" occurs twice (and no other word does), so it is the most frequent non-banned word in the paragraph.
Note that words in the paragraph are not case sensitive,
that punctuation is ignored (even if adjacent to words, such as "ball,"),
and that "hit" isn't the answer even though it occurs more because it is banned.


Note:

1 <= paragraph.length <= 1000.
0 <= banned.length <= 100.
1 <= banned[i].length <= 10.
The answer is unique, and written in lowercase (even if its occurrences in paragraph may have uppercase symbols,
and even if it is a proper noun.)
paragraph only consists of letters, spaces, or the punctuation symbols !?',;.
There are no hyphens or hyphenated words.
Words only consist of letters, never apostrophes or other punctuation symbols.
 */

package main

import (
	"fmt"
	"strings"
)

func trimmer(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		t := s[i]
		if t <= 'Z' && t >= 'A' || t >= 'a' && t <= 'z' {
			res += string(t)
		} else {
			res += " "
		}
	}
	return res
}

func mostCommonWord(paragraph string, banned []string) string {
	res := ""
	bannedDict := make(map[string]bool)
	for i := 0; i < len(banned); i++ {
		bannedDict[banned[i]] = true
	}
	fmt.Println(bannedDict)
	s := trimmer(paragraph)
	words := strings.Split(s, " ")
	dict := make(map[string]int)
	for i := 0; i < len(words); i++ {
		if len(words[i]) == 0 {
			continue
		}
		words[i] = strings.ToLower(words[i])
		if !bannedDict[words[i]] {
			fmt.Println(words[i])
			dict[words[i]]++
		}
	}
	max := 0
	for k, v := range dict {
		if v > max {
			max = v
			res = k
		}
	}
	return res
}

func main() {
	fmt.Println(mostCommonWord("a, a, a, a, b,b,b,c, c", []string{"a"}))
}
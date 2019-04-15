/*
A sentence S is given, composed of words separated by spaces. Each word consists of lowercase and uppercase letters only.

We would like to convert the sentence to "Goat Latin" (a made-up language similar to Pig Latin.)

The rules of Goat Latin are as follows:

If a word begins with a vowel (a, e, i, o, or u), append "ma" to the end of the word.
For example, the word 'apple' becomes 'applema'.

If a word begins with a consonant (i.e. not a vowel), remove the first letter and append it to the end, then add "ma".
For example, the word "goat" becomes "oatgma".

Add one letter 'a' to the end of each word per its word index in the sentence, starting with 1.
For example, the first word gets "a" added to the end, the second word gets "aa" added to the end and so on.
Return the final sentence representing the conversion from S to Goat Latin.



Example 1:

Input: "I speak Goat Latin"
Output: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"
Example 2:

Input: "The quick brown fox jumped over the lazy dog"
Output: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"


Notes:

S contains only uppercase, lowercase and spaces. Exactly one space between each word.
1 <= S.length <= 150.
 */

package main

import (
	"fmt"
	"strings"
)

func getNonVowel(s string) string {
	s1 := string(s[0])
	s2 := string(s[1:])
	return s2+s1
}

func toGoatLatin(S string) string {
	res := make([]string, 0)
	strs := strings.Split(S, " ")
	tmp := "a"
	vowelDict := make(map[string]bool)
	vowelDict["a"] = true
	vowelDict["e"] = true
	vowelDict["i"] = true
	vowelDict["o"] = true
	vowelDict["u"] = true
	for _, str := range strs {
		if vowelDict[strings.ToLower(string(str[0]))] {
			str = str + "ma" + tmp
		} else {
			str = getNonVowel(str) + "ma" + tmp
		}
		res = append(res, str)
		tmp += "a"
	}
	return strings.Join(res, " ")
}

func main() {
	fmt.Println(getNonVowel("speak"))
	fmt.Println(toGoatLatin("I speak Goat Latin"))
}
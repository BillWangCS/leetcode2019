/*
International Morse Code defines a standard encoding where each letter is mapped to a series of dots and dashes,
as follows: "a" maps to ".-", "b" maps to "-...", "c" maps to "-.-.", and so on.

For convenience, the full table for the 26 letters of the English alphabet is given below:

[a:".-",
b:"-...",
c:"-.-.",
d:"-..",
e:".",
f:"..-.",
g:"--.",
h:"....",
i:"..",
j:".---",
k:"-.-",
l:".-..",
m:"--",
n:"-.",
o:"---",
p:".--.",
q:"--.-",
r:".-.",
s:"...",
t:"-",
u:"..-",
v:"...-",
w:".--",
x:"-..-",
y:"-.--",
z:"--.."]
Now, given a list of words, each word can be written as a concatenation of the Morse code of each letter.
For example, "cab" can be written as "-.-..--...", (which is the concatenation "-.-." + "-..." + ".-").
We'll call such a concatenation, the transformation of a word.

Return the number of different transformations among all words we have.

Example:
Input: words = ["gin", "zen", "gig", "msg"]
Output: 2
Explanation:
The transformation of each word is:
"gin" -> "--...-."
"zen" -> "--...-."
"gig" -> "--...--."
"msg" -> "--...--."

There are 2 different transformations, "--...-." and "--...--.".
Note:

The length of words will be at most 100.
Each words[i] will have length in range [1, 12].
words[i] will only consist of lowercase letters.
 */

package main

import "fmt"

func translate(word string, dict map[string]string) string {
	res := ""
	for _, v := range word {
		res += dict[string(v)]
	}
	return res
}

func uniqueMorseRepresentations(words []string) int {
	res := 0
	dictMap := make(map[string]string)
	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	morseArr := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
	for i := 0; i < len(letters); i++ {
		dictMap[letters[i]] = morseArr[i]
	}
	morseMap := make(map[string]bool)
	for _, v := range words {
		morse := translate(v, dictMap)
		if !morseMap[morse] {
			morseMap[morse] = true
			res ++
		}
	}
	return res
}

func main() {
	fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
}
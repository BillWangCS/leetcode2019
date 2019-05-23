/*
Given an arbitrary ransom note string and another string containing letters from all the magazines,
write a function that will return true if the ransom note can be constructed from the magazines ; otherwise, it will return false.

Each letter in the magazine string can only be used once in your ransom note.

Note:
You may assume that both strings contain only lowercase letters.

canConstruct("a", "b") -> false
canConstruct("aa", "ab") -> false
canConstruct("aa", "aab") -> true
 */

package main

func canConstruct(ransomNote string, magazine string) bool {
	dict1 := make(map[byte]int)
	dict2 := make(map[byte]int)
	for _, v := range ransomNote {
		dict1[byte(v)]++
	}
	for _, v := range magazine {
		dict2[byte(v)]++
	}
	for _, v := range ransomNote {
		if dict1[byte(v)] > dict2[byte(v)] {
			return false
		}
	}
	return true
}

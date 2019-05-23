/*
Given a string, sort it in decreasing order based on the frequency of characters.

Example 1:

Input:
"tree"

Output:
"eert"

Explanation:
'e' appears twice while 'r' and 't' both appear once.
So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.
Example 2:

Input:
"cccaaa"

Output:
"cccaaa"

Explanation:
Both 'c' and 'a' appear three times, so "aaaccc" is also a valid answer.
Note that "cacaca" is incorrect, as the same characters must be together.
Example 3:

Input:
"Aabb"

Output:
"bbAa"

Explanation:
"bbaA" is also a valid answer, but "Aabb" is incorrect.
Note that 'A' and 'a' are treated as two different characters.
 */

package main

import "fmt"

type alphaInfo struct {
	 elem rune
	 freq int
}

func sortAlphaInfo(alphaArr *[]alphaInfo, start, end int) {
	if start > end {
		return
	}
	left := start
	right := end
	pivot := (*alphaArr)[start + (end - start) / 2]
	for left <= right {
		for left <= right && (*alphaArr)[left].freq > pivot.freq {
			left++
		}
		for left <= right && (*alphaArr)[right].freq < pivot.freq {
			right--
		}
		if left <= right {
			(*alphaArr)[left], (*alphaArr)[right] = (*alphaArr)[right], (*alphaArr)[left]
			left++
			right--
		}
	}
	sortAlphaInfo(alphaArr, left, end)
	sortAlphaInfo(alphaArr, start, right)
}

func frequencySort(s string) string {
	res := ""
	dict := make(map[rune]int)
	for _, v := range s {
		dict[v]++
	}
	alphaArr := make([]alphaInfo, 0)
	for k, v := range dict {
		alphaArr = append(alphaArr, alphaInfo{k, v})
	}
	sortAlphaInfo(&alphaArr, 0, len(alphaArr) - 1)
	for _, v := range alphaArr {
		for i := 0; i < v.freq; i++ {
			res += string(v.elem)
		}
	}
	return res
}

func main() {
	fmt.Println(frequencySort("t"))
}
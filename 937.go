/*
You have an array of logs.  Each log is a space delimited string of words.

For each log, the first word in each log is an alphanumeric identifier.  Then, either:

Each word after the identifier will consist only of lowercase letters, or;
Each word after the identifier will consist only of digits.
We will call these two varieties of logs letter-logs and digit-logs.
It is guaranteed that each log has at least one word after its identifier.

Reorder the logs so that all of the letter-logs come before any digit-log.
The letter-logs are ordered lexicographically ignoring identifier, with the identifier used in case of ties.
The digit-logs should be put in their original order.

Return the final order of the logs.



Example 1:

Input: ["a1 9 2 3 1","g1 act car","zo4 4 7","ab1 off key dog","a8 act zoo"]
Output: ["g1 act car","a8 act zoo","ab1 off key dog","a1 9 2 3 1","zo4 4 7"]


Note:

0 <= logs.length <= 100
3 <= logs[i].length <= 100
logs[i] is guaranteed to have an identifier, and a word after the identifier.
 */

package main

import (
	"fmt"
	"strings"
)

func compareWord(w1, w2 string) bool {
	i := 0
	j := 0
	lw1 := len(w1)
	lw2 := len(w2)
	for i < lw1 && j < lw2 {
		if w1[i] > w2[j] {
			return true
		} else if w1[i] < w2[j] {
			return false
		}
		i++
		j++
	}
	if lw1 > lw2 {
		return true
	}
	return false
}

func quickSortStrings(strs *[]string, start, end int) {
	if start >= end {
		return
	}
	left, right := start, end
	pivot := (*strs)[start + (end - start) / 2]
	for left <= right {
		for left <= right && compareWord(strings.Join(strings.Split(pivot, " ")[1:], " "), strings.Join(strings.Split((*strs)[left], " ")[1:], " ")) {
			left++
		}
		for left <= right && compareWord(strings.Join(strings.Split((*strs)[right], " ")[1:], " "), strings.Join(strings.Split(pivot, " ")[1:], " ")) {
			right--
		}
		if left <= right {
			(*strs)[left], (*strs)[right] = (*strs)[right], (*strs)[left]
			left++
			right--
		}
	}
	quickSortStrings(strs, left, end)
	quickSortStrings(strs, start, right)
}

func isDigit(s string) bool {
	for _, v := range s {
		if !(byte(v) >= '0' && byte(v) <= '9') {
			return false
		}
	}
	return true
}

func reorderLogFiles(logs []string) []string {
	strs := make([]string, 0)
	digits := make([]string, 0)
	for _, s := range logs {
		digit := strings.Split(s, " ")[1]
		if isDigit(digit) {
			digits = append(digits, s)
		} else {
			strs = append(strs, s)
		}
	}
	quickSortStrings(&strs, 0, len(strs) - 1)
	strs = append(strs, digits...)
	return strs
}

func main() {
	fmt.Println(reorderLogFiles([]string{"1 n u", "r 527", "j 893", "6 14", "6 82"}))
	fmt.Println(compareWord("act", "off"))
}
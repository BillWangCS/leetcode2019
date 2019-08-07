/*
Your friend is typing his name into a keyboard.
Sometimes, when typing a character c, the key might get long pressed, and the character will be typed 1 or more times.

You examine the typed characters of the keyboard.
Return True if it is possible that it was your friends name, with some characters (possibly none) being long pressed.



Example 1:

Input: name = "alex", typed = "aaleex"
Output: true
Explanation: 'a' and 'e' in 'alex' were long pressed.
Example 2:

Input: name = "saeed", typed = "ssaaedd"
Output: false
Explanation: 'e' must have been pressed twice, but it wasn't in the typed output.
Example 3:

Input: name = "leelee", typed = "lleeelee"
Output: true
Example 4:

Input: name = "laiden", typed = "laiden"
Output: true
Explanation: It's not necessary to long press any character.


Note:

name.length <= 1000
typed.length <= 1000
The characters of name and typed are lowercase letters.
 */


package main

import "fmt"

type characterInfo struct {
	s string
	freq int
}

func constructArr(s string) []characterInfo {
	arr := make([]characterInfo, 0)
	if len(s) == 0 {
		return arr
	}
	l := s[0]
	cnt := 1
	for i := 1; i < len(s); i++ {
		if s[i] == l {
			cnt++
		} else {
			arr = append(arr, characterInfo{string(l), cnt})
			cnt = 1
		}
		l = s[i]
	}
	arr = append(arr, characterInfo{string(l), cnt})
	return arr
}

func isLongPressedName(name string, typed string) bool {
	A := constructArr(name)
	B := constructArr(typed)
	if len(A) != len(B) {
		return false
	}
	for i := 0; i < len(A); i++ {
		if A[i].freq > B[i].freq {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isLongPressedName("alex", "aaleex"))
}
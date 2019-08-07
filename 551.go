/*
You are given a string representing an attendance record for a student. The record only contains the following three characters:
'A' : Absent.
'L' : Late.
'P' : Present.
A student could be rewarded if his attendance record doesn't contain more than one 'A' (absent) or more than two continuous 'L' (late).

You need to return whether the student could be rewarded according to his attendance record.

Example 1:
Input: "PPALLP"
Output: True
Example 2:
Input: "PPALLL"
Output: False
 */

package main

import (
	"fmt"
	"strings"
)

func checkRecord(s string) bool {
	if strings.Index(s, "LLL") != -1 {
		return false
	}
	cnt := 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "A" {
			cnt++
		}
	}
	if cnt > 1 {
		return false
	}
	return true
}

func main() {
	fmt.Println(checkRecord("PPALLP"))
}
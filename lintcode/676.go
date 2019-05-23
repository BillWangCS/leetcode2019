/*
A message containing letters from A-Z is being encoded to numbers using the following mapping way:

'A' -> 1
'B' -> 2
...
'Z' -> 26
Beyond that, now the encoded string can also contain the character *, which can be treated as one of the numbers from 1 to 9.
Given the encoded message containing digits and the character *, return the total number of ways to decode it.
Also, since the answer may be very large, you should return the output mod 10^9 + 7.

Example
Example 1

Input: "*"
Output: 9
Explanation: You can change it to "A", "B", "C", "D", "E", "F", "G", "H", "I".
Example 2

Input: "1*"
Output: 18
Notice
The length of the input string will fit in range [1, 10^5].
The input string will only contain the character * and digits 0 - 9.
 */

package main

import "fmt"

func numDecodings(s string) int {
	count := 1
	prvCount := 1
	for i := 0; i < len(s); i++ {
		newCount := 0
		if i > 0  {
			if s[i-1] == '1' && s[i] != '*' || (s[i-1] == '2' && s[i] < '7' && s[i] != '*') {
				newCount = (newCount + prvCount) % 1000000007
			} else if s[i-1] == '*' && s[i] < '7' && s[i] != '*' {
				newCount = (newCount + prvCount * 2) % 1000000007
			} else if s[i-1] == '1' && s[i] == '*' {
				newCount = (newCount + prvCount * 9) % 1000000007
			} else if s[i-1] == '2' && s[i] == '*' {
				newCount = (newCount + prvCount * 6) % 1000000007
			} else if s[i-1] == '*' && s[i] == '*' {
				newCount = (newCount + prvCount * 15) % 1000000007
			}
		}
		if s[i] != '*' && s[i] > '0' {
			newCount = (newCount + count) % 1000000007
		} else if s[i] == '*' {
			newCount = (newCount + count*9) % 1000000007
		}
		prvCount = count
		count = newCount
	}
	return count
}

func main() {
	fmt.Println(numDecodings("**1**"))
}
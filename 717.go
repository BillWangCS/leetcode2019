/*
We have two special characters. The first character can be represented by one bit 0.
The second character can be represented by two bits (10 or 11).

Now given a string represented by several bits. Return whether the last character must be a one-bit character or not.
The given string will always end with a zero.

Example 1:
Input:
bits = [1, 0, 0]
Output: True
Explanation:
The only way to decode it is two-bit character and one-bit character. So the last character is one-bit character.
Example 2:
Input:
bits = [1, 1, 1, 0]
Output: False
Explanation:
The only way to decode it is two-bit character and two-bit character. So the last character is NOT one-bit character.
Note:

1 <= len(bits) <= 1000.
bits[i] is always 0 or 1.
 */

package main

import "fmt"

func isOneBitCharacter(bits []int) bool {
	l := len(bits)
	if l == 1 {
		return true
	}
	if l > 1 && bits[l-2] == 0 {
		return true
	}
	i := 0
	for i + 1 < l {
		if i == l - 2 {
			return false
		}
		if bits[i] == 0 {
			i += 1
		} else {
			i += 2
		}
	}
	return true
}

func main() {
	fmt.Println(isOneBitCharacter([]int{0,1,0}))
}
/*
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai).
n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines,
which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.





The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case,
the max area of water (blue section) the container can contain is 49.



Example:

Input: [1,8,6,2,5,4,8,3,7]
Output: 49
 */

package main

import "fmt"

func maxArea(height []int) int {
	l := len(height)
	if l <= 1 {
		return 0
	}
	s := 0
	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			sTmp := 0
			if height[i] > height[j] {
				sTmp = (i - j) * height[j]
			} else {
				sTmp = (i - j) * height[i]
				if sTmp > s {
					s = sTmp
				}
				break
			}
			if sTmp > s {
				s = sTmp
			}
		}
	}
	return s
}

func main() {
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
}
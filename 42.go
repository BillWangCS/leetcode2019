/*
Given n non-negative integers representing an elevation map where the width of each bar is 1,
compute how much water it is able to trap after raining.


The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1].
In this case, 6 units of rain water (blue section) are being trapped. Thanks Marcos for contributing this image!

Example:

Input: [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
 */

package main

import "fmt"

func trap(height []int) int {
	l := len(height)
	if l < 3 {
		return 0
	}
	res := 0
	max := 0
	i := 0
	j := 0
	for i < l {
		if height[i] == 0 && max == 0 {
			i++
		}
		if height[i] > max {
			max = height[i]
		}
		for j = i+1; j < l; j++ {
			if height[j] >= height[i] {
				for k := i+1; k < j; k++ {
					res += height[i] - height[k]
				}
				i = j
			} else if height[j] > height[j-1] {
				for k := j-1; k > i; k-- {
					if height[k] < height[j] {
						res += height[j] - height[k]
						height[k] = height[j]
					} else {
						break
					}
				}
			}
		}
		if j >= l {
			break
		}
	}
	return res
}

func main() {
	fmt.Println(trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
}
/*
Given an integer array with no duplicates. A maximum tree building on this array is defined as follow:

The root is the maximum number in the array.
The left subtree is the maximum tree constructed from left part subarray divided by the maximum number.
The right subtree is the maximum tree constructed from right part subarray divided by the maximum number.
Construct the maximum tree by the given array and output the root node of this tree.

Example 1:
Input: [3,2,1,6,0,5]
Output: return the tree root node representing the following tree:

      6
    /   \
   3     5
    \    /
     2  0
       \
        1
Note:
The size of the given array will be in the range [1,1000].
 */

package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
 	Val int
 	Left *TreeNode
 	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	fmt.Println(nums, len(nums))
	if len(nums) == 0 {
		return nil
	}
	max := math.MinInt32
	index := 0
	for i, v := range nums {
		if v > max {
			max = v
			index = i
		}
	}
	root := &TreeNode{
		Val: max,
	}
	root.Left = constructMaximumBinaryTree(nums[:index])
	root.Right = constructMaximumBinaryTree(nums[index+1:])
	return root
}

func main() {
	constructMaximumBinaryTree([]int{3,2,1,6,0,5})
}
/*
Given the root of a binary tree, find the maximum value V for which there exists different
nodes A and B where V = |A.val - B.val| and A is an ancestor of B.

(A node A is an ancestor of B if either: any child of A is equal to B, or any child of A is an ancestor of B.)



Example 1:



Input: [8,3,10,1,6,null,14,null,null,4,7,13]
Output: 7
Explanation:
We have various ancestor-node differences, some of which are given below :
|8 - 3| = 5
|3 - 7| = 4
|8 - 1| = 7
|10 - 13| = 3
Among all possible differences, the maximum value of 7 is obtained by |8 - 1| = 7.


Note:

The number of nodes in the tree is between 2 and 5000.
Each node will have value between 0 and 100000.
 */

package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxInteger(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxAncestorDiffHelper(root *TreeNode, res *int) (maxnum int, minnum int) {
	if root.Left == nil && root.Right == nil {
		return root.Val, root.Val
	} else if root.Left != nil && root.Right == nil {
		leftMax, leftMin := maxAncestorDiffHelper(root.Left, res)
		if *res < leftMax - root.Val {
			*res = leftMax - root.Val
		}
		if *res < root.Val - leftMin {
			*res = root.Val - leftMin
		}
		return maxInteger(leftMax, root.Val), minInt(leftMin, root.Val)
	} else if root.Left == nil && root.Right != nil {
		rightMax, rightMin := maxAncestorDiffHelper(root.Right, res)
		if *res < rightMax - root.Val {
			*res = rightMax - root.Val
		}
		if *res < root.Val - rightMin {
			*res = root.Val - rightMin
		}
		return maxInteger(rightMax, root.Val), minInt(rightMin, root.Val)
	} else {
		leftMax, leftMin := maxAncestorDiffHelper(root.Left, res)
		rightMax, rightMin := maxAncestorDiffHelper(root.Right, res)
		fmt.Println(leftMax, leftMin, rightMax, rightMin, root.Val)
		if *res < root.Val - leftMin {
			*res = root.Val - leftMin
		}
		if *res < root.Val - rightMin {
			*res = root.Val - rightMin
		}
		if *res < leftMax - root.Val {
			*res = leftMax - root.Val
		}
		if *res < rightMax - root.Val {
			*res = rightMax - root.Val
		}
		return maxInteger(maxInteger(leftMax, rightMax), root.Val), minInt(minInt(leftMin, rightMin), root.Val)
	}
}

func maxAncestorDiff(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}
	maxAncestorDiffHelper(root, &res)
	return res
}
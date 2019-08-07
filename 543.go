/*
Given a binary tree, you need to compute the length of the diameter of the tree.
The diameter of a binary tree is the length of the longest path between any two nodes in a tree.
This path may or may not pass through the root.

Example:
Given a binary tree
          1
         / \
        2   3
       / \
      4   5
Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].

Note: The length of path between two nodes is represented by the number of edges between them.
 */

package main

import "fmt"

func maxInteger(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return maxInteger(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func dfs_543(root *TreeNode, max *int) {
	if root == nil {
		return
	}
	fmt.Println(root.Val, maxDepth(root.Left), maxDepth(root.Right))
	m := maxDepth(root.Left) + maxDepth(root.Right)
	if m > *max {
		*max = m
	}
	dfs_543(root.Left, max)
	dfs_543(root.Left, max)
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := 0
	dfs_543(root, &max)
	return max
}

// [4,-7,-3,null,null,-9,-3,9,-7,-4,null,6,null,-6,-6,null,null,0,6,5,null,9,null,null,-1,-4,null,null,null,-2]
/*
Given a Binary Search Tree (BST) with the root node root,
return the minimum difference between the values of any two different nodes in the tree.

Example :

Input: root = [4,2,6,1,3,null,null]
Output: 1
Explanation:
Note that root is a TreeNode object, not an array.

The given tree [4,2,6,1,3,null,null] is represented by the following diagram:

          4
        /   \
      2      6
     / \
    1   3

while the minimum difference in this tree is 1, it occurs between node 1 and node 2, also between node 3 and node 2.
Note:

The size of the BST will be between 2 and 100.
The BST is always valid, each node's value is an integer, and each node's value is different.
 */

package main

import "math"

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

func dfs_783(root *TreeNode, prevVal *int, minimum *int) {
	if root == nil {
		return
	}
	dfs_783(root.Left, prevVal, minimum)
	if abs(root.Val - *prevVal) < *minimum {
		*minimum = abs(root.Val - *prevVal)
	}
	*prevVal = root.Val
	dfs_783(root.Right, prevVal, minimum)
}

func minDiffInBST(root *TreeNode) int {
	res := math.MaxInt32
	prevVal := math.MinInt32
	dfs_783(root, &prevVal, &res)
	return res
}
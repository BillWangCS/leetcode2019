/*
Given a binary search tree with non-negative values, find the minimum absolute difference between values of any two nodes.

Example:

Input:

   1
    \
     3
    /
   2

Output:
1

Explanation:
The minimum absolute difference is 1, which is the difference between 2 and 1 (or between 2 and 3).


Note: There are at least two nodes in this BST.
 */

package main

import "math"

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

func dfs_530(root *TreeNode, prevVal *int, minimum *int) {
	if root == nil {
		return
	}
	dfs_530(root.Left, prevVal, minimum)
	if abs(root.Val - *prevVal) < *minimum {
		*minimum = abs(root.Val - *prevVal)
	}
	*prevVal = root.Val
	dfs_530(root.Right, prevVal, minimum)
}

func getMinimumDifference(root *TreeNode) int {
	res := math.MaxInt32
	prevVal := math.MinInt32
	dfs_530(root, &prevVal, &res)
	return res
}
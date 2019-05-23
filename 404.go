/*
Find the sum of all left leaves in a given binary tree.

Example:

    3
   / \
  9  20
    /  \
   15   7

There are two left leaves in the binary tree, with values 9 and 15 respectively. Return 24.
 */

package main

func dfs_404(root *TreeNode, sum *int, fromLeft bool) {
	if root == nil {
		return
	}
	if fromLeft && root.Left == nil && root.Right == nil {
		*sum += root.Val
	}
	dfs_404(root.Left, sum, true)
	dfs_404(root.Right, sum, false)
}

func sumOfLeftLeaves(root *TreeNode) int {
	res := 0
	dfs_404(root, &res, false)
	return res
}